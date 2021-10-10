package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	var port int

	flag.IntVar(&port, "port", 0, "port number")
	flag.Parse()

	if port == 0 {
		fmt.Println("port missing, please include: --port ####")
		return
	}

	b, err := exec.Command("lsof", "-t", "-i", fmt.Sprintf(":%d", port)).Output()
	if err != nil {
		if err.Error() == "exit status 1" {
			fmt.Println("running process not found")
			return
		} else {
			panic(err)
		}
	}

	pid := strings.TrimSpace(string(b))

	_, err = exec.Command("kill", "-kill", pid).Output()
	if err != nil {
		panic(err)
	}
}

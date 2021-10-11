package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	var port int
	var err error
	flag.Parse()

	if len(flag.Args()) == 1 {
		port, err = strconv.Atoi(flag.Args()[0])
		if err != nil {
			fmt.Println("Error: invalid port number")
			return
		}
	}

	if port == 0 {
		fmt.Println("Error: port missing number, useage: killport ####")
		return
	}

	b, err := exec.Command("lsof", "-t", "-i", fmt.Sprintf(":%d", port)).Output()
	if err != nil {
		if err.Error() == "exit status 1" {
			fmt.Println("Warning: running process not found")
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

	fmt.Printf("Success: process %s is dead\n", pid)
}

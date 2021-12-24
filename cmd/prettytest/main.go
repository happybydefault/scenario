package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	args := append([]string{"test"}, os.Args[1:]...)
	cmd := exec.Command("go", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if ok {
			os.Exit(exitErr.ExitCode())
		}
		log.Println(err)
		os.Exit(1)
	}
}

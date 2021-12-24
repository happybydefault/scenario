package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/gookit/color"
)

type prettifier struct {
}

func (w *prettifier) Write(p []byte) (n int, err error) {
	lines := strings.Split(string(p), "\n")
	for _, line := range lines {
		trimmed := strings.TrimLeft(line, " ")
		switch {
		case strings.HasPrefix(line, "PASS"):
			color.Greenln(line + " ✓")
		case strings.HasPrefix(line, "FAIL"):
			color.Redln(line + " ✘")
		case strings.HasPrefix(trimmed, "--- PASS: "):
			color.Greenln(line + " ✓")
		case strings.HasPrefix(trimmed, "--- FAIL: "):
			color.Redln(line + " ✘")
		default:
			fmt.Println(line)
		}
	}
	return len(p), err
}

func main() {
	args := append([]string{"test"}, os.Args[1:]...)
	cmd := exec.Command("go", args...)

	p := prettifier{}

	cmd.Stdout = &p
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if ok {
			os.Exit(exitErr.ExitCode())
		}
		log.Printf("error while running command: %v", err)
		os.Exit(1)
	}
}

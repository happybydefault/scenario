package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/gookit/color"
)

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

type prettifier struct{}

func (p *prettifier) Write(data []byte) (int, error) {
	if len(data) == 0 {
		return 0, nil
	}

	r := bufio.NewReader(bytes.NewReader(data))

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return 0, err
		}

		line = pretty(line)
		fmt.Print(line)
	}

	return len(data), nil
}

func pretty(line string) string {
	trimmed := strings.TrimLeft(line, " ")

	switch {
	case line == "PASS\n":
		return color.Green.Sprint(line)
	case line == "FAIL\n":
		return color.Red.Sprint(line)
	case
		strings.HasPrefix(line, "ok "),
		strings.HasPrefix(trimmed, "--- PASS: "):

		line = appendToLine(line, "✓")
		return color.Green.Sprint(line)
	case
		strings.HasPrefix(trimmed, "--- FAIL: "),
		strings.HasPrefix(line, "FAIL"):

		line = appendToLine(line, "✘")
		return color.Red.Sprint(line)
	}

	return line
}

func appendToLine(line string, suffix string) string {
	hasLineBreak := strings.HasSuffix(line, "\n")
	if hasLineBreak {
		line = strings.TrimSuffix(line, "\n")
	}

	line = fmt.Sprintf("%s %s", line, suffix)

	if hasLineBreak {
		line = line + "\n"
	}

	return line
}

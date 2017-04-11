package model

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

// InitStock ...
func InitStock() {
	fmt.Println("")
	// 1. connect to database
	// 2. get if inited
	// 3. init or update stock
	cmd := exec.Command("/usr/bin/python", "script/stock.py")
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		panic(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	go copyOutput(stdout)
	go copyOutput(stderr)
	cmd.Wait()

}

func copyOutput(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// UpdateStock ...
func UpdateStock() {
	fmt.Println("")
}

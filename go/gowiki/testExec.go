package main

import (
        "fmt"
        "os/exec"
	"bytes"
)

func main() {
	var out = ""
	cmd := exec.Command("python", "num.py")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		out += "ERROR\n"
		out += string(stderr.Bytes())
	}
	out += string(stdout.Bytes())
	fmt.Println(out)
}

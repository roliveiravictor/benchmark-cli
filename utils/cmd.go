package utils

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func Cmd(dir string, name string, arg ...string) {
	var log bytes.Buffer
	var cmdErr bytes.Buffer

	onTheFlyWriter := io.MultiWriter(os.Stdout, &log)

	cmd := exec.Command(name, arg[:]...)

	if dir != "" {
		cmd.Dir = dir
	}

	//cmd.Stdout = &log
	cmd.Stderr = &cmdErr
	cmd.Stdout = onTheFlyWriter

	err := cmd.Run()

	byteNilCheck(err, cmdErr)

	fmt.Println(log.String())
}

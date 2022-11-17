package utils

import (
	"bytes"
	"fmt"
)

func stringNilCheck(err error, message string) {
	if err != nil {
		exception := fmt.Errorf(message)
		if exception != nil {
			fmt.Println(err)
		}
	}
}

func byteNilCheck(err error, buffer bytes.Buffer) {
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + buffer.String())
	}
}

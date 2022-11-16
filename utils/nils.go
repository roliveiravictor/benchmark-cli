package utils

import (
	"fmt"
)

func nilCheck(err error, message string) {
	if err != nil {
		fmt.Errorf(message)
	}
}

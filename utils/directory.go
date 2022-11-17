package utils

import "os/user"

func Home() string {
	usr, err := user.Current()
	stringNilCheck(err, "Failed to acquire home directory")
	return usr.HomeDir
}

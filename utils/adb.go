package utils

func Uninstall(path string, appId string) {
	adb := "adb"
	uninstall := "uninstall"
	Cmd(path, adb, uninstall, appId)
}

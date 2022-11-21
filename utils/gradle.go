package utils

import (
	"os"
	"strings"
)

func GradleStop(path string) {
	gradle := "./gradlew"
	arg0 := "--stop"
	Cmd(path, gradle, arg0)
}

func GradleRun(path string) {
	files, err := os.ReadDir(os.Getenv("SDKMAN_CANDIDATES_DIR") + "/java")
	stringNilCheck(err, "Failed reading root dir")

	var link string
	for _, file := range files {
		if strings.Contains(file.Name(), "amzn") {
			link = file.Name()
			break
		}
	}

	gradle := "./gradlew"
	arg0 := "clean"
	arg1 := ":benchmark:connectedAndroidTest"
	arg2 := "-Dorg.gradle.java.home=" + os.Getenv("SDKMAN_CANDIDATES_DIR") + "/java/" + link
	Cmd(path, gradle, arg0, arg1, arg2)
}

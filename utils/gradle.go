package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"strings"
)

func GradleStop(path string) {
	gradle := "./gradlew"
	arg0 := "--stop"
	Cmd(path, gradle, arg0)
}

func GradleConnectedAndroidTest(path string, module string) {
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
	arg1 := ":" + module + ":connected" + cases.Title(language.Und).String(module) + "AndroidTest"
	arg2 := "-Dorg.gradle.java.home=" + os.Getenv("SDKMAN_CANDIDATES_DIR") + "/java/" + link
	arg3 := "-Dorg.gradle.unsafe.configuration-cache=false"
	Cmd(path, gradle, arg0, arg1, arg2, arg3)
}

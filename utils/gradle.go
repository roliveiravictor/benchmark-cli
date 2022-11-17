package utils

import (
	"os"
	"strings"
)

func GradleStop() {
	dir := Home() + "/StudioProjects/fury_startup-initializer-android/"
	gradle := "./gradlew"
	arg0 := "--stop"
	Cmd(dir, gradle, arg0)
}

func GradleRun() {
	files, err := os.ReadDir(os.Getenv("SDKMAN_CANDIDATES_DIR") + "/java")
	stringNilCheck(err, "Failed reading root dir")

	var link string
	for _, file := range files {
		if strings.Contains(file.Name(), "amzn") {
			link = file.Name()
			break
		}
	}

	dir := Home() + "/StudioProjects/play-store-buy-me-a-coffee/"
	gradle := "./gradlew"
	arg0 := "clean"
	arg1 := ":app:assemble"
	arg2 := "-Dorg.gradle.java.home=" + os.Getenv("SDKMAN_CANDIDATES_DIR") + "/java/" + link
	Cmd(dir, gradle, arg0, arg1, arg2)
}

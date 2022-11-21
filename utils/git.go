package utils

func Checkout(path string, branch string) {
	git := "git"

	fetch := "fetch"
	Cmd(path, git, fetch)

	checkout := "checkout"
	Cmd(path, git, checkout, branch)

	pull := "pull"
	origin := "origin"
	Cmd(path, git, pull, origin, branch)
}

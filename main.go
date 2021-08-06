package main

import (
	"gabor-boros/travaux-crawler/cmd"
)

var version string
var commit string

func main() {
	cmd.Execute(version, commit)
}

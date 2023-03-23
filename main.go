package main

import "github.com/thiagozs/aws-tools/cmd"

var (
	Version string
)

func main() {
	cmd.Version = Version
	cmd.Execute()
}

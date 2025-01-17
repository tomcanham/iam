package main

import (
	"github.com/tomcanham/iam/cli"
	"github.com/tomcanham/iam/cmd"
)

func main() {
	cmd.Execute(cli.New())
}

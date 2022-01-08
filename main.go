package main

import (
	"github.com/alierkilic/do-cli/cmd"
	"github.com/alierkilic/do-cli/data"
)

func main() {
	data.Open()
	cmd.RootCmd.Execute()
}

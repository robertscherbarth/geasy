package main

import (
	"github.com/spf13/cobra"

	"github.com/robertscherbarth/geasy/cmd"
)

func main() {

	command := cmd.NewRootCommand()

	cobra.CheckErr(command.Execute())
}

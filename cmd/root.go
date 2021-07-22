package cmd

import (
	"github.com/robertscherbarth/geasy/cmd/get"
	"github.com/robertscherbarth/geasy/internal/flags"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   "geasy",
		Short: "The geasy is a tool to evaluate gcp resources ...",
	}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().BoolVarP(&flags.Verbose, "verbose", "v", false, "verbose output")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(
		get.NewGetCmd(),
	)

	return rootCmd
}

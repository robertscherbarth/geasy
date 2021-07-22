package get

import (
	"github.com/robertscherbarth/geasy/internal/flags"
	"github.com/spf13/cobra"
)

func NewGetCmd() *cobra.Command {
	cmdGet := &cobra.Command{
		Use:   "get",
		Short: "get specific gcp resources information",
	}
	cmdGet.PersistentFlags().StringVarP(&flags.Output, "output", "o", "json", "specify the output format")

	cmdGet.AddCommand(
		newRuntimeConfigurationCmd(),
	)

	return cmdGet
}

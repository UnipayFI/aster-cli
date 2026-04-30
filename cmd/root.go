package cmd

import (
	"errors"

	"github.com/UnipayFI/aster-cli/config"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:               "aster-cli",
	Short:             "AsterDEX API for CLI version",
	PersistentPreRunE: checkCredentials,
}

func init() {
	initCommandConfig()
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func initCommandConfig() {
	RootCmd.CompletionOptions.DisableDefaultCmd = true
	RootCmd.PersistentFlags().BoolVar(&config.Config.OutputJSON, "json", false, "Output JSON instead of a table")
}

func checkCredentials(cmd *cobra.Command, args []string) error {
	if config.Config.APIAddress == "" || config.Config.APIPrivateKey == "" {
		return errors.New("API_ADDRESS and API_PRIVATE_KEY must be set")
	}
	return nil
}

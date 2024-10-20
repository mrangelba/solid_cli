package cmd

import (
	"github.com/mrangelba/solid_cli/cmd/account"
	"github.com/mrangelba/solid_cli/cmd/files"
	"github.com/mrangelba/solid_cli/cmd/pod"
	"github.com/spf13/cobra"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "solid-cli",
		Short: "A SOLID Community management CLI",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(pod.PodCmd, account.AccountCmd, files.FilesCmd)
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of SOLID CLI",
	Long:  `All software has versions. This is SOLID CLI's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SOLID CLI v0.1")
	},
}

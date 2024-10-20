package account

import (
	"github.com/spf13/cobra"
)

var (
	AccountCmd = &cobra.Command{
		Use:   "account",
		Short: "Manage SOLID accounts",
		Long:  `Manage SOLID accounts, including listing and deleting them.`,
	}
)

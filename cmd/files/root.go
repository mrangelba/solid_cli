package files

import (
	"github.com/spf13/cobra"
)

var (
	FilesCmd = &cobra.Command{
		Use:   "files",
		Short: "Manage account files",
		Long:  `Listing account files.`,
	}
)

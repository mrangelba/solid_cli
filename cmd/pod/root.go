package pod

import (
	"github.com/spf13/cobra"
)

var (
	PodCmd = &cobra.Command{
		Use:   "pod",
		Short: "Manage SOLID Pods",
		Long:  `Manage SOLID Pods, including listing and deleting them.`,
	}
)

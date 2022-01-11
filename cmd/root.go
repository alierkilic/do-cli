package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
	Long: `Task is a CLI task manager written in Go.
                Complete documentation is available at https://github.com/alierkilic/do-cli`,
}

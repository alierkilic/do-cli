package cmd

import (
	"fmt"

	"github.com/alierkilic/do-cli/data"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		err := data.CreateTable()
		if err != nil {
			fmt.Printf("Error happended creating table %v", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}

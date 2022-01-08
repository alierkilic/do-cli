package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		switch taskLength := len(args); taskLength {
		case 0:
			fmt.Println("Please give a task")
		case 1:
			fmt.Printf("added task: %s", args[0])
		default:
			fmt.Printf("added tasks: ")
			for _, arg := range args {
				fmt.Printf("%s ", arg)
			}
			fmt.Println()
		}

	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}

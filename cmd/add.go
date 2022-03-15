package cmd

import (
	"fmt"
	"strings"

	"github.com/alierkilic/do-cli/data"
	"github.com/alierkilic/do-cli/model"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		switch taskLength := len(args); taskLength {
		case 0:
			fmt.Println("Don't be lazy... give me a task")
		default:

			dstatus, _ := cmd.Flags().GetBool("daily")
			task := &model.NewTask{
				Task:  strings.Join(args, " "),
				Daily: dstatus,
			}
			data.SaveTask(task)
			fmt.Printf("added task: %s\n", task.Task)
		}

	},
}

func init() {
	RootCmd.AddCommand(addCmd)

	addCmd.Flags().BoolP("daily", "D", false, "List daily items")
}

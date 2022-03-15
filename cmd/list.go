package cmd

import (
	"fmt"

	"github.com/alierkilic/do-cli/data"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {

		dstatus, _ := cmd.Flags().GetBool("done")
		if dstatus { // if status is true, call addFloat
			ListDoneTodos()
		} else {
			dailystatus, _ := cmd.Flags().GetBool("daily")
			if dailystatus {
				ListTodos(true)
			} else {
				ListTodos(false)
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("done", "d", false, "List done items")
	listCmd.Flags().BoolP("daily", "D", false, "List daily items")

}

func ListTodos(daily bool) {
	tasks := data.GetTasks(daily)
	for _, task := range tasks {
		fmt.Printf("Task %d: %s\n", task.ID, task.Task)
	}
}

func ListDoneTodos() {
	tasks := data.GetDoneTasks()
	for _, task := range tasks {
		fmt.Printf("Task %d: %s\n", task.ID, task.Task)
	}
}

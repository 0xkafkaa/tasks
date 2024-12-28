package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteTaskCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long: `Delete a task from the task list based on the task ID.
	For example:
	
	task delete -d <taskid>
	task delete -d 34`,
	Run: deleteTasks,
}

func deleteTasks(cmd *cobra.Command, args []string) {
	taskID, _ := cmd.Flags().GetString("task-id")
	records := readCsvFile("/home/mklno/projects/tasks/tasks.csv")
	deleteId := taskID
	var updatedRecords [][]string
	for _, row := range records {
		if len(row) > 0 && row[0] != deleteId {
			updatedRecords = append(updatedRecords, row)
		}
	}
	updateCSV(updatedRecords)
	fmt.Println("Task has been deleted")
}

func init() {
	rootCmd.AddCommand(deleteTaskCmd)
	deleteTaskCmd.Flags().StringP("task-id", "d", "0", "Enter task ID to delete")
}

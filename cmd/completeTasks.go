package cmd

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var completeTasksCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a task.",
	Long: `Complete a task from the task list based on the task ID.
	For example:
	
	task complete -t <taskid>
	task complete -t 34`,
	Run: completeTasks,
}

func completeTasks(cmd *cobra.Command, args []string) {
	taskID, _ := cmd.Flags().GetString("task-id")

	records := readCsvFile("/home/mklno/projects/tasks/tasks.csv")
	updateId := taskID

	for i, row := range records {
		if len(row) > 0 && row[0] == updateId {
			// update status
			records[i] = []string{row[0], row[1], row[2], "true"}
			break
		}
	}
	// file, _ := os.Create("/home/mklno/projects/tasks/tasks.csv")
	// defer file.Close()
	// csvWriter := csv.NewWriter(file)
	// csvWriter.WriteAll(records)
	updateCSV(records)
	fmt.Println("Task has been marked completed")
}

func updateCSV(records [][]string) {
	file, _ := os.Create("/home/mklno/projects/tasks/tasks.csv")
	defer file.Close()
	csvWriter := csv.NewWriter(file)
	csvWriter.WriteAll(records)
}
func init() {
	rootCmd.AddCommand(completeTasksCmd)
	completeTasksCmd.Flags().StringP("task-id", "t", "0", "Enter task ID to mark complete")
}

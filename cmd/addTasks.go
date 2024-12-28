package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var addTaskCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task",
	Long: `Add a task to the task list.
	For example:
	
	task add -a <task>
	task add -a "Clean my desk"`,
	Run: addTasks,
}

func addTasks(cmd *cobra.Command, args []string) {
	task, _ := cmd.Flags().GetString("add-task")
	records := readCsvFile("/home/mklno/projects/tasks/tasks.csv")
	if task == "dummy" {
		return
	}
	taskID, _ := strconv.Atoi(records[len(records)-1][0])
	ID := strconv.Itoa(taskID + 1)
	createdTime := time.Now().Format(time.RFC3339)
	taskDone := "false"
	newRecord := []string{ID, task, createdTime, taskDone}
	fmt.Println(newRecord)
	records = append(records, newRecord)
	updateCSV(records)
}

func init() {
	rootCmd.AddCommand(addTaskCmd)
	addTaskCmd.Flags().StringP("add-task", "a", "dummy", "Enter a task")
}

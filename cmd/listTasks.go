package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var listTasksCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks.",
	Long: `List all tasks.
For example:

tasks list => List all pending tasks.
tasks list -a => List all tasks.`,
	Run: listTasks,
}

func listTasks(cmd *cobra.Command, args []string) {
	isAllTasks, _ := cmd.Flags().GetBool("all-tasks")
	records := readCsvFile("/home/mklno/projects/tasks/tasks.csv")
	if isAllTasks {
		getAllTasks(records)
	} else {
		getPendingTasks(records)
	}
}

func readCsvFile(filepath string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("unable to read input file"+filepath, err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("unable to parse csv file"+filepath, err)
	}
	return records
}

func formatRecords(records [][]string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	layout := "2006-01-02T15:04:05-07:00"
	for _, outer := range records {
		var content string
		for j, value := range outer {
			if j == 2 && value != "Created" {
				timestamp, _ := time.Parse(layout, value)
				content += timediff.TimeDiff(timestamp)
			} else {
				content += value
			}
			content += "\t"
		}
		fmt.Fprintln(w, content)
	}
	w.Flush()
}

func getAllTasks(records [][]string) {
	formatRecords(records)
}
func getPendingTasks(records [][]string) {
	if len(records) == 0 {
		fmt.Println("No pending tasks")
		return
	}
	header := records[0]
	isCompleteIndex := -1

	for i, col := range header {
		if col == "Done" {
			isCompleteIndex = i
			break
		}
	}
	if isCompleteIndex == -1 {
		fmt.Println("No 'IsComplete' column found")
		return
	}
	var pendingTasks [][]string
	pendingTasks = append(pendingTasks, header[0:3])

	for _, row := range records[1:] {
		if len(row) <= isCompleteIndex {
			continue
		}
		if row[isCompleteIndex] == "false" {
			pendingTasks = append(pendingTasks, row[0:3])
		}
	}
	formatRecords(pendingTasks)
}
func init() {
	rootCmd.AddCommand(listTasksCmd)
	listTasksCmd.Flags().BoolP("all-tasks", "a", false, "List all tasks")
}

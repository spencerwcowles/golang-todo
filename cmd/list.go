package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mergestat/timediff"
	"io"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var showAll bool

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"all"},
	Short:   "List tasks",
	Long:    "List tasks, use (--all, -a) flag to show all tasks including completed ones",
	Run: func(cmd *cobra.Command, args []string) {
		// Open the jsonFile
		// TODO: maybe change this from json to a basic sql database? like mysql
		path := "todo.json"
		jsonFile, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer jsonFile.Close()

		byteValue, _ := io.ReadAll(jsonFile)
		var jsonData JsonFile
		json.Unmarshal(byteValue, &jsonData)

		// use tabwriter to print out the tasks with a good amount of spacing
		w := tabwriter.NewWriter(os.Stdout, 0, 2, 4, ' ', 0)

		// show all flag, prints all tasks and extra data
		// shows task id, name, completed
		if showAll {
			fmt.Fprintln(w, "ID\tTask\tCreated\tCompleted\t")
			for i := 0; i < len(jsonData.Tasks); i++ {
				timeDiff := timediff.TimeDiff(jsonData.Tasks[i].Time)
				fmt.Fprintf(w, "%d\t%s\t%s\t%t\t\n", jsonData.Tasks[i].Id, jsonData.Tasks[i].Title, timeDiff, jsonData.Tasks[i].Completed)
			}
		} else {
			// show tasks normally
			// shows task id, name
			fmt.Fprintln(w, "ID\tTask\tCreated")
			for i := 0; i < len(jsonData.Tasks); i++ {
				timeDiff := timediff.TimeDiff(jsonData.Tasks[i].Time)
				if jsonData.Tasks[i].Completed && !showAll {
					continue
				}
				fmt.Fprintf(w, "%d\t%s\t%s\n", jsonData.Tasks[i].Id, jsonData.Tasks[i].Title, timeDiff)
			}
		}

		// get rid of the tabwriter to make sure it prints correctly
		w.Flush()

	},
}

func init() {
	// '--all' flag to show all tasks including complete tasks
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all tasks")
	rootCmd.AddCommand(listCmd)
}

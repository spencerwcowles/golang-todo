package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var showAll bool

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"all"},
	Short:   "list all tasks",
	Long:    "list all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		// Open the jsonFile
		// TODO: maybe change this from json to a basic sql database? like mysql
		jsonFile, err := os.Open("todo.json")
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
		// TODO: add time information
		// github.com/mergestat/timediff for displaying relative friendly time differences (1 hour ago, 10 minutes ago, etc)

		if showAll {
			fmt.Fprintln(w, "ID\tTask\tCompleted\t")
			for i := 0; i < len(jsonData.Tasks); i++ {
				fmt.Fprintf(w, "%d\t%s\t%t\t\n", jsonData.Tasks[i].Id, jsonData.Tasks[i].Title, jsonData.Tasks[i].Completed)
			}
		} else {
			// show tasks normally
			// shows task id, name
			fmt.Fprintln(w, "ID\tTask")
			for i := 0; i < len(jsonData.Tasks); i++ {
				if jsonData.Tasks[i].Completed && !showAll {
					continue
				}
				fmt.Fprintf(w, "%d\t%s\t\n", jsonData.Tasks[i].Id, jsonData.Tasks[i].Title)
			}
		}

		// get rid of the tabwriter to make sure it prints correctly
		w.Flush()

	},
}

func init() {
	// add '--all' flag
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all tasks")
	rootCmd.AddCommand(listCmd)
}

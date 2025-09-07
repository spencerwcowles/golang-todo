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
		// Open our jsonFile
		jsonFile, err := os.Open("todo.json")
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
			return
		}
		defer jsonFile.Close()

		byteValue, _ := io.ReadAll(jsonFile)
		var jsonData JsonFile
		json.Unmarshal(byteValue, &jsonData)
		// fmt.Println("-------------------------------")
		// fmt.Println("Total Number of tasks: ", len(jsonData.Tasks))
		// fmt.Println("-------------------------------")
		// for i := 0; i < len(jsonData.Tasks); i++ {
		// 	if jsonData.Tasks[i].Completed && !showAll {
		// 		continue
		// 	}
		// 	fmt.Printf("Task ID: %d | Task Title: %s | Completed: %t\n", jsonData.Tasks[i].Id, jsonData.Tasks[i].Title, jsonData.Tasks[i].Completed)
		// }

		w := tabwriter.NewWriter(os.Stdout, 0, 2, 4, ' ', 0)

		if showAll {

			fmt.Fprintln(w, "ID\tTask\tCompleted\t")
			for i := 0; i < len(jsonData.Tasks); i++ {
				fmt.Fprintf(w, "%d\t%s\t%t\t\n", jsonData.Tasks[i].Id, jsonData.Tasks[i].Title, jsonData.Tasks[i].Completed)
			}

		} else {
			fmt.Fprintln(w, "ID\tTask")
			for i := 0; i < len(jsonData.Tasks); i++ {
				if jsonData.Tasks[i].Completed && !showAll {
					continue
				}
				fmt.Fprintf(w, "%d\t%s\t\n", jsonData.Tasks[i].Id, jsonData.Tasks[i].Title)
			}
		}

		w.Flush()

	},
}

func init() {
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all tasks")
	rootCmd.AddCommand(listCmd)
}

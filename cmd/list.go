package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

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
		// fmt.Println("Successfully Opened todo.json")
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()
		// parse json file to list all tasks

		byteValue, _ := io.ReadAll(jsonFile)
		var jsonData JsonFile
		json.Unmarshal(byteValue, &jsonData)
		fmt.Println("-------------------------------")
		fmt.Println("Total Number of tasks: ", len(jsonData.Tasks))
		fmt.Println("-------------------------------")
		for i := 0; i < len(jsonData.Tasks); i++ {
			if jsonData.Tasks[i].Completed && !showAll {
				continue
			}
			fmt.Printf("Task ID: %d | Task Title: %s | Completed: %t\n", jsonData.Tasks[i].Id, jsonData.Tasks[i].Title, jsonData.Tasks[i].Completed)
		}

	},
}

func init() {
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all tasks")
	rootCmd.AddCommand(listCmd)
}

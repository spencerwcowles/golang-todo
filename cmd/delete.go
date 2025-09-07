package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "remove", "rm", "rem"},
	Short:   "Delete a task",
	Long:    "Delete a task using a task ID found by listing the current tasks.",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := "todo.json"
		// Open our jsonFile
		jsonFile, err := os.Open(path)
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Println("Successfully Opened todo.json")
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		byteValue, _ := io.ReadAll(jsonFile)
		var jsonData JsonFile
		json.Unmarshal(byteValue, &jsonData)

		taskId := args[0]
		var taskFound bool = false

		for i := 0; i < len(jsonData.Tasks); i++ {
			if fmt.Sprintf("%d", jsonData.Tasks[i].Id) == taskId {
				jsonData.Tasks = append(jsonData.Tasks[:i], jsonData.Tasks[i+1:]...)
				taskFound = true
				break
			}
		}
		if !taskFound {
			fmt.Println("Task not found")
			return
		}

		// Marshal updated data
		updatedData, err := json.MarshalIndent(jsonData, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}

		// Write back to file
		err = os.WriteFile("todo.json", updatedData, 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		fmt.Println("-------------------------------")
		fmt.Printf("Task %v deleted successfully!\n", taskId)
		fmt.Println("-------------------------------")

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

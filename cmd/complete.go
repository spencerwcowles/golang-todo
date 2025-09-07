package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:     "complete",
	Aliases: []string{"done"},
	Short:   "complete a task",
	Long:    "complete a task",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		jsonFile, err := os.Open("todo.json")
		if err != nil {
			fmt.Println(err)
			return
		}

		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		byteValue, _ := io.ReadAll(jsonFile)
		var jsonData JsonFile
		json.Unmarshal(byteValue, &jsonData)

		taskId := args[0]
		var taskFound bool = false

		for i := 0; i < len(jsonData.Tasks); i++ {
			if fmt.Sprintf("%d", jsonData.Tasks[i].Id) == taskId {
				jsonData.Tasks[i].Completed = true
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
		fmt.Printf("Task %v completed successfully!\n", taskId)
		fmt.Println("-------------------------------")

	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}

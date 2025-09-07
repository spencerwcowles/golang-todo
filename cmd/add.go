package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"new"},
	Short:   "add a task short description", // TODO: add documentation ( to all of the commands)
	Long:    "add a task long description",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Open our jsonFile
		jsonFile, err := os.Open("todo.json")
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
			return
		}
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		byteValue, _ := io.ReadAll(jsonFile)
		var jsonData JsonFile
		json.Unmarshal(byteValue, &jsonData)
		newTask := &Task{
			Id:        jsonData.NextId,
			Title:     args[0],
			Completed: false,
		}
		jsonData.Tasks = append(jsonData.Tasks, *newTask)
		jsonData.NextId += 1

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
		fmt.Printf("Task %v added successfully!\n", newTask.Id)
		fmt.Println("-------------------------------")

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

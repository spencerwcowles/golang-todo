package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type JsonFile struct {
	Tasks  []Task `json:"tasks"`
	NextId int    `json:"nextId"`
}

type Task struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "simple todo app to learn golang",
	Long:  "using golang to make a simple todo cli app",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing todo '%s'\n", err)
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"
	"os"

	"github.com/darleneeon/todo/database"
	"github.com/darleneeon/todo/model"
	"github.com/spf13/cobra"
)

// newCmd represents the add command
var newCmd = &cobra.Command{
	Use:   "add [TEXT]",
	Short: "Add new task",
	Long:  `Add (todo add) will add new task to the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Usage: todo add [TEXT]")
			return
		}

		task := model.Task{Text: args[0]}
		if err := database.AddTask(&task); err != nil {
			fmt.Printf("Error while saving task: %s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(newCmd)
}

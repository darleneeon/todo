package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/darleneeon/todo/database"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done [ID]",
	Short: "Mark task as done",
	Long:  `done (todo done) will remove task.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Usage: todo done [ID]")
			return
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("[ID] argument must be a number!")
			os.Exit(1)
		}

		if err := database.DeleteTask(id); err != nil {
			fmt.Printf("Error while deleting task: %s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(doneCmd)
}

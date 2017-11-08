package cmd

import (
	"fmt"
	"os"

	"github.com/darleneeon/todo/database"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show all available tasks",
	Long:  `Show (todo show) will print list of available tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("No arguments required!")
			return
		}

		tasks, err := database.GetTasks()
		if err != nil {
			fmt.Printf("Error while getting tasks: %s\n", err)
			os.Exit(1)
		}

		for i := 0; i < len(tasks); i++ {
			task := tasks[i]
			fmt.Printf("[%d]\t %s\n", task.ID, task.Text)
		}
	},
}

func init() {
	RootCmd.AddCommand(showCmd)
}

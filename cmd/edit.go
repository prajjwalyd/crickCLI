package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"github.com/prajjwalyd/crickCLI/internal/config"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the stored match ID",
	Run: func(cmd *cobra.Command, args []string) {
		var newMatchID string
		fmt.Print("Enter a new match ID: ")
		fmt.Scanln(&newMatchID)

		err := config.StoreMatchID(newMatchID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Match ID updated successfully.")
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}

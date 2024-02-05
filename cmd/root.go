package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"github.com/prajjwalyd/crickCLI/internal/config"
	"github.com/prajjwalyd/crickCLI/internal/cricklib"
	"github.com/prajjwalyd/crickCLI/pkg/http"
)

var matchID string

var rootCmd = &cobra.Command{
	Use:   "crickcli",
	Short: "A CLI application for live cricket scores",
	Run: func(cmd *cobra.Command, args []string) {
		if matchID == "" {
			storedMatchID, err := config.ReadStoredMatchID()
			if err != nil {
				log.Fatal(err)
			}
			matchID = storedMatchID
		}

		if matchID == "" {
			log.Fatal("Match ID not set. Use 'crickcli edit' to set a match ID.")
		}

		url := fmt.Sprintf("https://cric-api-orpin.vercel.app/score?id=%s", matchID)
		response, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		matchData, err := cricklib.ParseResponse(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		cricklib.PrintLiveScore(matchData)
	},
}

// Execute initializes the CLI application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

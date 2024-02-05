// main.go

package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var matchID string

// Config file path
var configFilePath = getConfigFilePath()

// MatchData struct to represent the API response
type MatchData struct {
	Title           string `json:"title"`
	Update          string `json:"update"`
	LiveScore       string `json:"livescore"`
	RunRate         string `json:"runrate"`
	BatsmanOne      string `json:"batterone"`
	BatsmanOneRun   string `json:"batsmanonerun"`
	BatsmanOneBall  string `json:"batsmanoneball"`
	BatsmanOneSR    string `json:"batsmanonesr"`
	BatsmanTwo      string `json:"battertwo"`
	BatsmanTwoRun   string `json:"batsmantworun"`
	BatsmanTwoBall  string `json:"batsmantwoball"`
	BatsmanTwoSR    string `json:"batsmantwosr"`
	BowlerOne       string `json:"bowlerone"`
	BowlerOneOver   string `json:"bowleroneover"`
	BowlerOneRun    string `json:"bowleronerun"`
	BowlerOneWickets string `json:"bowleronewickers"`
	BowlerOneEcon   string `json:"bowleroneeconomy"`
	BowlerTwo       string `json:"bowlertwo"`
	BowlerTwoOver   string `json:"bowlertwoover"`
	BowlerTwoRun    string `json:"bowlertworun"`
	BowlerTwoWickets string `json:"bowlertwowickers"`
	BowlerTwoEcon   string `json:"bowlertwoeconomy"`
}

var rootCmd = &cobra.Command{
	Use:   "crickcli",
	Short: "A CLI application for live cricket scores",
	Run: func(cmd *cobra.Command, args []string) {
		if matchID == "" {
			// If match ID is not provided as a flag, check the config file
			storedMatchID, err := readStoredMatchID()
			if err != nil {
				log.Fatal(err)
			}
			matchID = storedMatchID
		}

		if matchID == "" {
			log.Fatal("Match ID not set. Use 'crickcli edit' to set a match ID.")
		}

		// Fetch live cricket score based on match ID
		url := fmt.Sprintf("https://cric-api-orpin.vercel.app/score?id=%s", matchID)
		response, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		// Parse the response
		var matchData MatchData
		decoder := json.NewDecoder(response.Body)
		err = decoder.Decode(&matchData)
		if err != nil {
			log.Fatal(err)
		}

		// Display the live cricket score with color formatting
		printLiveScore(matchData)
	},
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the stored match ID",
	Run: func(cmd *cobra.Command, args []string) {
		var newMatchID string
		fmt.Print("Enter a new match ID: ")
		fmt.Scanln(&newMatchID)

		err := storeMatchID(newMatchID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Match ID updated successfully.")
	},
}

func printLiveScore(matchData MatchData) {
	// Color formatting for better visualization
	title := color.New(color.Bold, color.FgHiCyan).Sprint(matchData.Title)
	update := color.New(color.FgHiYellow).Sprint(matchData.Update)
	liveScore := color.New(color.Bold, color.FgHiGreen).Sprint(matchData.LiveScore)
	runRate := color.New(color.FgHiMagenta).Sprint(matchData.RunRate)

	// Batsman details
	batsmanDetails := fmt.Sprintf("Batting: \n %s %s%s SR:%s\n %s %s%s SR:%s\n",
		matchData.BatsmanOne, matchData.BatsmanOneRun, matchData.BatsmanOneBall, matchData.BatsmanOneSR,
		matchData.BatsmanTwo, matchData.BatsmanTwoRun, matchData.BatsmanTwoBall, matchData.BatsmanTwoSR)

	// Bowler details
	bowlerDetails := fmt.Sprintf("Bowling: \n %s %s-%s (%s)\n %s %s-%s (%s)\n",
		matchData.BowlerOne, matchData.BowlerOneWickets, matchData.BowlerOneRun, matchData.BowlerOneOver,
		matchData.BowlerTwo, matchData.BowlerTwoWickets, matchData.BowlerTwoRun, matchData.BowlerTwoOver)

	// Print formatted live cricket score with batsman and bowler details
	fmt.Printf("\n%s\n%s\n%s\n%s\n\n%s\n%s", title, update, liveScore, runRate, batsmanDetails, bowlerDetails)
}



func readStoredMatchID() (string, error) {
	// Read match ID from the config file
	content, err := os.ReadFile(configFilePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func storeMatchID(newMatchID string) error {
	// Store the new match ID in the config file
	return os.WriteFile(configFilePath, []byte(newMatchID), 0644)
}

func getConfigFilePath() string {
	// Get the user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// Create the config directory if it doesn't exist
	configDir := filepath.Join(home, ".crickcli")
	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Return the path to the config file
	return filepath.Join(configDir, "config.txt")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

func init() {
	// Define flags for the command line
	rootCmd.Flags().StringVarP(&matchID, "matchID", "m", "", "Match ID for live cricket score")

	// Add 'edit' command to the root command
	rootCmd.AddCommand(editCmd)
}

func main() {
	Execute()
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

var matchID string

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

func printLiveScore(matchData MatchData) {
	// Color formatting for better visualization
	title := color.New(color.Bold, color.FgHiCyan).Sprint(matchData.Title)
	update := color.New(color.FgHiYellow).Sprint(matchData.Update)
	liveScore := color.New(color.Bold, color.FgHiGreen).Sprint(matchData.LiveScore)
	runRate := color.New(color.FgHiMagenta).Sprint(matchData.RunRate)

	// Print formatted live cricket score
	fmt.Printf("%s\n%s\n%s\n%s\n", title, update, liveScore, runRate)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

func init() {
	// Define flags for the command line
	rootCmd.Flags().StringVarP(&matchID, "matchID", "m", "", "Match ID for live cricket score (required)")

	// Mark the flag as required
	_ = rootCmd.MarkFlagRequired("matchID")
}

func main() {
	Execute()
}

package cricklib

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io"
)

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

// ParseResponse decodes the JSON response from the API
func ParseResponse(body io.Reader) (MatchData, error) {
	var matchData MatchData
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&matchData)
	return matchData, err
}

// PrintLiveScore displays the live cricket score with color formatting
func PrintLiveScore(matchData MatchData) {
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
	bowlerDetails := fmt.Sprintf("Bowling: \n %s %s-%s (%s)\n %s %s-%s (%s)\n\n",
		matchData.BowlerOne, matchData.BowlerOneWickets, matchData.BowlerOneRun, matchData.BowlerOneOver,
		matchData.BowlerTwo, matchData.BowlerTwoWickets, matchData.BowlerTwoRun, matchData.BowlerTwoOver)

	// Print formatted live cricket score with batsman and bowler details
	fmt.Printf("\n%s\n%s\n%s\n%s\n\n%s\n%s", title, update, liveScore, runRate, batsmanDetails, bowlerDetails)
}


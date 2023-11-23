/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// currentTimeCmd represents the currentTime command
var currentTimeCmd = &cobra.Command{
	Use:   "currentTime",
	Short: "List current time based on timezone",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		CurrentTime()
	},
}

func init() {
	rootCmd.AddCommand(currentTimeCmd)
}

type Time struct {
	Time     string `json:"time"`
	Timezone string `json:"timezone"`
}

func CurrentTime() {

	API := "https://timeapi.io/api/Time/current/zone?timeZone="
	var tz string

	fmt.Print("Enter Timezone: ")

	fmt.Scan(&tz)

	CAPI := API + tz

	resp, err := http.Get(CAPI)
	if err != nil {
		error := fmt.Errorf("Kindly check the entered value")
		fmt.Println(error.Error())
		os.Exit(1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		error := fmt.Errorf("API Doesn't exist. kindly verify the timezone using list option")
		fmt.Println(error.Error())
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		error := fmt.Errorf("Something gone wrong, Waiting for sometime & then re-try.")
		fmt.Println(error.Error())
		os.Exit(1)
	}

	var timex Time
	err = json.Unmarshal(body, &timex)

	if err != nil {
		error := fmt.Errorf("Something gone wrong, Waiting for sometime & then re-try.")
		fmt.Println(error.Error())
		os.Exit(1)
	}

	fmt.Printf("The current time in %v is %v\n", timex.Timezone, timex.Time)

}

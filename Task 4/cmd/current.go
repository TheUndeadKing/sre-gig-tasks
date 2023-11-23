package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// currentCmd represents the current command
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "List current time based on timezone",
	Long: `List current time based on timezone. 
Syntax: ctime current [timezone]
Example: ctime current Asia/Tokyo`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			for _, zone := range args {
				CurrentTime(zone)
			}
		} else {
			fmt.Println("Kindly provide a valid timezone\nSyntax: ctime current [timezone]\nExample: ctime current Asia/Tokyo")
		}
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)

}

type Time struct {
	Time     string `json:"time"`
	Timezone string `json:"timezone"`
}

func CurrentTime(zone string) {

	API := "https://timeapi.io/api/Time/current/zone?timeZone=" + zone

	resp, err := http.Get(API)
	if err != nil {
		log.Println("Unable to get response")
		os.Exit(1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("Alert!!. check if you access 'https://worldtimeapi.org' site & check provided timezone is a valid entry using list option. Current HTTP code - %v\n", resp.StatusCode)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("Something gone wrong, please retry")
		os.Exit(1)
	}

	var timex Time
	err = json.Unmarshal(body, &timex)

	if err != nil {
		log.Println("Unable to unmarshal the json response.")
		os.Exit(1)
	}

	fmt.Printf("The current time in %v is %v\n", timex.Timezone, timex.Time)

}

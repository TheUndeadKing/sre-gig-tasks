/*
Copyright © 2023 Anwar
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all timezone",
	Long: `With the list paramter you can list all timezone
Syntax: ctime list`,
	Run: func(cmd *cobra.Command, args []string) {
		List()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}

func List() {
	resp, err := http.Get("http://worldtimeapi.org/api/timezone.txt")
	if err != nil {
		log.Println("Unable to get response")
		os.Exit(1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("Alert!!. check if you access 'https://worldtimeapi.org' site. Current HTTP code - %v\n", resp.StatusCode)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("Something gone wrong, please retry")
		os.Exit(1)
	}

	fmt.Println(string(body))

}

/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all timezone",
	Long:  `With the list paramter you can list all timezone`,
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
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("Api not avaiable")
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

}

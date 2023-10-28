/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert date &time between timezones",
	Long: ` Convert date time from one timezones to another timezones.
	EXAMPLE:
	========
	Enter from timezone:Europe/Amsterdam
    Enter from time:2021-03-14
    Enter from date:17:45:00
    Enter to timezone:America/Los_Angeles
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ConvertTime()
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

}

type Post struct {
	ConversionResult struct {
		Time     string `json:"time"`
		Date     string `json:"date"`
		TimeZone string `json: "timeZone"`
	} `json:"conversionResult"`
	FromTimezone string `json:"fromTimezone"`
}

func ConvertTime() {

	API := "https://timeapi.io/api/Conversion/ConvertTimeZone"

	var a, b, c, d string

	fmt.Print("Enter from timezone:")
	fmt.Scan(&a)
	fmt.Print("Enter from time:")
	fmt.Scan(&b)
	fmt.Print("Enter from date:")
	fmt.Scan(&c)
	fmt.Print("Enter to timezone:")
	fmt.Scan(&d)

	jsonData := []byte(fmt.Sprintf(`{
				"fromTimeZone":  "%v",
				"dateTime": "%v %v",
				"toTimeZone": "%v",
				"dstAmbiguity": ""
			}`, a, b, c, d))

	request, error := http.NewRequest("POST", API, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	var post Post
	json.Unmarshal(body, &post)

	fmt.Printf("Time - %v \nDate - %v \nTimezone - %v \n", post.ConversionResult.Time, post.ConversionResult.Date, post.ConversionResult.TimeZone)
}

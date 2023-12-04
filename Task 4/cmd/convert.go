package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert date & time from timezones to another timezones",
	Long:  `Convert date time from one timezones to another timezones.`,
	Run: func(cmd *cobra.Command, args []string) {
		ConvertTime()
	},
}

var From string
var To string
var FromTime string
var Date string

func init() {
	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().StringVarP(&From, "from", "f", "default", "from (required)")
	convertCmd.Flags().StringVarP(&To, "to", "t", "default", "to (required)")
	convertCmd.Flags().StringVarP(&FromTime, "fromtime", "e", "default", "time (required)")
	convertCmd.Flags().StringVarP(&Date, "date", "d", "default", "date (required)")

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

	jsonData := []byte(fmt.Sprintf(`{
				"fromTimeZone":  "%v",
				"dateTime": "%v %v",
				"toTimeZone": "%v",
				"dstAmbiguity": ""
			}`, From, Date, FromTime, To))

	request, error := http.NewRequest("POST", API, bytes.NewBuffer(jsonData))
	if error != nil {
		log.Println("Unable to Perform POST Requst")
		os.Exit(1)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		log.Println("Unable get response")
		os.Exit(1)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Unable get response")
		os.Exit(1)
	}

	var post Post
	err = json.Unmarshal(body, &post)

	if err != nil {
		log.Println("Unable get unmarhsal json. Please check the input argument once.")
		os.Exit(1)
	}

	fmt.Printf("Time - %v \nDate - %v \nTimezone - %v \n", post.ConversionResult.Time, post.ConversionResult.Date, post.ConversionResult.TimeZone)
}

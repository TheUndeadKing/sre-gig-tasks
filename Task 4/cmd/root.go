package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ctime",
		Short: "time based Applications",
		Long: `ctime is a CLI library. It can be used to convert time & date based on timezone.
Get current time for all timezone.`,
		Run: func(cmd *cobra.Command, args []string) {
			Info := `ctime v1.0!
ctime is a CLI library. It can be used to convert time & date based on timezone.
Get current time for all timezone.
		
Usage:
ctime [command]
		
Commands:
convert     Convert date & time from timezones to another timezones
current     List current time based on timezone
help        Help about any command
list        List all timezone`
			fmt.Println(Info)
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

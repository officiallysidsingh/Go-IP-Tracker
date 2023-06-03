package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP Address",
	Long: `Trace the IP Address and get the Location of the IP Address.
			Example:
					iptracker trace 1111.1111.1111.1111
			`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				fmt.Println(ip)
			}
		} else {
			fmt.Println("Please provide the IP Address")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

func showData() {}

func getData() {
	url := "http://ipinfo.io/1.1.1.1/geo"

	http.Get(url)
}

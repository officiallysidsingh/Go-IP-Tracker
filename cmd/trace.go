package cmd

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/fatih/color"
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
				showData(ip)
			}
		} else {
			red := color.New(color.FgRed).Add(color.Bold)
			red.Println("Please provide the IP Address")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

type IPAddress struct {
	IP			string		`json:"ip"`
	City		string		`json:"city"`
	Region		string		`json:"region"`
	Country		string		`json:"country"`
	Loc			string		`json:"loc"`
	Postal		string		`json:"postal"`
	Timezone	string		`json:"timezone"`
	Org			string		`json:"org"`
}

func showData(ip string) {
	url := "http://ipinfo.io/"+ ip +"/geo"
	resByte := getData(url)

	var data IPAddress

	err := json.Unmarshal(resByte, &data)
	if err != nil {
		red := color.New(color.FgRed).Add(color.Bold)
		red.Println("Unable to Unmarshal the data")
	}
	
	dataFound := color.New(color.FgGreen).Add(color.Underline).Add(color.Bold)
	dataFound.Printf("\nData Found")
	nextline := color.New(color.FgGreen).Add(color.Bold)
	nextline.Printf(" :- \n\n")

	results := color.New(color.FgMagenta).Add(color.Bold)
	results.Printf("IP :\t\t%s\nCITY:\t\t%s\nREGION:\t\t%s\nCOUNTRY:\t%s\nLOC:\t\t%s\nPOSTAL:\t\t%s\nTIMEZONE:\t%s\nORG:\t\t%s\n", 
				data.IP, 
				data.City, 
				data.Region, 
				data.Country, 
				data.Loc, 
				data.Postal, 
				data.Timezone, 
				data.Org)
}

func getData(url string) []byte {

	res, err := http.Get(url)

	if err != nil {
		red := color.New(color.FgRed).Add(color.Bold)
		red.Println("Unable to get the response")
	}

	resByte, err := io.ReadAll(res.Body)

	if err != nil {
		red := color.New(color.FgRed).Add(color.Bold)
		red.Println("Unable to read the response")
	}

	return resByte
}

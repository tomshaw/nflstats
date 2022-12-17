package parse

import (
	"flag"
	"github.com/tomshaw/nflstats/common"
)

func Flag(Info *common.InputInfoStruct) {
	flag.StringVar(&Info.ApiName, "api", "NewsByTeam", "The API endpoint to query.")

	// TeamGameStats
	flag.IntVar(&Info.SeasonYear, "year", 2022, "The NFL season year.")
	flag.IntVar(&Info.SeasonWeek, "week", 1, "The NFL season week. 1-17 18 week period")
	flag.StringVar(&Info.SeasonType, "type", "REG", "The NFL season type. reg/pre/post")

	// NewsByTeam
	flag.StringVar(&Info.Team, "team", "DAL", "The NFL team abbrevation.")

	// Export/Print/Table
	flag.BoolVar(&Info.Export, "export", false, "Export results to JSON file.")
	flag.BoolVar(&Info.Print, "print", false, "Pretty prints results to console output.")
	flag.BoolVar(&Info.Table, "table", false, "Generates results to console table output.")

	flag.BoolVar(&Info.Color, "color", true, "Colorize output messages.")
	flag.BoolVar(&Info.DryRun, "dryrun", false, "Prints running options only.")
	flag.BoolVar(&Info.Silent, "silent", false, "Print results to screen.")
	flag.BoolVar(&Info.List, "list", false, "Prints available API's to console output.")

	flag.BoolVar(&Info.Help, "help", false, "Display command line help.")

	flag.Parse()
}

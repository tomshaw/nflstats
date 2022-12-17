package parse

import (
	"github.com/tomshaw/nflstats/common"
)

func Parse(InputInfo *common.InputInfoStruct, RunningInfo *common.RunningInfoStruct) {
	RunningInfo.ApiName = InputInfo.ApiName

	// TeamGameStats
	RunningInfo.SeasonYear = InputInfo.SeasonYear
	RunningInfo.SeasonWeek = InputInfo.SeasonWeek
	RunningInfo.SeasonType = InputInfo.SeasonType // REG/PRE/POST

	// NewsByTeam
	RunningInfo.Team = InputInfo.Team // DAL/ARI

	// Output
	RunningInfo.Export = InputInfo.Export
	RunningInfo.Print = InputInfo.Print
	RunningInfo.Table = InputInfo.Table

	RunningInfo.Color = InputInfo.Color
	RunningInfo.DryRun = InputInfo.DryRun
	RunningInfo.Silent = InputInfo.Silent
	RunningInfo.List = InputInfo.List

	RunningInfo.Help = InputInfo.Help
}

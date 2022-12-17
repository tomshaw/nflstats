package api

import (
	"github.com/tomshaw/nflstats/common"
)

type (
	Endpoint interface {
		Request(apiKey string)
	}
)

func Factory(RunningInfo *common.RunningInfoStruct) Endpoint {
	switch common.InputInfo.ApiName {
	case "TeamGameStats":
		return TeamGameStats{
			SeasonYear: common.InputInfo.SeasonYear,
			SeasonWeek: common.InputInfo.SeasonWeek,
			SeasonType: common.InputInfo.SeasonType,
		}
	case "NewsByTeam":
		return NewsByTeam{
			Team: common.InputInfo.Team,
		}
	default:
		return nil
	}
}

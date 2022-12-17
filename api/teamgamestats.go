package api

import (
	"encoding/json"
	"github.com/tomshaw/nflstats/api/models"
	"github.com/tomshaw/nflstats/common"
	"github.com/tomshaw/nflstats/utils/logger"
	"github.com/tomshaw/nflstats/utils/printer"
	"github.com/tomshaw/nflstats/utils/request"
	"github.com/tomshaw/nflstats/utils/writer"
	"strconv"
)

type TeamGameStats struct {
	SeasonYear int
	SeasonWeek int
	SeasonType string
}

func (api TeamGameStats) Request(apiKey string) {
	url := "https://api.sportsdata.io/v3/nfl/scores/json/TeamGameStats/" + strconv.Itoa(api.SeasonYear) + api.SeasonType + "/" + strconv.Itoa(api.SeasonWeek) + "?key=" + apiKey

	if common.RunningInfo.DryRun {
		logger.General(0, url, "URL")
	}

	resp := request.SendRequest(url) //nolint

	if common.RunningInfo.Export {
		writer.Export(resp, "data/TeamGameStats.json")
	}

	if common.RunningInfo.Silent {
		return
	}

	var results []models.TeamGameStatsResponse

	if common.RunningInfo.Print {
		printer.Print(resp, results)
	}

	if common.RunningInfo.Table {
		if err := json.Unmarshal(resp, &results); err != nil {
			logger.Fatal(err.Error())
		}

		printer.TeamGameStatsPrinter(results, "NFL Game Stats")
	}
}

package api

import (
	"encoding/json"
	"github.com/tomshaw/nflstats/api/models"
	"github.com/tomshaw/nflstats/common"
	"github.com/tomshaw/nflstats/utils/logger"
	"github.com/tomshaw/nflstats/utils/printer"
	"github.com/tomshaw/nflstats/utils/request"
	"github.com/tomshaw/nflstats/utils/writer"
)

type NewsByTeam struct {
	Team string
}

func (api NewsByTeam) Request(apiKey string) {
	url := "https://api.sportsdata.io/v3/nfl/scores/json/NewsByTeam/" + api.Team + "?key=" + apiKey

	if common.RunningInfo.DryRun {
		logger.General(0, url, "URL")
	}

	resp := request.SendRequest(url) //nolint

	if common.RunningInfo.Export {
		writer.Export(resp, "data/NewsByTeam.json")
	}

	if common.RunningInfo.Silent {
		return
	}

	var results []models.NewsByTeamResponse

	if common.RunningInfo.Print {
		printer.Print(resp, results)
	}

	if common.RunningInfo.Table {
		if err := json.Unmarshal(resp, &results); err != nil {
			logger.Fatal(err.Error())
		}

		printer.NewByTeamPrinter(results, "News By Team")
	}
}

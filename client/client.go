package client

import (
	"github.com/tomshaw/nflstats/common"
	"github.com/tomshaw/nflstats/utils"
	"github.com/tomshaw/nflstats/utils/logger"
	"github.com/tomshaw/nflstats/utils/printer"
	"os"
)

func Options() {
	logger.Info(logger.Blue("Starting application..."))

	api := common.RunningInfo.ApiName

	if common.RunningInfo.Help {
		logger.Info(logger.LightGreen(">> Display help screen"))
		os.Exit(1)
	}

	if _, ok := common.NFLApis[api]; ok {
		logger.Info(logger.LightGreen("Initializing API: ") + logger.White(api))
	} else {
		logger.Fatal(logger.LightGreen("Selected API does not exist: ") + logger.Red(api))
	}

	apiKey := utils.GetEnv("API_KEY")
	if len(apiKey) != 0 {
		common.RunningInfo.ApiKey = apiKey
	}

	if common.RunningInfo.DryRun {
		logger.Info(logger.LightGreen(common.RunningInfo.ApiKey))
	}

	if common.RunningInfo.List {
		printer.PrintNFLApis(common.NFLApis)
	}
}

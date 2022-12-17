package main

import (
	"fmt"
	"github.com/tomshaw/nflstats/api"
	"github.com/tomshaw/nflstats/client"
	"github.com/tomshaw/nflstats/common"
	"github.com/tomshaw/nflstats/parse"
	"github.com/tomshaw/nflstats/utils"
	"github.com/tomshaw/nflstats/utils/logger"
	"time"
)

func main() {
	startTime := time.Now()
	utils.PrintBanner(common.Version)
	parse.Flag(&common.InputInfo)
	parse.Parse(&common.InputInfo, &common.RunningInfo)
	client.Options()
	api := api.Factory(&common.RunningInfo)
	api.Request(common.RunningInfo.ApiKey)
	logger.Info(fmt.Sprintf("Finished execution... consumption time: %s", time.Since(startTime)))
}

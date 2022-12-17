package common

var Version = "0.0.1"

type InputInfoStruct struct {
	ApiName string

	SeasonYear int
	SeasonWeek int
	SeasonType string

	Team string

	Export bool
	Print  bool
	Table  bool

	Color  bool
	DryRun bool
	Silent bool
	List   bool

	Help bool
}

type RunningInfoStruct struct {
	ApiName string
	ApiKey  string

	SeasonYear int
	SeasonWeek int
	SeasonType string

	Team string

	Export bool
	Print  bool
	Table  bool

	Color  bool
	DryRun bool
	Silent bool
	List   bool

	Help bool
}

var InputInfo InputInfoStruct
var RunningInfo RunningInfoStruct

var NFLApis = map[string]string{
	"NewsByTeam":    "https://api.sportsdata.io/v3/nfl/scores/json/NewsByTeam/{team}",
	"TeamGameStats": "https://api.sportsdata.io/v3/nfl/scores/json/TeamGameStats/{season}/{week}",
}

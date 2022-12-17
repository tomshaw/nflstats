package printer

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/tomshaw/nflstats/api/models"
	"os"
)

func PrintNFLApis(apis map[string]string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)
	t.AppendHeader(table.Row{"API", "Endpoint"})
	for key, val := range apis {
		t.AppendRow(table.Row{key, val})
		t.AppendSeparator()
	}
	t.Render()
}

func NewByTeamPrinter(results []models.NewsByTeamResponse, title string) {
	t := table.NewWriter()
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Number:    1,
			AutoMerge: true,
			WidthMin:  100,
			WidthMax:  100,
		},
	})
	t.SetStyle(table.StyleRounded)
	t.SetTitle(title)
	t.Style().Title.Align = text.AlignLeft
	t.Style().Title.Format = text.FormatUpper
	t.SetOutputMirror(os.Stdout)
	t.Style().Options.SeparateRows = true

	for _, v := range results {
		header := fmt.Sprintf("%s - %s", v.Source, v.Updated)
		content := fmt.Sprintf("%s\n%s", v.Title, v.Content)
		footer := fmt.Sprintf("%s Last Updated: %s", v.Author, v.TimeAgo)
		t.AppendHeader(table.Row{header})
		t.AppendRow(table.Row{content})
		t.AppendFooter(table.Row{footer})
		t.AppendSeparator()
	}

	t.Render()
}

func TeamGameStatsPrinter(results []models.TeamGameStatsResponse, title string) {
	t := table.NewWriter()
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Number:      1,
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
		},
		{
			Number:      2,
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
		},
		{
			Number:      3,
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
		},
		{
			Number:      4,
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
		},
		{
			Number:      5,
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
		},
		{
			Number:      6,
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
		},
		{
			Number:      7,
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
		},
		{
			Number:      8,
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
		},
		{
			Number:      9,
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
		},
		{
			Number:      10,
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
		},
		{
			Number:      11,
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
		},
	})
	t.SetStyle(table.StyleRounded)
	t.SetTitle(title)
	t.Style().Title.Align = text.AlignCenter
	t.Style().Title.Format = text.FormatUpper
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Team", "Opponent", "Score", "Touchdowns", "RushingYards", "PassingYards", "Sacks", "Fumbles", "Punts", "Stadium", "Date"})
	t.Style().Options.SeparateRows = true

	for _, v := range results {
		homeTeamScore := v.ScoreQuarter1 + v.ScoreQuarter2 + v.ScoreQuarter3 + v.ScoreQuarter4
		opponentScore := v.OpponentScoreQuarter1 + v.OpponentScoreQuarter2 + v.OpponentScoreQuarter3 + v.OpponentScoreQuarter4
		finalScore := fmt.Sprintf("%d/%d", homeTeamScore, opponentScore)
		t.AppendRow(table.Row{v.Team, v.Opponent, finalScore, v.Touchdowns, v.RushingYards, v.PassingYards, v.Sacks, v.Fumbles, v.Punts, v.Stadium, v.Date})
		t.AppendSeparator()
	}

	t.Render()
}

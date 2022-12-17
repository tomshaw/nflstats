package models

type NewsByTeamResponse struct {
	NewsID            int         `json:"NewsID"`
	Source            string      `json:"Source"`
	Updated           string      `json:"Updated"`
	TimeAgo           string      `json:"TimeAgo"`
	Title             string      `json:"Title"`
	Content           string      `json:"Content"`
	URL               string      `json:"Url"`
	TermsOfUse        string      `json:"TermsOfUse"`
	Author            string      `json:"Author"`
	Categories        string      `json:"Categories"`
	PlayerID          int         `json:"PlayerID"`
	TeamID            int         `json:"TeamID"`
	Team              string      `json:"Team"`
	PlayerID2         interface{} `json:"PlayerID2"`
	TeamID2           interface{} `json:"TeamID2"`
	Team2             interface{} `json:"Team2"`
	OriginalSource    string      `json:"OriginalSource"`
	OriginalSourceURL string      `json:"OriginalSourceUrl"`
}

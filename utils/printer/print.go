package printer

import (
	"encoding/json"
	"fmt"
	"github.com/tomshaw/nflstats/utils/logger"
)

func Print(resp []byte, result interface{}) {
	if err := json.Unmarshal(resp, &result); err != nil {
		logger.Fatal(err.Error())
	}
	fmt.Println(PrettyPrint(result))
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

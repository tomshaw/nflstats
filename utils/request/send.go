package request

import (
	"github.com/tomshaw/nflstats/utils/logger"
	"io"
	"net/http"
)

func SendRequest(url string) (resBody []byte) {
	resp, err := http.Get(url) //nolint
	if err != nil {
		logger.Fatal(err.Error())
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Fatal(err.Error())
	}
	/*if len(print) > 0 {
		fmt.Printf("http response code: %d\n", resp.StatusCode)
		fmt.Printf("http response body: %s\n", respBody)
	}*/
	return respBody
}

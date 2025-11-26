package writer

import (
	"bytes"
	"github.com/tomshaw/nflstats/utils/logger"
	"io"
	"os"
)

func Export(resp []byte, file string) {
	out, err := os.Create(file) // #nosec G304
	if err != nil {
		logger.Fatal(logger.Red("File creation error."))
	}
	defer func() {
		if closeErr := out.Close(); closeErr != nil {
			logger.Error(logger.Red("File close error."))
		}
	}()
	r := bytes.NewReader(resp) // []byte to reader
	_, err = io.Copy(out, r)
	if err != nil {
		logger.Fatal(logger.Red("File creation error."))
	}
}

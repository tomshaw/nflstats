package writer

import (
	"bytes"
	"github.com/tomshaw/nflstats/utils/logger"
	"io"
	"os"
)

func Export(resp []byte, file string) {
	out, err := os.Create(file)
	if err != nil {
		logger.Fatal(logger.Red("File creation error."))
	}
	defer out.Close()
	r := bytes.NewReader(resp) // []byte to reader
	_, err = io.Copy(out, r)
	if err != nil {
		logger.Fatal(logger.Red("File creation error."))
	}
}

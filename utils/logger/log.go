package logger

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/tomshaw/nflstats/common"
	"github.com/tomshaw/nflstats/utils"
	"os"
	"regexp"
)

var (
	Red         = color.Red.Render
	Cyan        = color.Cyan.Render
	Yellow      = color.Yellow.Render
	White       = color.White.Render
	Blue        = color.Blue.Render
	Green       = color.Green.Render
	Purple      = color.Style{color.Magenta, color.OpBold}.Render
	LightRed    = color.Style{color.Red, color.OpBold}.Render
	LightGreen  = color.Style{color.Green, color.OpBold}.Render
	LightWhite  = color.Style{color.White, color.OpBold}.Render
	LightCyan   = color.Style{color.Cyan, color.OpBold}.Render
	LightYellow = color.Style{color.Yellow, color.OpBold}.Render
)

func log(l Level, msg string) {
	if common.RunningInfo.Color {
		fmt.Println(msg)
	} else {
		fmt.Println(StripAnsi(msg))
	}

	if common.RunningInfo.Export {
		WriteLogFile(StripAnsi(msg), "logs.txt")
	}

	if l == LevelFatal {
		os.Exit(0)
	}
}

func Fatal(msg string) {
	log(LevelFatal, fmt.Sprintf("[%s] [%s] %s", Cyan(utils.GetTime()), Red("FATAL"), msg))
}

func Error(msg string) {
	log(LevelError, fmt.Sprintf("[%s] [%s] %s", Cyan(utils.GetTime()), LightRed("ERROR"), msg))
}

func Warning(msg string) {
	log(LevelWarning, fmt.Sprintf("[%s] [%s] %s", Cyan(utils.GetTime()), Yellow("WARNING"), msg))
}

func Success(msg string) {
	log(LevelSuccess, fmt.Sprintf("[%s] [%s] %s", Cyan(utils.GetTime()), Green("SUCCESS"), msg))
}

func Info(msg string) {
	log(LevelInfo, fmt.Sprintf("[%s] [%s] %s", Cyan(utils.GetTime()), LightGreen("INFO"), msg))
}

func General(level Level, msg string, title string) {
	log(level, fmt.Sprintf("[%s] [%s] %s", Cyan(utils.GetTime()), LightGreen(title), msg))
}

func WriteLogFile(result string, filename string) {
	var text = []byte(result + "\n")
	fl, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600) // #nosec G304
	if err != nil {
		fmt.Printf("Open %s error, %v\n", filename, err)
		return
	}
	_, err = fl.Write(text)
	if closeErr := fl.Close(); closeErr != nil {
		fmt.Printf("Close %s error, %v\n", filename, closeErr)
	}
	if err != nil {
		fmt.Printf("Write %s error, %v\n", filename, err)
	}
}

func StripAnsi(str string) string {
	const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"
	var re = regexp.MustCompile(ansi)
	return re.ReplaceAllString(str, "")
}

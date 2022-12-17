package logger

type Level int

const (
	LevelFatal Level = iota
	LevelError
	LevelWarning
	LevelSuccess
	LevelInfo
)

package logger

import (
	"github.com/fatih/color"
	"log"
	"strings"
)

type Type string

const (
	//Logger types
	LogTypeHeader Type = "Header"
	LogTypeDebug   Type = "Debug"
	LogTypeWarning Type = "Warning"
	LogTypeFatal   Type = "Fatal"
)

type Logger struct {
	tag string
}

func (logger *Logger) Write(msg string, logType Type){
	switch logType {
	case LogTypeHeader:
		log.Println(color.GreenString(strings.Repeat("*", len(msg)+8)))
		log.Printf(color.GreenString("%s ")+"%s "+color.GreenString("%s"), strings.Repeat("*", 3), msg, strings.Repeat("*", 3))
		log.Println(color.GreenString(strings.Repeat("*", len(msg)+8)))
	case LogTypeDebug:
		log.Println(msg)
	case LogTypeWarning:
		log.Println(msg)
	case LogTypeFatal:
		log.Println(msg)
	}
}
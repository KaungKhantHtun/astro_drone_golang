package config

import (
	"os"

	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
)

func InitLog() {

	log.SetLevel(getLoggerLevel(os.Getenv("LOG_LEVEL")))
	log.SetReportCaller(true)
	log.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
		TimestampFormat: "2000-01-01 00:00:00",
		ShowFullLevel:   true,
		CallerFirst:     true,
	})

}

func getLoggerLevel(value string) log.Level {
	switch value {
	case "DEBUG":
		return log.DebugLevel
	case "TRACE":
		return log.TraceLevel
	default:
		return log.InfoLevel
	}
}

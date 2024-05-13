package tools

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

const format = "2006-01-02 15:04:05"

func ConvertTimestampToTime(timestamp int64) time.Time {
	t := time.Unix(timestamp, 0)
	return t
}

func ConvertTimeToTimestamp(date string) time.Time {
	//Check the documentation on Go for the const variables!
	//They need to be exactly as they are shown in the documentation to be read correctly!

	t, err := time.Parse(format, date)

	if err != nil {
		log.Errorf("Failed to convert date to timestamp: %v \n %e", date, err)
	} else {
		return t
	}
	return time.Now()
}

func getEnv() string {
	logLevel := os.Getenv("LOG_LEVEL")
	if len(logLevel) == 0 {
		return "debug"
	}
	return logLevel
}

func SetupLogging() {

	logLevel := getEnv()
	switch logLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)

	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "trace":
		log.SetLevel(log.TraceLevel)
		file, _ := os.OpenFile("/var/log/trace.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		mw := io.MultiWriter(file, os.Stdout)
		log.SetOutput(mw)
	default:
		log.SetLevel(log.InfoLevel)
	}
}

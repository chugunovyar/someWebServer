package tools

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
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

func GetEnv(requestedEnvVar string, defaultVar string) string {
	envVariable := os.Getenv(requestedEnvVar)
	if len(envVariable) == 0 {
		log.Errorf("Environment string variable %s is empty %s", requestedEnvVar, envVariable)
		return defaultVar
	}
	return envVariable
}

func GetEnvAsInt(requestedEnvVar string, defaultVar int64) int64 {
	envVariable := os.Getenv(requestedEnvVar)
	if len(envVariable) == 0 {
		log.Errorf("Environment int64 variable %s is empty %s", requestedEnvVar, envVariable)
		return defaultVar
	}
	envVarAsInt, _ := strconv.ParseInt(envVariable, 10, 64)
	return envVarAsInt
}

func SetupLogging() {

	logLevel := GetEnv("LOG_LEVEL", "debug")
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
	default:
		log.SetLevel(log.InfoLevel)
	}
}

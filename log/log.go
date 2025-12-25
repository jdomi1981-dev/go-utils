package log

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

const (
	ERROR_LEVEL = "ERROR"
)

type Log struct {
	Environment string    `json:"environment"`
	LoggerName  string    `json:"loggerName"`
	Datetime    time.Time `json:"datetime"`
	Level       string    `json:"level"`
	Message     string    `json:"message"`
	Error       string    `json:"error,omitempty"`
}

func Logger(env string, level string, loggerName string, message string, err error) Log {
	var log Log
	log.Environment = env
	log.Message = message
	log.Datetime = time.Now().UTC()
	log.Level = level
	log.LoggerName = loggerName
	if strings.ToUpper(level) == ERROR_LEVEL {
		log.Error = err.Error()
	}
	return log
}

func PrintLog(log Log) {
	str, _ := json.Marshal(log)
	fmt.Println(string(str))
}

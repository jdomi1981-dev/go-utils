package utils

import (
	"strings"
	"time"
)

// HasWhiteSpaces indicate if a string has only white spaces
func HasWhiteSpaces(value string) bool {
	var invalidStr bool
	if len(strings.TrimSpace(value)) == 0 {
		invalidStr = true
	}
	return invalidStr
}

// SearchValue search a value inside a string
func SearchValue(evaluatedValue string, line string, separator string) bool {
	var vFind bool
	splitProperty := strings.SplitSeq(line, separator)

	for line := range splitProperty {
		if strings.EqualFold(evaluatedValue, line) {
			vFind = true
		}
	}
	return vFind
}

// GetCurrentTimeMillis provides the current time in milliseconds.
func GetCurrentTimeMillis(timeProvider func() time.Time) int64 {
	return timeProvider().UnixNano() / 1e6
}

func NowPostgres() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}

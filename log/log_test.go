package log_test

import (
	"testing"

	"github.com/jdomi1981-dev/go-utils/log"
)

func TestLogger(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		env        string
		level      string
		loggerName string
		message    string
		err        error
		want       log.Log
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := log.Logger(tt.env, tt.level, tt.loggerName, tt.message, tt.err)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("Logger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintLog(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		log log.Log
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.PrintLog(tt.log)
		})
	}
}

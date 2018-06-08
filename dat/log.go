package dat

import (
	stdlog "log"

	"github.com/helloeave/dat/internal/log"
)

type LogFunc = log.LogFunc

var StdLogger = func(msg string, kvs ...interface{}) {
	stdlog.Println(msg, kvs)
}

func init() {
	SetErrorLogger(StdLogger)
}

// SetDebugLogger sets a logger for use when dat encounters interesting debug information. Defaults to a NoOp logger.
func SetDebugLogger(l LogFunc) {
	log.DebugLogFn = l
}

// SetErrorLogger sets a logger all error occurrences in dat. Defaults to stdlib log.Printf.
func SetErrorLogger(l LogFunc) {
	log.ErrorLogFn = l
}

// SetSQLLogger sets a logger for recording sql queries and metrics. Defaults to a NoOp logger. By setting a logger, sql queries will be logged.
func SetSQLLogger(l LogFunc) {
	log.SQLLogFn = l
}

package log

import (
	"fmt"
	"log"
	"os"
)

var (
	debugLogFn LogFunc = nil
	sqlLogFn   LogFunc = nil
	errLogFn   LogFunc = log.Printf
)

type LogFunc = func(string, ...interface{})

// HasDebugLogger is a way to check for a noop on debug logging and pre-empt an expensive formatting call.
func HasDebugLogger() bool {
	return debugLogFn != nil
}

// HasErrLogger is a way to check for a noop on error logging and pre-empt an expensive formatting call.
func HasErrLogger() bool {
	return errLogFn != nil
}

// HasSqlLogger is a way to check for a noop on sql logging and pre-empt an expensive formatting call.
func HasSqlLogger() bool {
	return sqlLogFn != nil
}

// SetDebug sets a logger for use when dat encounters interesting debug information. Defaults to a NoOp logger.
func SetDebug(l LogFunc) {
	debugLogFn = l
}

// SetErr sets a logger all error occurrences in dat. Defaults to stdlib log.Printf.
func SetErr(l LogFunc) {
	errLogFn = l
}

// SetSql sets a logger for recording sql queries and metrics. Defaults to a NoOp logger. By setting a logger, sql queries will be logged.
func SetSql(l LogFunc) {
	sqlLogFn = l
}

// Error is a temporary helper to replace logger.Error from logxi
func Error(msg string, vals ...interface{}) error {
	Err(msg, vals)
	return fmt.Errorf(fmt.Sprintln(msg, vals))
}

// Fatal is a temporary helper to replace logger.Fatal from logxi
func Fatal(msg string, vals ...interface{}) {
	Err(msg, vals)
	os.Exit(1)
}

// Debug logs to the debugLogFn if it is set. Noop default.
func Debug(msg string, vals ...interface{}) {
	if HasDebugLogger() {
		debugLogFn(msg, vals)
	}
}

// Err logs to the errLogFn if it is set. stdlib log.Printf default.
func Err(msg string, vals ...interface{}) {
	if HasErrLogger() {
		errLogFn(msg, vals)
	}
}

// Sql logs to the sqlLogFn if it is set. Noop default.
func Sql(msg string, vals ...interface{}) {
	if HasSqlLogger() {
		sqlLogFn(msg, vals)
	}
}

package log

import (
	"fmt"
	"os"
)

var (
	DebugLogFn LogFunc = nil
	SQLLogFn   LogFunc = nil
	ErrorLogFn LogFunc = nil
)

type LogFunc = func(string, ...interface{})

// HasDebugLogger is a way to check for a noop on debug logging and pre-empt an expensive formatting call.
func HasDebugLogger() bool {
	return DebugLogFn != nil
}

// HasErrLogger is a way to check for a noop on error logging and pre-empt an expensive formatting call.
func HasErrLogger() bool {
	return ErrorLogFn != nil
}

// HasSQLLogger is a way to check for a noop on sql logging and pre-empt an expensive formatting call.
func HasSQLLogger() bool {
	return SQLLogFn != nil
}

// ErrorE is a temporary helper to replace logger.ErrorE from logxi
func ErrorE(msg string, vals ...interface{}) error {
	Error(msg, vals...)
	return fmt.Errorf(fmt.Sprintln(msg, vals))
}

// Fatal is a temporary helper to replace logger.Fatal from logxi
func Fatal(msg string, vals ...interface{}) {
	Error(msg, vals...)
	os.Exit(1)
}

// Debug logs to the debugLogFn if it is set. Noop default.
func Debug(msg string, vals ...interface{}) {
	if HasDebugLogger() {
		DebugLogFn(msg, vals...)
	}
}

// Error logs to the errLogFn if it is set. stdlib log.Printf default.
func Error(msg string, vals ...interface{}) {
	if HasErrLogger() {
		ErrorLogFn(msg, vals...)
	}
}

// SQL logs to the sqlLogFn if it is set. Noop default.
func SQL(msg string, vals ...interface{}) {
	if HasSQLLogger() {
		SQLLogFn(msg, vals...)
	}
}

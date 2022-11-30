package logging

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"log"
	"os"
)

var logger struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	fatalLog *log.Logger
}

func init() {
	dsn := os.Getenv("SENTRY_DSN")
	if dsn == "" {
		panic("SENTRY_DSN environmental variable not set\n")
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		TracesSampleRate: 1.0,
	})

	if err != nil {
		m := fmt.Sprintf("Sentry initialization failed: %s", err)
		panic(m)
	}

	logger.infoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	logger.errorLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
	logger.fatalLog = log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime)
}

// LogInfo logs a message that is not an error.
// For errors, use LogError or LogFatal.
func LogInfo(message string) {
	logger.infoLog.Println(message)
}

// LogError logs a non-fatal error (with an error type).
// It also sends the error to Sentry.
// For errors with string messages, use LogErrorMessage.
// For fatal errors, use LogFatal.
func LogError(err error) {
	m := fmt.Sprintf("Error: %s", err)
	e := fmt.Errorf(m)

	sentry.CaptureException(e)
	logger.errorLog.Println(m)
}

// LogErrorMessage logs a non-fatal error (with a string message).
// It also sends the error to Sentry.
// For errors with error types, use LogError.
// For fatal errors, use LogFatalMessage.
func LogErrorMessage(message string) {
	m := fmt.Sprintf("Error: %s", message)
	e := fmt.Errorf(m)

	sentry.CaptureException(e)
	logger.errorLog.Println(m)
}

// LogFatal logs a fatal error (with an error type).
// It also sends the error to Sentry and exits the program with code 1.
// For errors with string messages, use LogFatalMessage.
// For non-fatal errors, use LogError.
func LogFatal(err error) {
	m := fmt.Sprintf("Fatal: %s", err)
	e := fmt.Errorf(m)

	sentry.CaptureException(e)
	logger.fatalLog.Println(m)

	os.Exit(1)
}

// LogFatalMessage logs a fatal error (with a string message).
// It also sends the error to Sentry and exits the program with code 1.
// For errors with error types, use LogFatal.
// For non-fatal errors, use LogErrorMessage.
func LogFatalMessage(message string) {
	m := fmt.Sprintf("Fatal: %s", message)
	e := fmt.Errorf(m)

	sentry.CaptureException(e)
	logger.fatalLog.Println(m)

	os.Exit(1)
}

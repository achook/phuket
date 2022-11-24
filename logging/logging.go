package logging

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"log"
	"os"
)

type Logger struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	fatalLog *log.Logger
}

func (l *Logger) Initialize() error {
	dsn := os.Getenv("SENTRY_DSN")
	if dsn == "" {
		return fmt.Errorf("SENTRY_DSN environmental variable not set\n")
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		TracesSampleRate: 1.0,
	})

	if err != nil {
		return fmt.Errorf("Sentry initilization failed: %s\n", err)
	}

	l.infoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.errorLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.fatalLog = log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}

func (l *Logger) LogInfo(message string) {
	l.infoLog.Println(message)
}

func (l *Logger) LogError(message string) {
	m := fmt.Sprintf("Error: %s", message)
	sentry.CaptureMessage(m)
	
	l.errorLog.Println(message)
}

func (l *Logger) LogFatal(message string) {
	m := fmt.Sprintf("Fatal: %s", message)
	sentry.CaptureMessage(m)

	l.fatalLog.Println(message)
}

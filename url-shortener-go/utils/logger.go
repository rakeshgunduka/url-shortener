package utils

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	Info  func(message string, rest ...string)
	Error func(err error, message string, rest ...string)
}

func CreateLogger(logContext string) (context *Logger) {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
	})

	context = &Logger{
		Info: func(message string, rest ...string) {
			m := []string{message}
			if len(rest) > 0 {
				m = append(m, rest...)
			}
			logger.Info(strings.Join(m, ":"))
		},
		Error: func(err error, message string, rest ...string) {
			m := []string{message}
			if len(rest) > 0 {
				m = append(m, rest...)
			}
			logger.WithError(err).Error(strings.Join(m, ":"))
		},
	}

	return
}

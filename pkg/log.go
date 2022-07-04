package pkg

import (
	"os"

	"github.com/sirupsen/logrus"
)

// NewLogger — конструктор логгера
func NewLogger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)

	return log
}

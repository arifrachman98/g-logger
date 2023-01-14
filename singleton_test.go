package gologging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingletonLogging(t *testing.T) {
	logrus.Info("Nope")
	logrus.Warn("Nothing")

	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	logrus.Info("Duarrrr!!!")
}

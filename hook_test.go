// The package name.
package gologging

// Importing the testing package.
import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type SampleHook struct {
}

// Returning the levels that the hook will be called for.
func (s *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

// A function that is called when the hook is fired.
func (s *SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Sample Hook", entry.Level, entry.Message)
	return nil
}

// logger, and then logs a message
func TestLogHook(t *testing.T) {
	// Creating a new instance of the logrus logger.
	log := logrus.New()
	// Adding a hook to the logger.
	log.AddHook(&SampleHook{})

	// Logging the message to the console.
	log.Info("Log info")
	log.Warn("Log warn")
	log.Error("Log Warn")
	log.Debug("Log Warn")

}

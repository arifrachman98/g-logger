// The package name.
package gologging
// Importing the packages that are needed for the program to run.
import (
	"fmt"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLog(t *testing.T) {
	// Creating a new instance of the logger.
	logger := logrus.New()

	// Printing the message to the console.
	logger.Println("Hello logger sirupsen")
	fmt.Println("Hello logger Go-Lang")
}

func TestLogLevel(t *testing.T) {
	// It creates a new instance of the logger.
	log := logrus.New()

	// A log level.
	log.Trace("This is trace logger")
	log.Debug("This is debug logger")
	log.Info("This is info logger")
	log.Warn("This is warn logger")
	log.Error("This is error logger")
}

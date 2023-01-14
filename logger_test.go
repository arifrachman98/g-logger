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

func TestLogSetLevel(t *testing.T) {
	// Setting the log level to trace level.
	log := logrus.New()
	log.SetLevel(logrus.TraceLevel)

	// Printing the message to the console.
	log.Trace("Trace level")
	log.Debug("Debug level")
	log.Info("Info level")
	log.Warn("Warn level")
	log.Error("Error level")
}

func TestLogOutputFile(t *testing.T) {
	// It creates a new instance of the logger.
	log := logrus.New()

	// Creating a file called app.log and writing to it.
	file, _ := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	// Setting the output of the log to the file.
	log.SetOutput(file)

	// Printing the message to the console.
	log.Info("Info Logger")
	log.Warn("Warn Logger")
	log.Error("Error Logger")

}

func TestLogOutputJSONFormatter(t *testing.T) {
	// It creates a new instance of the logger.
	log := logrus.New()

	// Setting the formatter to JSON.
	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	// Printing the message to the console.
	log.Info("Log info over here")
	log.Warn("Log warn is here")
	log.Error("Log error was here")
}

func TestLogOutputJSONxFile(t *testing.T) {
	// It creates a new instance of the logger.
	log := logrus.New()

	// Creating a file called dir.log and writing to it.
	file, _ := os.OpenFile("dir.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	// Setting the formatter to JSON.
	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	// Setting the output of the log to the file.
	log.SetOutput(file)

	// Printing the message to the console.
	log.Trace("Trace level good")
	log.Debug("Debug level bad enough")
	log.Info("Info level was clear")
	log.Warn("Warn level negative")
	log.Error("Error level secure")
}

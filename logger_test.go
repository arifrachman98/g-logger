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

func TestLogWithField(t *testing.T) {
	// It creates a new instance of the logger.
	log := logrus.New()

	// Creating a file called dir.log and writing to it.
	file, _ := os.OpenFile("user.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	// Setting the formatter to JSON.
	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	// Setting the output of the log to the file.
	log.SetOutput(file)

	// Adding a field to the log.
	log.WithField("username", "jokokendic").Info("Joko was here")
	log.WithField("name", "Jaka").
		WithField("username", "jaakee").
		Info("Jaka the other side")
}

func TestLogWithFields(t *testing.T) {
	// It creates a new instance of the logger.
	log := logrus.New()

	// Creating a file called dir.log and writing to it.
	file, _ := os.OpenFile("address.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	// Setting the formatter to JSON.
	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	// Setting the output of the log to the file.
	log.SetOutput(file)

	// Adding a field to the log.
	log.WithFields(logrus.Fields{
		"Address": "Phenompend",
		"City":    "Logero",
	}).Info("address user 1")

	log.WithFields(logrus.Fields{
		"Address": "Kindape X",
		"City":    "Koraemo",
	}).Info("address user 2")

	log.WithFields(logrus.Fields{
		"Address": "Lomera",
		"City":    "Jigeso",
	}).Info("address user 3")
}

func TestLogEntry(t *testing.T) {
	// It creates a new instance of the logger.
	log := logrus.New()

	// Creating a file called dir.log and writing to it.
	file, _ := os.OpenFile("bid.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	// Setting the formatter to JSON.
	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	// Setting the output of the log to the file.
	log.SetOutput(file)

	// A for loop that is looping 100 times.
	for i := 1; i <= 100; i++ {
		// Creating a new entry for the log.
		entry := logrus.NewEntry(log)
		// Checking if the number is even.
		if i%2 == 0 {
			entry.WithFields(logrus.Fields{
				"name": "Joko",
				"bid":  i,
			}).Info("Rejected")
			// Checking if the number is divisible by 3.
		} else if i%3 == 0 {
			entry.WithFields(logrus.Fields{
				"name": "Joko",
				"bid":  i,
			}).Info("Deny")
			// Checking if the number is divisible by 5.
		} else if i%5 == 0 {
			entry.WithFields(logrus.Fields{
				"name": "Joko",
				"bid":  i,
			}).Info("OUT")
			// The else statement.
		} else {
			entry.WithFields(logrus.Fields{
				"name": "Joko",
				"bid":  i,
			}).Info("Accept")
		}
	}
}

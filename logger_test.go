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

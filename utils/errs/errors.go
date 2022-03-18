package errs

import (
	"log"
)

func HandleFatalError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

package utils

import (
	"log"
	"time"
)

func VerboseTime(category string, verbose bool) {
	if verbose {
		log.Printf("%s: %s", category, time.Now().String())
	}
}

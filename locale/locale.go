package locale

import (
	"os"
	"time"
)

// SetLocale setting locale
func SetLocale() {
	loc, err := time.LoadLocation(os.Getenv("LOCATION"))
	if err != nil {
		loc = time.FixedZone(os.Getenv("LOCATION"), 9*60*60)
	}
	time.Local = loc
}

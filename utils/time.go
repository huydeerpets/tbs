package utils

import "time"

var now time.Time

// Now - Get current date and time
func Now() time.Time {
	if IsTest() {
		return now
	}

	return time.Now()
}

// SetNow - Set the current date and time
func SetNow(n time.Time) {
	now = n
}

package AdvancedGoConcurrencyPatterns

import "time"

// Fetch fetches Items for uri and returns the time when the next
// fetch should be attempted.  On failure, Fetch returns an error.
func Fetch(uri string) (items []Item, next time.Time, err error) {
	return nil, time.Time{}, nil
}

package AdvancedGoConcurrencyPatterns

import "time"

type Fetcher interface {
	// Fetch fetches Items for uri and returns the time when the next
	// fetch should be attempted.  On failure, Fetch returns an error.
	Fetch() (items []Item, next time.Time, err error)
}

type fetchResult struct {
	fetched []Item
	next    time.Time
	err     error
}

// fetches Items from domain
func Fetch(domain string) Fetcher {
	return nil
}

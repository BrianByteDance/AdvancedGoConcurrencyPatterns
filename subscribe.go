package AdvancedGoConcurrencyPatterns

import "time"

const maxPending = 10

type Subscription interface {
	Updates() <-chan Item // stream of Items
	Close() error         // shuts down the stream
}

// converts Fetches to a stream
func Subscribe(fetcher Fetcher) Subscription {
	s := &sub{
		fetcher: fetcher,
		updates: make(chan Item), // for Updates
	}
	go s.loop()
	return s
}

// merges several streams
func Merge(subs ...Subscription) Subscription {
	return nil
}

// sub implements the Subscription interface.
type sub struct {
	fetcher Fetcher   // fetches items
	updates chan Item // delivers items to the user
	closing chan chan error
}

// loop fetches items using s.fetcher and sends them
// on s.updates.  loop exits when s.Close is called.
func (s *sub) loop() {
	var pending []Item               // appended by fetch; consumed by send
	var next time.Time               // initially January 1, year 0
	var err error                    // set when Fetch fails
	var seen = make(map[string]bool) // set of item.GUIDs
	var fetchDone chan fetchResult   // if non-nil, Fetch is running
	for {
		var fetchDelay time.Duration // initially 0 (no delay)
		if now := time.Now(); next.After(now) {
			fetchDelay = next.Sub(now)
		}
		var startFetch <-chan time.Time
		if fetchDone == nil && len(pending) < maxPending {
			startFetch = time.After(fetchDelay) // enable fetch case
		}

		var first Item
		var updates chan Item
		if len(pending) > 0 {
			first = pending[0]
			updates = s.updates // enable send case
		}

		select {
		case errc := <-s.closing:
			errc <- err
			close(s.updates) // tells receiver we're done
			return
		case <-startFetch:
			fetchDone = make(chan fetchResult, 1)
			go func() {
				fetched, next, err := s.fetcher.Fetch()
				fetchDone <- fetchResult{fetched, next, err}
			}()
		case result := <-fetchDone:
			fetchDone = nil
			// Use result.fetched, result.next, result.err
			if result.err != nil {
				next = time.Now().Add(10 * time.Second)
				break
			}
			next = result.next
			for _, item := range result.fetched {
				if !seen[item.GUID] {
					pending = append(pending, item)
					seen[item.GUID] = true
				}
			}
		case updates <- first:
			pending = pending[1:]
		}
	}
}
func (s *sub) Updates() <-chan Item {
	return s.updates
}

func (s *sub) Close() error {
	errc := make(chan error)
	s.closing <- errc
	return <-errc
}

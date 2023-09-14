package AdvancedGoConcurrencyPatterns

import "time"

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
}

// loop fetches items using s.fetcher and sends them
// on s.updates.  loop exits when s.Close is called.
func (s *sub) loop() {}

func (s *sub) Updates() <-chan Item {
	return s.updates
}

func (s *sub) Close() error {
	// TODO: make loop exit
	// TODO: find out about any error
	return nil
}

type naiveSub struct {
	sub
	closed bool
	err    error
}

func (s *naiveSub) loop() {
	for {
		if s.closed {
			close(s.updates)
			return
		}
		items, next, err := s.fetcher.Fetch()
		if err != nil {
			s.err = err
			time.Sleep(10 * time.Second)
			continue
		}
		for _, item := range items {
			s.updates <- item
		}
		if now := time.Now(); next.After(now) {
			time.Sleep(next.Sub(now))
		}
	}
}

func (s *naiveSub) Close() error {
	s.closed = true
	return s.err
}

package AdvancedGoConcurrencyPatterns

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

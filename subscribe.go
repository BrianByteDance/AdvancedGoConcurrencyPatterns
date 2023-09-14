package AdvancedGoConcurrencyPatterns

type Subscription interface {
	Updates() <-chan Item // stream of Items
	Close() error         // shuts down the stream
}

// converts Fetches to a stream
func Subscribe(fetcher Fetcher) Subscription {
	return nil
}

// merges several streams
func Merge(subs ...Subscription) Subscription {
	return nil
}

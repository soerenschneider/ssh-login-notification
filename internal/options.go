package internal

// Options contains all user set runtime options.
type Options struct {
	GeoLookup     bool
	DnsLookup     bool
	TelegramId    int64
	TelegramToken string
}

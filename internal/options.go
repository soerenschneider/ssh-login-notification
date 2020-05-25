package internal

// Options contains all user set runtime options.
type Options struct {
	GeoLookup        bool
	DnsLookup        bool
	IgnorePrivateIps bool
	TelegramId       int64
	TelegramToken    string
}

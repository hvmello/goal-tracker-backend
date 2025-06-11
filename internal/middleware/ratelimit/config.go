package ratelimit

type Config struct {
	RequestsPerMinute int
	BurstSize         int
}

func NewConfig(requestsPerMinute, burstSize int) Config {
	return Config{
		RequestsPerMinute: requestsPerMinute,
		BurstSize:         burstSize,
	}
}

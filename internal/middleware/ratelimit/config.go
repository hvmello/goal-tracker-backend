package ratelimit

type Config struct {
	RequestsPerMinute int
	BurstSize         int
	RateLimit         struct {
		Enabled           bool `env:"RATE_LIMIT_ENABLED" envDefault:"true"`
		RequestsPerMinute int  `env:"RATE_LIMIT_REQUESTS_PER_MINUTE" envDefault:"60"`
		BurstSize         int  `env:"RATE_LIMIT_BURST_SIZE" envDefault:"20"`
	}
}

func NewConfig(requestsPerMinute, burstSize int) Config {
	return Config{
		RequestsPerMinute: requestsPerMinute,
		BurstSize:         burstSize,
	}
}

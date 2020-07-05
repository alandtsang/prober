package config

var DefaultConfig = &Config{
	InitialDelaySeconds: 10,
	TimeoutSeconds:      10,
	PeriodSeconds:       10,
}

type Config struct {
	InitialDelaySeconds int32
	TimeoutSeconds      int32
	PeriodSeconds       int32
}

// NewConfig return Config pointer.
func NewConfig(initDelaySec, timeoutSec, periodSec int32) *Config {
	return &Config{
		InitialDelaySeconds: initDelaySec,
		TimeoutSeconds:      timeoutSec,
		PeriodSeconds:       periodSec,
	}
}

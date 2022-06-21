package config

import "sync"

type (
	businessConfig struct {
		sync.Mutex
	}
)

var runtimeBusinessConfig = newBusinessConfig()

func newBusinessConfig() businessConfig {
	// fallback values
	return businessConfig{}
}

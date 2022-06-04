package config

var RuntimeConfig = NewBusinessConfig()

type (
	BusinessConfig struct {
		User UserConfig `mapstructure:",squash"`
	}

	UserConfig struct {
		CachingEnabled bool `mapstructure:"USER_CACHING_ENABLED"`
	}
)

// NewBusinessConfig initiate business config object.
func NewBusinessConfig() BusinessConfig {
	// fallback values
	return BusinessConfig{
		User: UserConfig{
			CachingEnabled: false,
		},
	}
}

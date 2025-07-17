package luksdk

func NewLukSDKFromConfig(config *Config) *LukSDK {
	config.parseDefault()
	luksdk := &LukSDK{
		config: *config,
	}
	luksdk.Apis = newApis(luksdk)
	return luksdk
}

func NewLukSDKWithOptions(opts ...Option) *LukSDK {
	config := NewConfig()
	config.WithOptions(opts...)
	return NewLukSDKFromConfig(config)
}

func NewLukSDKWithConfigurators(configurators ...Configurator) *LukSDK {
	config := NewConfig()
	for _, configurator := range configurators {
		configurator.Configure(config)
	}
	return NewLukSDKFromConfig(config)
}

type LukSDK struct {
	config Config
	*Apis
}

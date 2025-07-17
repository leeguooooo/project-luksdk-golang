package luksdk

import "resty.dev/v3"

func NewConfig() *Config {
	return &Config{}
}

type Configurator interface {
	Configure(config *Config)
}

type ConfiguratorFN func(config *Config)

func (fn ConfiguratorFN) Configure(config *Config) {
	fn(config)
}

type Option func(config *Config)

type Config struct {
	AppId      int64         // APP ID
	AppSecret  string        // APP 密钥
	Domain     string        // API 域名
	Debug      bool          // 是否开启调试模式
	HttpClient *resty.Client // HTTP 客户端
}

func (c *Config) parseDefault() {
	if c.HttpClient == nil {
		c.HttpClient = resty.New()
	}

	c.HttpClient.
		SetDebug(c.Debug).
		SetBaseURL(c.Domain).
		SetHeader("User-Agent", "golang/v1.0.0")
}

func (c *Config) WithAppId(appId int64) *Config {
	c.AppId = appId
	return c
}

func WithAppId(appId int64) Option {
	return func(config *Config) {
		config.WithAppId(appId)

	}
}

func (c *Config) WithAppSecret(appSecret string) *Config {
	c.AppSecret = appSecret
	return c
}

func WithAppSecret(appSecret string) Option {
	return func(config *Config) {
		config.WithAppSecret(appSecret)
	}
}

func (c *Config) WithOptions(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
}

func (c *Config) WithDomain(domain string) *Config {
	c.Domain = domain
	return c
}

func WithDomain(domain string) Option {
	return func(config *Config) {
		config.WithDomain(domain)
	}
}

func (c *Config) WithDebug(debug bool) *Config {
	c.Debug = debug
	return c
}

func WithDebug(debug bool) Option {
	return func(config *Config) {
		config.WithDebug(debug)
	}
}

func (c *Config) WithHttpClient(client *resty.Client) *Config {
	c.HttpClient = client
	return c
}

func WithHttpClient(client *resty.Client) Option {
	return func(config *Config) {
		config.WithHttpClient(client)
	}
}

func WithOptions(opts ...Option) Option {
	return func(config *Config) {
		config.WithOptions(opts...)
	}
}

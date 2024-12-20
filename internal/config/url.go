package config

type urlConfig struct {
	url string
}

// NewURLConfig создает новый экземпляр URLConfig
func NewURLConfig() URLProvider {
	return &urlConfig{
		url: baseURL,
	}
}

func (cfg *urlConfig) BaseURL() string {
	return cfg.url
}

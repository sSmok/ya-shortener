package config

import "os"

const baseURLEnv = "BASE_URL"

type urlConfig struct {
	url string
}

// NewURLConfig создает новый экземпляр URLConfig
func NewURLConfig() URLProvider {
	var URL string
	URL = os.Getenv(baseURLEnv)
	if len(URL) == 0 {
		URL = baseURL
	}
	return &urlConfig{
		url: URL,
	}
}

func (cfg *urlConfig) BaseURL() string {
	return cfg.url
}

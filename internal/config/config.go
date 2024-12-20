package config

import "flag"

// AddressProvider предоставляет адрес HTTP сервера
type AddressProvider interface {
	Address() string
}

// URLProvider предоставляет базовый URL для коротких ссылок
type URLProvider interface {
	BaseURL() string
}

var (
	httpAddress string
	baseURL     string
)

func Load() error {
	flag.StringVar(&httpAddress, "a", defaultAddress, "address for HTTP server")
	flag.StringVar(&baseURL, "b", defaultBaseURL, "base URL for short links")
	flag.Parse()

	return nil
}

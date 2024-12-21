package config

import (
	"flag"

	"github.com/joho/godotenv"
)

const (
	defaultAddress = "localhost:8080"
	defaultBaseURL = "http://localhost:8080"
)

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

// Load загружает конфигурацию из флагов
func Load() error {
	flag.StringVar(&httpAddress, "a", defaultAddress, "address for HTTP server")
	flag.StringVar(&baseURL, "b", defaultBaseURL, "base URL for short links")
	flag.Parse()

	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

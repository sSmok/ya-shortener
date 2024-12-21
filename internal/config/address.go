package config

import "os"

const addressEnv = "HTTP_ADDRESS"

type addressConfig struct {
	address string
}

// NewAddressConfig создает новый экземпляр HTTPAddress
func NewAddressConfig() AddressProvider {
	var addr string
	addr = os.Getenv(addressEnv)
	if len(addr) == 0 {
		addr = httpAddress
	}
	return &addressConfig{
		address: addr,
	}
}

func (cfg *addressConfig) Address() string {
	return cfg.address
}

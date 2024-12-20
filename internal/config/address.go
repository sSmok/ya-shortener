package config

type addressConfig struct {
	address string
}

// NewAddressConfig создает новый экземпляр HTTPAddress
func NewAddressConfig() AddressProvider {
	return &addressConfig{
		address: httpAddress,
	}
}

func (cfg *addressConfig) Address() string {
	return cfg.address
}

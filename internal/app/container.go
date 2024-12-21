package app

import (
	"github.com/sSmok/ya-shortener/internal/api/link"
	"github.com/sSmok/ya-shortener/internal/config"
	"github.com/sSmok/ya-shortener/internal/repository"
	linkRepository "github.com/sSmok/ya-shortener/internal/repository/link"
)

type container struct {
	linkAPI       *link.API
	linkRepo      repository.LinkRepository
	urlConfig     config.URLProvider
	addressConfig config.AddressProvider
}

func newContainer() *container {
	return &container{}
}

func (c *container) LinkAPI() *link.API {
	if c.linkAPI == nil {
		c.linkAPI = link.NewAPI(c.LinkRepo(), c.URLConfig().BaseURL())
	}

	return c.linkAPI
}

func (c *container) LinkRepo() repository.LinkRepository {
	if c.linkRepo == nil {
		c.linkRepo = linkRepository.NewRepository()
	}

	return c.linkRepo
}

func (c *container) URLConfig() config.URLProvider {
	if c.urlConfig == nil {
		c.urlConfig = config.NewURLConfig()
	}

	return c.urlConfig
}

func (c *container) AddressConfig() config.AddressProvider {
	if c.addressConfig == nil {
		c.addressConfig = config.NewAddressConfig()
	}

	return c.addressConfig
}

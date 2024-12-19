package app

import (
	"github.com/sSmok/ya-shortener/internal/api/link"
	"github.com/sSmok/ya-shortener/internal/repository"
	linkRepository "github.com/sSmok/ya-shortener/internal/repository/link"
)

type container struct {
	linkAPI  *link.API
	linkRepo repository.LinkRepository
}

func newContainer() *container {
	return &container{}
}

func (c *container) LinkAPI() *link.API {
	if c.linkAPI == nil {
		c.linkAPI = link.NewAPI(c.LinkRepo())
	}

	return c.linkAPI
}

func (c *container) LinkRepo() repository.LinkRepository {
	if c.linkRepo == nil {
		c.linkRepo = linkRepository.NewRepository()
	}

	return c.linkRepo
}

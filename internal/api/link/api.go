package link

import (
	"github.com/sSmok/ya-shortener/internal/repository"
)

// API представляет собой обработчик запросов
type API struct {
	linkRepo repository.LinkRepository
}

// NewAPI создает новый экземпляр API
func NewAPI(linkRepo repository.LinkRepository) *API {
	return &API{
		linkRepo: linkRepo,
	}
}

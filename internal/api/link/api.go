package link

import (
	"github.com/sSmok/ya-shortener/internal/repository"
)

// API представляет собой обработчик запросов
type API struct {
	linkRepo repository.LinkRepository
	baseURL  string
}

// NewAPI создает новый экземпляр API
func NewAPI(linkRepo repository.LinkRepository, baseURL string) *API {
	return &API{
		linkRepo: linkRepo,
		baseURL:  baseURL,
	}
}

package link

import (
	"github.com/sSmok/ya-shortener/internal/repository"
)

type repo struct {
	storage map[string]string
}

// NewRepository создает новый экземпляр репозитория для хранилища ссылок
func NewRepository() repository.LinkRepository {
	return &repo{
		storage: make(map[string]string),
	}
}

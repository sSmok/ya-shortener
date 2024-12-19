package repository

// LinkRepository представляет собой контракт для работы с репозиторием ссылок
type LinkRepository interface {
	Create(url string) (string, error)
	Get(short string) (string, error)
}

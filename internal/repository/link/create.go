package link

import (
	"math/rand/v2"

	"github.com/sSmok/ya-shortener/internal/utils"
)

func (r *repo) Create(url string) (string, error) {
	// #nosec G404 // и так используем math/rand/v2
	short := utils.Base62Encode(rand.Uint64())
	r.storage[short] = url

	return short, nil
}

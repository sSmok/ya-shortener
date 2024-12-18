package link

import (
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Create обрабатывает запросы на создание короткой ссылки
func (api *API) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	originalURL := string(body)
	fmt.Println("REQ BODY: ", originalURL)

	// #nosec G404 // и так используем math/rand/v2
	short := base62Encode(rand.Uint64())
	api.storage[short] = originalURL
	w.WriteHeader(http.StatusCreated)
	shortURL := fmt.Sprintf("http://%s/%s", address, short)
	_, err = w.Write([]byte(shortURL))
	if err != nil {
		return
	}

}

func base62Encode(number uint64) string {
	length := uint(len(alphabet))
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(10)
	for ; number > 0; number = number / uint64(length) {
		encodedBuilder.WriteByte(alphabet[(number % uint64(length))])
	}

	return encodedBuilder.String()
}

package link

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const address = "localhost:8080"

// Create обрабатывает запросы на создание короткой ссылки
func (api *API) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	defer func() {
		err := r.Body.Close()
		if err != nil {
			log.Fatalf("failed to close request body: %v", err)
		}
	}()

	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	short, err := api.linkRepo.Create(string(body))
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusCreated)
	shortURL := fmt.Sprintf("http://%s/%s", address, short)
	_, err = w.Write([]byte(shortURL))
	if err != nil {
		return
	}

}

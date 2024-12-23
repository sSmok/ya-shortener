package link

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mailru/easyjson"
	"github.com/sSmok/ya-shortener/internal/model"
)

// Shorten обрабатывает запросы на создание короткой ссылки
func (api *API) Shorten(w http.ResponseWriter, r *http.Request) {
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

	oLink := &model.OriginalLink{}
	err = easyjson.Unmarshal(body, oLink)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	short, err := api.linkRepo.Create(oLink.URL)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	sLink := &model.ShortLink{URL: fmt.Sprintf("%s/%s", api.baseURL, short)}

	marshal, err := easyjson.Marshal(sLink)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(marshal)
	if err != nil {
		return
	}
}

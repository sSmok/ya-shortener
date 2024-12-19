package link

import (
	"net/http"
)

// Short обрабатывает запросы по коротким ссылкам и возвращает оригинальную ссылку
func (api *API) Short(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	short := r.PathValue("id")
	originalURL, err := api.linkRepo.Get(short)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Location", originalURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

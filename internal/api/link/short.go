package link

import (
	"log"
	"net/http"
)

// Short обрабатывает запросы по коротким ссылкам и возвращает оригинальную ссылку
func (api *API) Short(w http.ResponseWriter, r *http.Request) {
	short := r.PathValue("id")
	log.Println("SHORT: ", short)
	originalURL, ok := api.storage[short]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Location", originalURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

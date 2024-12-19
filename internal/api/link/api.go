package link

import (
	"net/http"
	"time"
)

const address = "localhost:8080"

// API представляет собой обработчик запросов
type API struct {
	storage map[string]string
	mux     *http.ServeMux
}

// NewAPI создает новый экземпляр API
func NewAPI() *API {
	return &API{
		storage: make(map[string]string),
		mux:     http.NewServeMux(),
	}
}

// Run запускает сервер
func (api *API) Run() error {
	api.registerHandlers()

	server := &http.Server{
		Addr:              address,
		ReadHeaderTimeout: 2 * time.Second,
		Handler:           api.mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (api *API) registerHandlers() {
	api.mux.HandleFunc("/", api.Create)
	api.mux.HandleFunc("/{id}", api.Short)
}

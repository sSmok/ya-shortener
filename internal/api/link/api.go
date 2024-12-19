package link

import (
	"github.com/sSmok/ya-shortener/internal/repository"
)

//const address = "localhost:8080"

// API представляет собой обработчик запросов
type API struct {
	linkRepo repository.LinkRepository
	//mux      *http.ServeMux
}

// NewAPI создает новый экземпляр API
func NewAPI(linkRepo repository.LinkRepository) *API {
	return &API{
		linkRepo: linkRepo,
		//mux:      http.NewServeMux(),
	}
}

//// Run запускает сервер
//func (api *API) Run() error {
//	api.registerHandlers()
//
//	server := &http.Server{
//		Addr:              address,
//		ReadHeaderTimeout: 2 * time.Second,
//		Handler:           api.mux,
//	}
//
//	err := server.ListenAndServe()
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func (api *API) registerHandlers() {
//	api.mux.HandleFunc("/", api.Create)
//	api.mux.HandleFunc("/{id}", api.Short)
//}

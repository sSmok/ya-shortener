package app

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sSmok/ya-shortener/internal/config"
	"github.com/sSmok/ya-shortener/internal/middleware"
)

// App представляет собой приложение
type App struct {
	container  *container
	mux        *chi.Mux
	httpServer *http.Server
}

// NewApp создает новое приложение
func NewApp() (*App, error) {
	app := &App{}
	err := app.initDeps()
	if err != nil {
		return nil, err
	}
	return app, nil
}

// Run запускает http сервер
func (a *App) Run() error {
	log.Printf("server running on %s", a.container.AddressConfig().Address())
	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initDeps() error {
	deps := []func() error{
		a.initConfig,
		a.initContainer,
		a.initMux,
		a.initHTTPServer,
	}

	for _, f := range deps {
		err := f()
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initContainer() error {
	a.container = newContainer()
	return nil
}

func (a *App) initMux() error {
	a.mux = chi.NewRouter()

	a.mux.Use(middleware.Log)

	a.mux.Post("/", a.container.LinkAPI().Create)
	a.mux.Get("/{id}", a.container.LinkAPI().Short)
	return nil
}

func (a *App) initHTTPServer() error {
	a.httpServer = &http.Server{
		Addr:              a.container.AddressConfig().Address(),
		ReadHeaderTimeout: 2 * time.Second,
		Handler:           a.mux,
	}

	return nil
}

func (a *App) initConfig() error {
	return config.Load()
}

package app

import (
	"net/http"
	"time"
)

const address = "localhost:8080"

// App представляет собой приложение
type App struct {
	container  *container
	mux        *http.ServeMux
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
	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initDeps() error {
	deps := []func() error{
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
	a.mux = http.NewServeMux()
	a.mux.HandleFunc("/", a.container.LinkAPI().Create)
	a.mux.HandleFunc("/{id}", a.container.LinkAPI().Short)
	return nil
}

func (a *App) initHTTPServer() error {
	a.httpServer = &http.Server{
		Addr:              address,
		ReadHeaderTimeout: 2 * time.Second,
		Handler:           a.mux,
	}

	return nil
}

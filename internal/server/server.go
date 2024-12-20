package server

import (
	"fmt"
	"net/http"
)

type AppConfig struct {
	Port string
}

func Start(appConfig AppConfig) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", appConfig.Port),
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

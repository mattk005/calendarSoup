package httpserver

import (
	"log"
	"net/http"

	"wfm-json-to-ics/internal/config"
)

func StartServer(cfg *config.Config) {
	lCfg := newLocalCfg(cfg)
	port := "8080"
	mux := http.NewServeMux()
	// let apache 2 serve the files
	// filePathRoot := "."
	// fsHanlder := http.StripPrefix("/app/", http.FileServer(http.Dir(filePathRoot)))
	// mux.Handle("/app/", fsHanlder)

	mux.HandleFunc("GET /app/getCal/{wfmID}", lCfg.handlerGetCal)
	mux.HandleFunc("POST /api/refreshCal", lCfg.handlerRefreshCal)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	log.Fatal(srv.ListenAndServe())
}

type LocalCfg struct {
	*config.Config // promoted
}

func newLocalCfg(cfg *config.Config) *LocalCfg {
	return &LocalCfg{Config: cfg}
}

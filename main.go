package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/muzavan/the-versus/league"
	"github.com/muzavan/the-versus/sql"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout)
	sqlStore := &sql.Store{}
	leagueService, err := league.NewService(sqlStore, sqlStore)
	if err != nil {
		logger.Fatal().Err(err).Msg("league.NewService fails")
	}

	handler := initHandler(leagueService)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		logger.Error().Err(err).Msg("http.ListenAndServe fails")
	}
}

func initHandler(service *league.Service) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/v1/competition", nil).Methods("POST")
	router.HandleFunc("/v1/competition/{id}/matches", nil).Methods("GET")
	router.HandleFunc("/v1/competition/{id}/table", nil).Methods("GET")
	router.HandleFunc("/v1/competition/{id}/report", nil).Methods("POST")
	return router
}

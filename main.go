package main

import (
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	"github.com/muzavan/the-versus/league"
	"github.com/muzavan/the-versus/sql"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout)
	sqlStore := &sql.Store{}
	leagueServer, err := league.NewServer(sqlStore, sqlStore, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("league.NewServer fails")
	}

	handler := initHandler(leagueServer)
	logger.Info().Msg("Starting the leagueServer")
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() {
		err := server.ListenAndServe()
		logger.Error().Err(err).Msg("http.ListenAndServe fails")
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	server.Close()

	logger.Info().Msg("Done!")
}

func initHandler(service *league.Server) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/v1/competition", service.CreateCompetition).Methods("POST")
	router.HandleFunc("/v1/competition/{id}/matches", service.Matches).Methods("GET")
	router.HandleFunc("/v1/competition/{id}/table", service.Table).Methods("GET")
	router.HandleFunc("/v1/competition/{id}/report", service.Report).Methods("POST")
	return router
}

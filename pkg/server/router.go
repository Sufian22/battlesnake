package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sufian22/battlesnake/lib/api/middlewares"
	"github.com/sufian22/battlesnake/pkg/server/handlers"
)

const (
	GetInfoPath   = "/"
	StartGamePath = "/start"
	MovePath      = "/move"
	EndGamePath   = "/end"
	MetricsPath   = "/metrics"
)

func NewRouter(bs *BattlesnakeServer) http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middlewares.MetricsMiddleware, middlewares.LoggingMiddleware(bs.logger))

	router.HandleFunc(GetInfoPath, handlers.GetInfoHandler(bs.logger, bs.config)).Methods("GET")
	router.HandleFunc(StartGamePath, handlers.StartGameHandler(bs.logger, &bs.games)).Methods("POST")
	router.HandleFunc(MovePath, handlers.MoveHandler(bs.logger, &bs.games)).Methods("POST")
	router.HandleFunc(EndGamePath, handlers.EndGameHandler(bs.logger, &bs.games)).Methods("POST")
	router.Handle(MetricsPath, promhttp.Handler()).Methods("GET")

	return router
}

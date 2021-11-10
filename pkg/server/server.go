package server

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/config"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/server/handlers"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/server/middlewares"
)

type BattlesnakeServer struct {
	config config.ServerConfig
	games  sync.Map
	server *http.Server
	logger *logrus.Logger
}

func NewBattlesnakeServerFunc(logger *logrus.Logger) func(config.ServerConfig) (*BattlesnakeServer, error) {
	return func(config config.ServerConfig) (*BattlesnakeServer, error) {
		bs := &BattlesnakeServer{
			config: config,
			games:  sync.Map{},
			logger: logger,
		}

		bs.server = &http.Server{
			Addr:         config.Port,
			Handler:      configureRouter(bs),
			ReadTimeout:  time.Second * 15,
			WriteTimeout: time.Second * 15,
			IdleTimeout:  time.Second * 60,
		}

		return bs, nil
	}
}

func configureRouter(bs *BattlesnakeServer) http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middlewares.MetricsMiddleware, middlewares.LoggingMiddleware(bs.logger))

	router.HandleFunc("/", handlers.GetInfoHandler(bs.logger, bs.config)).Methods("GET")
	router.HandleFunc("/start", handlers.StartGameHandler(bs.logger, &bs.games)).Methods("POST")
	router.HandleFunc("/move", handlers.MoveHandler(bs.logger, &bs.games)).Methods("POST")
	router.HandleFunc("/end", handlers.EndGameHandler(bs.logger, &bs.games)).Methods("POST")
	router.Handle("/metrics", promhttp.Handler()).Methods("GET")

	return router
}

func (bs *BattlesnakeServer) Start() error {
	bs.logger.Printf("server listening on port %s", bs.server.Addr)
	if err := bs.server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (bs *BattlesnakeServer) Shutdown(ctx context.Context) {
	bs.server.Shutdown(ctx)
}

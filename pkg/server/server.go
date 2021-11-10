package server

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/config"
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

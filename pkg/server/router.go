package server

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"

	"github.com/BattlesnakeOfficial/rules/cli/commands"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/config"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/errors"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/middlewares"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/types"
)

func configureRouter(bs *BattlesnakeServer) http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middlewares.LoggingMiddleware(bs.logger))

	router.Handle("/metrics", promhttp.Handler()).Methods("GET")

	router.HandleFunc("/", GetInfoHandler(bs.logger, bs.config)).Methods("GET")
	router.HandleFunc("/start", StartGameHandler(bs.logger, &bs.games)).Methods("POST")
	router.HandleFunc("/move", MoveHandler(bs.logger, &bs.games)).Methods("POST")
	router.HandleFunc("/end", EndGameHandler(bs.logger, &bs.games)).Methods("POST")

	return router
}

func GetInfoHandler(logger *logrus.Logger, config config.ServerConfig) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(config.Info)
	}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func StartGameHandler(logger *logrus.Logger, games *sync.Map) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req commands.ResponsePayload
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{err.Error()})
			return
		}

		// game types not typed
		if req.Game.Ruleset.Name != "solo" {
			err := errors.UnsupportedGameModeErr{}
			logger.WithFields(logrus.Fields{
				"action": types.START_GAME,
				"mode":   req.Game.Ruleset.Name,
			}).Error(err.Error())

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{err.Error()})
			return
		}

		gameID := req.Game.Id
		if _, ok := games.Load(gameID); ok {
			err := errors.GameAlreadyStartedErr{}
			logger.WithFields(logrus.Fields{
				"action": types.START_GAME,
				"gameID": gameID,
			}).Error(err.Error(gameID))

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{err.Error(gameID)})
			return
		}

		games.Store(gameID, &commands.ResponsePayload{
			Game:  req.Game,
			Turn:  0, // ignore turn from request because its a new game
			Board: req.Board,
			You:   req.You,
		})

		logger.WithFields(logrus.Fields{
			"action": types.START_GAME,
			"gameID": gameID,
			"mode":   req.Game.Ruleset.Name,
		}).Info()

		w.WriteHeader(http.StatusOK)
	}
}

func MoveHandler(logger *logrus.Logger, games *sync.Map) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req commands.ResponsePayload
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		gameID := req.Game.Id
		if _, ok := games.Load(gameID); !ok {
			err := errors.UnknownGameErr{}
			logger.WithFields(logrus.Fields{
				"action": types.END_GAME,
			}).Error(err.Error(gameID))

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{err.Error(gameID)})
			return
		}

		// movements not typed
		validMoves := []string{"up", "down", "left", "right"}
		// random movements
		movement := validMoves[rand.Intn(len(validMoves))]

		logger.WithFields(logrus.Fields{
			"action":   types.MOVE,
			"snakeID":  req.You.Id,
			"movement": movement,
		}).Info()

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(commands.PlayerResponse{
			Move:  string(movement),
			Shout: fmt.Sprintf("moving %s!", movement),
		})
	}
}

func EndGameHandler(logger *logrus.Logger, games *sync.Map) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req commands.ResponsePayload
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{err.Error()})
			return
		}

		gameID := req.Game.Id
		val, ok := games.Load(gameID)
		if !ok {
			err := errors.UnknownGameErr{}
			logger.WithFields(logrus.Fields{
				"action": types.END_GAME,
			}).Error(err.Error(gameID))

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{err.Error(gameID)})
			return
		}

		game := val.(*commands.ResponsePayload)
		// evaluate game results
		if len(game.Board.Food) > 0 {
		}

		games.Delete(gameID)

		logger.WithFields(logrus.Fields{
			"action": types.END_GAME,
			"gameID": gameID,
		}).Info()

		w.WriteHeader(http.StatusOK)
	}
}

package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/config"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/server/errors"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/server/metrics"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/server/models"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/types"
)

func GetInfoHandler(logger *logrus.Logger, config config.ServerConfig) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(config.Info)
	}
}

func StartGameHandler(logger *logrus.Logger, games *sync.Map) func(w http.ResponseWriter, r *http.Request) {
	action := types.START_GAME
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.StartGameRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		if err := req.Game.Ruleset.Name.IsSupported(); err != nil {
			logger.WithFields(logrus.Fields{
				"action": action,
				"mode":   req.Game.Ruleset.Name,
			}).Error(err.Error())

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		gameID := req.Game.ID
		if _, ok := games.Load(gameID); ok {
			err := errors.GameAlreadyStartedErr{}
			logger.WithFields(logrus.Fields{
				"action": action,
				"gameID": gameID,
			}).Error(err.Error(gameID))

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Error: err.Error(gameID),
			})
			return
		}

		games.Store(gameID, &models.GameRequest{
			Game:  req.Game,
			Turn:  0, // ignore turn from request because its a new game
			Board: req.Board,
			You:   req.You,
		})

		logger.WithFields(logrus.Fields{
			"action": action,
			"gameID": gameID,
			"mode":   req.Game.Ruleset.Name,
		}).Info()

		w.WriteHeader(http.StatusOK)
	}
}

func MoveHandler(logger *logrus.Logger, games *sync.Map) func(w http.ResponseWriter, r *http.Request) {
	action := types.MOVE
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.MoveRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		gameID := req.Game.ID
		if _, ok := games.Load(gameID); !ok {
			err := errors.UnknownGameErr{}
			logger.WithFields(logrus.Fields{
				"action": action,
			}).Error(err.Error(gameID))

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Error: err.Error(gameID),
			})
			return
		}

		// random movements
		movement := types.Movements[rand.Intn(len(types.Movements))]

		logger.WithFields(logrus.Fields{
			"action":   action,
			"snakeID":  req.You.ID,
			"movement": movement,
		}).Info()

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.MoveResponse{
			Move:  movement,
			Shout: fmt.Sprintf("moving %s!", movement),
		})
	}
}

func EndGameHandler(logger *logrus.Logger, games *sync.Map) func(w http.ResponseWriter, r *http.Request) {
	action := types.END_GAME
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.EndGameRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		gameID := req.Game.ID
		val, ok := games.Load(gameID)
		if !ok {
			err := errors.UnknownGameErr{}
			logger.WithFields(logrus.Fields{
				"action": action,
			}).Error(err.Error(gameID))

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Error: err.Error(gameID),
			})
			return
		}

		game := val.(*models.GameRequest)
		// evaluate game results
		if game.Game.Ruleset.Name == types.Solo {
			// snake performance on games
			if len(game.Board.Food) > 0 {
				metrics.TotalLosses.With(prometheus.Labels{"snakeID": game.You.ID}).Inc()
			} else {
				metrics.TotalWins.With(prometheus.Labels{"snakeID": game.You.ID}).Inc()
			}
		}

		games.Delete(gameID)

		logger.WithFields(logrus.Fields{
			"action": action,
			"gameID": gameID,
		}).Info()

		w.WriteHeader(http.StatusOK)
	}
}

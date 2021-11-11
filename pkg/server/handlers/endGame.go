package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/server/errors"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/server/metrics"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/server/models"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/types"
)

func EndGameHandler(logger *logrus.Logger, games *sync.Map) func(w http.ResponseWriter, r *http.Request) {
	loggerEntry := logger.WithFields(logrus.Fields{
		"action": types.END_GAME,
	})

	return func(w http.ResponseWriter, r *http.Request) {
		var req models.EndGameRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			loggerEntry.Error(err.Error())
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
			loggerEntry.Error(err.Error(gameID))

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

		loggerEntry.WithFields(logrus.Fields{
			"gameID": gameID,
		}).Info()

		w.WriteHeader(http.StatusOK)
	}
}

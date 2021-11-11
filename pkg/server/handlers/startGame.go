package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/server/errors"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/server/models"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/types"
)

func StartGameHandler(logger *logrus.Logger, games *sync.Map) func(w http.ResponseWriter, r *http.Request) {
	loggerEntry := logger.WithFields(logrus.Fields{
		"action": types.START_GAME,
	})

	return func(w http.ResponseWriter, r *http.Request) {
		var req models.StartGameRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			loggerEntry.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		gameID := req.Game.ID
		loggerEntry = loggerEntry.WithFields(logrus.Fields{
			"gameID": gameID,
		})

		if _, ok := games.Load(gameID); ok {
			err := errors.GameAlreadyStartedErr{}
			loggerEntry.Error(err.Error(gameID))

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

		loggerEntry.WithFields(logrus.Fields{
			"mode": req.Game.Ruleset.Name,
		}).Info()

		w.WriteHeader(http.StatusOK)
	}
}

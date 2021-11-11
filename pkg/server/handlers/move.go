package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/server/errors"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/server/models"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/types"
)

func MoveHandler(logger *logrus.Logger, games *sync.Map) func(w http.ResponseWriter, r *http.Request) {
	loggerEntry := logger.WithFields(logrus.Fields{
		"action": types.MOVE,
	})

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
			loggerEntry.Error(err.Error(gameID))

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Error: err.Error(gameID),
			})
			return
		}

		loggerEntry = loggerEntry.WithFields(logrus.Fields{
			"snakeID": req.You.ID,
			"gameID":  req.Game.ID,
		})

		movement := calculateNextMovement(req)

		if movement == "" {
			err := errors.NoMovementErr{}
			loggerEntry.Warn(err.Error())
		} else {
			loggerEntry.WithFields(logrus.Fields{
				"movement": movement,
			}).Info()
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.MoveResponse{
			Move:  movement,
			Shout: fmt.Sprintf("moving %s!", movement),
		})
	}
}

func calculateNextMovement(req models.MoveRequest) types.Movement {
	var forbiddenCells types.Coordinates
	forbiddenCells = append(forbiddenCells, req.You.Body...)
	forbiddenCells = append(forbiddenCells, req.Board.Hazards...)

	if req.Game.Ruleset.Name != types.Solo && req.Game.Ruleset.Name != types.Standard {
		forbiddenCells = append(forbiddenCells, req.Board.Snakes...)
	}

	validMovements := []types.Movement{}
	for k, v := range types.ValidMovementCoordinates {
		nextCoordinate := types.Coordinate{
			X: req.You.Head.X + v.X,
			Y: req.You.Head.Y + v.Y,
		}

		if !nextCoordinate.IsValid(req.Board.Height, req.Board.Width) {
			continue
		}

		if isNextMovementForbid(forbiddenCells, nextCoordinate) {
			validMovements = append(validMovements, k)
		}
	}

	var movement types.Movement
	if len(validMovements) > 0 {
		movement = validMovements[rand.Intn(len(validMovements))]
	}

	return movement
}

func isNextMovementForbid(forbidden types.Coordinates, c types.Coordinate) bool {
	for _, v := range forbidden {
		if v == c {
			return false
		}
	}
	return true
}

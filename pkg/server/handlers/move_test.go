package handlers_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sufian22/battlesnake/pkg/models"
	"github.com/sufian22/battlesnake/pkg/server/handlers"
	httpModels "github.com/sufian22/battlesnake/pkg/server/models"
	"github.com/sufian22/battlesnake/pkg/types"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetOutput(ioutil.Discard)
}

func TestMoveHandler(t *testing.T) {
	games := sync.Map{}
	gameID := "123456"

	CreateGameForTest(logger, &games, gameID, models.Board{})

	tt := []struct {
		name             string
		gameID           string
		gameMode         types.GameMode
		you              models.Snake
		board            models.Board
		posibleMovements []types.Movement
		expectedStatus   int
	}{
		{"validGameButEmptyBoard", gameID, types.Solo, models.Snake{}, models.Board{}, []types.Movement{}, http.StatusOK},
		{
			name:     "validSnakePosition",
			gameID:   gameID,
			gameMode: types.Solo,
			you: models.Snake{
				Head: types.Coordinate{
					X: 0,
					Y: 0,
				},
			},
			board: models.Board{
				Height: 2,
				Width:  2,
			},
			posibleMovements: []types.Movement{types.Up, types.Right},
			expectedStatus:   http.StatusOK,
		},
		{
			name:     "validSnakePosition",
			gameID:   gameID,
			gameMode: types.Solo,
			you: models.Snake{
				Head: types.Coordinate{
					X: 0,
					Y: 1,
				},
			},
			board: models.Board{
				Height: 2,
				Width:  2,
			},
			posibleMovements: []types.Movement{types.Down, types.Right},
			expectedStatus:   http.StatusOK,
		},
		{
			name:     "validSnakePosition",
			gameID:   gameID,
			gameMode: types.Solo,
			you: models.Snake{
				Head: types.Coordinate{
					X: 1,
					Y: 1,
				},
			},
			board: models.Board{
				Height: 2,
				Width:  2,
			},
			posibleMovements: []types.Movement{types.Down, types.Left},
			expectedStatus:   http.StatusOK,
		},
		{
			name:     "validSnakePosition",
			gameID:   gameID,
			gameMode: types.Solo,
			you: models.Snake{
				Head: types.Coordinate{
					X: 1,
					Y: 0,
				},
			},
			board: models.Board{
				Height: 2,
				Width:  2,
			},
			posibleMovements: []types.Movement{types.Up, types.Left},
			expectedStatus:   http.StatusOK,
		},
		{
			name:     "noMovementPosible",
			gameID:   gameID,
			gameMode: types.Solo,
			you: models.Snake{
				Head: types.Coordinate{
					X: 0,
					Y: 0,
				},
				Body: []types.Coordinate{
					{
						X: 0,
						Y: 1,
					},
					{
						X: 1,
						Y: 0,
					},
					{
						X: 1,
						Y: 1,
					},
				},
			},
			board: models.Board{
				Height: 2,
				Width:  2,
			},
			posibleMovements: []types.Movement{},
			expectedStatus:   http.StatusOK,
		},
		{
			name:     "noMovementPosible",
			gameID:   gameID,
			gameMode: types.Solo,
			you: models.Snake{
				Head: types.Coordinate{
					X: 0,
					Y: 0,
				},
			},
			board: models.Board{
				Height: 2,
				Width:  2,
				Hazards: []types.Coordinate{
					{
						X: 0,
						Y: 1,
					},
					{
						X: 1,
						Y: 0,
					},
					{
						X: 1,
						Y: 1,
					},
				},
			},
			posibleMovements: []types.Movement{},
			expectedStatus:   http.StatusOK,
		},
		{"invalidGameMode", gameID, types.GameMode("test"), models.Snake{}, models.Board{}, []types.Movement{}, http.StatusBadRequest},
		{"unknownGameID", "test", types.Solo, models.Snake{}, models.Board{}, []types.Movement{}, http.StatusBadRequest},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			reqBody := httpModels.MoveRequest{
				GameRequest: httpModels.GameRequest{
					Game: models.Game{
						ID:      tc.gameID,
						Source:  types.Custom,
						Timeout: 1,
						Ruleset: models.Ruleset{
							Name:    tc.gameMode,
							Version: "1",
						},
					},
					You:   tc.you,
					Board: tc.board,
				},
			}

			body, _ := json.Marshal(reqBody)
			req := httptest.NewRequest(http.MethodPost, "/move", bytes.NewBuffer(body))
			rr := httptest.NewRecorder()

			moveHandlerFunc := handlers.MoveHandler(logger, &games)
			handler := http.HandlerFunc(moveHandlerFunc)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tc.expectedStatus)
			}

			if tc.expectedStatus == http.StatusOK {
				_, ok := games.Load(gameID)
				if !ok {
					t.Errorf("game has not been stored correctly")
				}

				resp := httpModels.MoveResponse{}
				err := json.Unmarshal(rr.Body.Bytes(), &resp)
				if err != nil && len(tc.posibleMovements) > 0 {
					t.Error(err)
				}

				if len(tc.posibleMovements) > 0 {
					found := false
					for _, v := range tc.posibleMovements {
						if v == resp.Move {
							found = true
							break
						}
					}

					if !found {
						t.Errorf("unexpected movement in response got %s want any of %v", resp.Move, tc.posibleMovements)
					}
				}
			}
		})
	}
}

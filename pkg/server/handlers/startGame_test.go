package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sufian22/battlesnake/lib/battlesnake/models"
	"github.com/sufian22/battlesnake/lib/battlesnake/types"
	"github.com/sufian22/battlesnake/pkg/server/handlers"
	httpModels "github.com/sufian22/battlesnake/pkg/server/models"
)

func TestStartGameHandler(t *testing.T) {
	games := sync.Map{}
	gameID := "123456"

	tt := []struct {
		name           string
		source         types.Source
		gameMode       types.GameMode
		expectedStatus int
	}{
		{"validGame", types.Custom, types.Solo, http.StatusOK},
		{"unknownSource", types.Source("test"), types.Solo, http.StatusBadRequest},
		{"unknownGameMode", types.Custom, types.GameMode("test"), http.StatusBadRequest},
		{"unsupportedGameMode", types.Custom, types.Squad, http.StatusBadRequest},
		{"repeatedGame", types.Custom, types.Solo, http.StatusBadRequest},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			rr := CreateGameForTest(logger, &games, gameID, models.Board{})

			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tc.expectedStatus)
			}

			if tc.expectedStatus == http.StatusOK {
				_, ok := games.Load(gameID)
				if !ok {
					t.Errorf("game has not been stored correctly")
				}
			}
		})
	}
}

func CreateGameForTest(logger *logrus.Logger, games *sync.Map, gameID string, board models.Board) *httptest.ResponseRecorder {
	gameRequest := httpModels.GameRequest{
		Game: models.Game{
			ID:      gameID,
			Source:  types.Custom,
			Timeout: 1,
			Ruleset: models.Ruleset{
				Name:    types.Solo,
				Version: "1",
			},
		},
		Board: board,
	}

	body, _ := json.Marshal(httpModels.StartGameRequest{GameRequest: gameRequest})
	req := httptest.NewRequest(http.MethodPost, "/start", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	startGameHandlerFunc := handlers.StartGameHandler(logger, games)
	handler := http.HandlerFunc(startGameHandlerFunc)

	handler.ServeHTTP(rr, req)

	return rr
}

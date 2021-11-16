package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/sufian22/battlesnake/lib/battlesnake/models"
	"github.com/sufian22/battlesnake/lib/battlesnake/types"
	"github.com/sufian22/battlesnake/pkg/server/handlers"
	httpModels "github.com/sufian22/battlesnake/pkg/server/models"
)

func TestEndGameHandler(t *testing.T) {
	games := sync.Map{}

	gameIDWithFood := "123456"
	gameIDWithoutFood := "654321"

	_ = CreateGameForTest(logger, &games, gameIDWithFood, models.Board{Food: []types.Coordinate{{1, 1}}})
	_ = CreateGameForTest(logger, &games, gameIDWithoutFood, models.Board{})

	tt := []struct {
		name           string
		gameID         string
		source         types.Source
		gameMode       types.GameMode
		expectedStatus int
	}{
		{"validGameWithFood", gameIDWithFood, types.Custom, types.Solo, http.StatusOK},
		{"validGameWithoutFood", gameIDWithoutFood, types.Custom, types.Solo, http.StatusOK},
		{"unknownSource", gameIDWithFood, types.Source("test"), types.Solo, http.StatusBadRequest},
		{"unknownGameMode", gameIDWithFood, types.Custom, types.GameMode("test"), http.StatusBadRequest},
		{"unsupportedGameMode", gameIDWithFood, types.Custom, types.Squad, http.StatusBadRequest},
		{"repeatedGame", gameIDWithFood, types.Custom, types.Solo, http.StatusBadRequest},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			reqBody := httpModels.EndGameRequest{
				GameRequest: httpModels.GameRequest{
					Game: models.Game{
						ID:      tc.gameID,
						Source:  tc.source,
						Timeout: 1,
						Ruleset: models.Ruleset{
							Name:    tc.gameMode,
							Version: "1",
						},
					},
				},
			}

			body, _ := json.Marshal(reqBody)
			req := httptest.NewRequest(http.MethodPost, "/end", bytes.NewBuffer(body))
			rr := httptest.NewRecorder()

			endGameHandlerFunc := handlers.EndGameHandler(logger, &games)
			handler := http.HandlerFunc(endGameHandlerFunc)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tc.expectedStatus)
			}

			if tc.expectedStatus == http.StatusOK {
				_, ok := games.Load(tc.gameID)
				if ok {
					t.Errorf("game has not been deleted correctly")
				}
			}
		})
	}
}

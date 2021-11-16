package server_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sufian22/battlesnake/lib/battlesnake/models"
	"github.com/sufian22/battlesnake/lib/battlesnake/types"
	"github.com/sufian22/battlesnake/pkg/config"
	"github.com/sufian22/battlesnake/pkg/server"
	httpModels "github.com/sufian22/battlesnake/pkg/server/models"
)

func TestNewRouter(t *testing.T) {
	newServerFunc := server.NewBattlesnakeServerFunc(logrus.New())
	bsServer, err := newServerFunc(config.ServerConfig{})
	if err != nil {
		t.Error(err)
	}

	handler := server.NewRouter(bsServer)
	validGameRequest := httpModels.GameRequest{
		Game: models.Game{
			ID:      "1123445",
			Source:  types.Custom,
			Timeout: 1,
			Ruleset: models.Ruleset{
				Name:    types.Solo,
				Version: "1",
			},
		},
	}

	tt := []struct {
		name               string
		path               string
		body               interface{}
		methods            []string
		expectedStatusCode int
	}{
		{
			name:               "getInfoRouteFail",
			path:               server.GetInfoPath,
			body:               nil,
			methods:            []string{http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
			expectedStatusCode: http.StatusMethodNotAllowed,
		},
		{
			name:               "getInfoRouteSuccess",
			path:               server.GetInfoPath,
			body:               nil,
			methods:            []string{http.MethodGet},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "startGameRouteFail",
			path:               server.StartGamePath,
			body:               validGameRequest,
			methods:            []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete},
			expectedStatusCode: http.StatusMethodNotAllowed,
		},
		{
			name:               "startGameRouteSuccess",
			path:               server.StartGamePath,
			body:               validGameRequest,
			methods:            []string{http.MethodPost},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "moveRouteFail",
			path:               server.MovePath,
			body:               validGameRequest,
			methods:            []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete},
			expectedStatusCode: http.StatusMethodNotAllowed,
		},
		{
			name:               "moveRouteSuccess",
			path:               server.MovePath,
			body:               validGameRequest,
			methods:            []string{http.MethodPost},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "endGameRouteFail",
			path:               server.EndGamePath,
			body:               validGameRequest,
			methods:            []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete},
			expectedStatusCode: http.StatusMethodNotAllowed,
		},
		{
			name:               "endGameRouteFail",
			path:               server.EndGamePath,
			body:               validGameRequest,
			methods:            []string{http.MethodPost},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			for _, method := range tc.methods {
				body, _ := json.Marshal(tc.body)

				req := httptest.NewRequest(method, tc.path, bytes.NewBuffer(body))
				rr := httptest.NewRecorder()

				handler.ServeHTTP(rr, req)

				if rr.Code != tc.expectedStatusCode {
					t.Errorf("unexpected status code in response got %v expected %v", rr.Code, tc.expectedStatusCode)
				}
			}
		})
	}
}

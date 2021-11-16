package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/sufian22/battlesnake/pkg/config"
	"github.com/sufian22/battlesnake/pkg/server/handlers"
)

func TestGetInfoHandler(t *testing.T) {
	serverConfig := config.ServerConfig{
		Info: config.BattlesnakeInfo{
			APIVersion: "1",
			Author:     "test",
			Color:      "#000000",
			Head:       "default",
			Tail:       "default",
			Version:    "v1",
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	infoHandlerFunc := handlers.GetInfoHandler(logger, serverConfig)
	handler := http.HandlerFunc(infoHandlerFunc)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	infoReceived := config.BattlesnakeInfo{}
	if err := json.Unmarshal(rr.Body.Bytes(), &infoReceived); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(serverConfig.Info, infoReceived) {
		t.Errorf("handler returned unexpected body: got %v want %v", infoReceived, serverConfig.Info)
	}
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/sufian22/battlesnake/pkg/config"
)

func GetInfoHandler(logger *logrus.Logger, config config.ServerConfig) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(config.Info)
	}
}

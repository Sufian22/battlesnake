package models

import "github.com/sufian22/battlesnake/lib/battlesnake/models"

type GameRequest struct {
	Game  models.Game  `json:"game"`
	Turn  int32        `json:"turn"`
	Board models.Board `json:"board"`
	You   models.Snake `json:"you"`
}

type StartGameRequest struct {
	GameRequest
}

type MoveRequest struct {
	GameRequest
}

type EndGameRequest struct {
	GameRequest
}

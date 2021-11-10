package models

import "gitlab.com/etoshi/testingroom/battlesnake/pkg/models"

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

package models

import (
	"github.com/sufian22/battlesnake/pkg/config"
	"github.com/sufian22/battlesnake/pkg/types"
)

type MoveResponse struct {
	Move  types.Movement `json:"move"`
	Shout string         `json:"shout"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type GetInfoResponse struct {
	config.BattlesnakeInfo
}

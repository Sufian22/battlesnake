package models

import (
	"github.com/sufian22/battlesnake/lib/battlesnake/types"
	"github.com/sufian22/battlesnake/pkg/config"
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

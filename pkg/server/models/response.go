package models

import (
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/config"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/types"
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

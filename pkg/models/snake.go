package models

import "gitlab.com/etoshi/testingroom/battlesnake/pkg/types"

type Snake struct {
	ID      string            `json:"id"`
	Name    string            `json:"name"`
	Health  int32             `json:"health"`
	Body    types.Coordinates `json:"body"`
	Latency string            `json:"latency"`
	Head    types.Coordinate  `json:"head"`
	Length  int32             `json:"length"`
	Shout   string            `json:"shout"`
	Squad   string            `json:"squad"`
}

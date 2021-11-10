package models

import "gitlab.com/etoshi/testingroom/battlesnake/pkg/types"

type Board struct {
	Height  int32             `json:"height"`
	Width   int32             `json:"width"`
	Food    types.Coordinates `json:"food"`
	Hazards types.Coordinates `json:"hazards"`
	Snakes  types.Coordinates `json:"snakes"`
}

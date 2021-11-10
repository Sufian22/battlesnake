package models

import "gitlab.com/etoshi/testingroom/battlesnake/pkg/types"

type Game struct {
	ID      string       `json:"id"`
	Ruleset Ruleset      `json:"ruleset"`
	Timeout int32        `json:"timeout"`
	Source  types.Source `json:"source"`
}

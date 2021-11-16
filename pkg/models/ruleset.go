package models

import "github.com/sufian22/battlesnake/pkg/types"

type Ruleset struct {
	Name     types.GameMode  `json:"name"`
	Version  string          `json:"version"`
	Settings RulesetSettings `json:"settings"`
}

type RulesetSettings struct {
	StandardSettings
	Royale RoyaleSettings `json:"royale"`
	Squad  SquadSettings  `json:"squad"`
}

type StandardSettings struct {
	FoodSpawnChance     int32 `json:"foodSpawnChance"`
	MinimumFood         int32 `json:"minimumFood"`
	HazardDamagePerTurn int32 `json:"hazardDamagePerTurn"`
}

type RoyaleSettings struct {
	ShrinkEveryNTurns int32 `json:"shrinkEveryNTurns"`
}

type SquadSettings struct {
	AllowBodyCollisions bool `json:"allowBodyCollisions"`
	SharedElimination   bool `json:"sharedElimination"`
	SharedHealth        bool `json:"sharedHealth"`
	SharedLength        bool `json:"sharedLength"`
}

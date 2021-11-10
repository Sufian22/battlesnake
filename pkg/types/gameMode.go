package types

import (
	"encoding/json"

	"gitlab.com/etoshi/testingroom/battlesnake/pkg/server/errors"
)

type GameMode string

const (
	Standard GameMode = "standard"
	Solo     GameMode = "solo"
	Royale   GameMode = "royale"
	Squad    GameMode = "squad"
)

func (gm *GameMode) UnmarshalJSON(b []byte) error {
	var value string
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	mode := GameMode(value)
	if err := mode.IsSupported(); err != nil {
		return err
	}

	*gm = mode
	return nil
}

func (gm *GameMode) IsSupported() error {
	switch *gm {
	case Solo:
		return nil
	}

	return &errors.UnsupportedGameModeErr{}
}

package types

import (
	"encoding/json"

	"gitlab.com/etoshi/testingroom/battlesnake/pkg/server/errors"
)

type Movement string

const (
	Up    Movement = "up"
	Down  Movement = "down"
	Left  Movement = "left"
	Right Movement = "right"
)

var Movements = []Movement{Up, Down, Left, Right}

func (m *Movement) UnmarshalJSON(b []byte) error {
	var value string
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	move := Movement(value)
	if err := move.IsValid(); err != nil {
		return err
	}

	*m = move
	return nil
}

func (m *Movement) IsValid() error {
	switch *m {
	case Up, Down, Left, Right:
		return nil
	}
	return &errors.InvalidMovementErr{}
}

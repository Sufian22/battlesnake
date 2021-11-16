package types

import (
	"encoding/json"

	"github.com/sufian22/battlesnake/pkg/server/errors"
)

type Movement string

const (
	Up    Movement = "up"
	Down  Movement = "down"
	Left  Movement = "left"
	Right Movement = "right"
)

var ValidMovementCoordinates map[Movement]Coordinate = map[Movement]Coordinate{
	Up: {
		X: 0,
		Y: 1,
	},
	Down: {
		X: 0,
		Y: -1,
	},
	Left: {
		X: -1,
		Y: 0,
	},
	Right: {
		X: 1,
		Y: 0,
	},
}

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
	_, ok := ValidMovementCoordinates[*m]
	if ok {
		return nil
	}
	return &errors.InvalidMovementErr{}
}

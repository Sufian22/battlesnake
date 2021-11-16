package types

import (
	"encoding/json"

	"github.com/sufian22/battlesnake/pkg/server/errors"
)

type Source string

const (
	League Source = "league"
	Custom Source = "custom"
)

func (s *Source) UnmarshalJSON(b []byte) error {
	var value string
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	source := Source(value)
	if err := source.IsValid(); err != nil {
		return err
	}

	*s = source
	return nil
}

func (s *Source) IsValid() error {
	switch *s {
	case League, Custom:
		return nil
	}
	return &errors.InvalidSourceErr{}
}

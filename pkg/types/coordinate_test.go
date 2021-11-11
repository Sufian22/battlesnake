package types_test

import (
	"testing"

	"gitlab.com/etoshi/testingroom/battlesnake/pkg/types"
)

func TestIsValid(t *testing.T) {
	validCoordinate := types.Coordinate{
		X: 1,
		Y: 1,
	}

	if !validCoordinate.IsValid(2, 2) {
		t.Error("coordinate expected to be valid")
	}

	invalidCoordinate := types.Coordinate{
		X: -1,
		Y: 0,
	}

	if invalidCoordinate.IsValid(2, 2) {
		t.Error("coordinate expected to be invalid")
	}
}

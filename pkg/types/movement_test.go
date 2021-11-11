package types_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"gitlab.com/etoshi/testingroom/battlesnake/pkg/server/models"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/types"
)

func TestMovementUnmarshalJSON(t *testing.T) {
	resp := models.MoveResponse{}
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"move": "%s"}`, types.Up)), &resp); err != nil {
		t.Error("it should be valid movemnt")
	}

	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"move": "%s"}`, types.Movement("test"))), &resp); err == nil {
		t.Error("it should be an invalid movement")
	}
}

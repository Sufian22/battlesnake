package types_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/sufian22/battlesnake/lib/battlesnake/types"
	"github.com/sufian22/battlesnake/pkg/server/models"
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

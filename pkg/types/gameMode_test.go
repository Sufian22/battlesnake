package types_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/sufian22/battlesnake/pkg/models"
	"github.com/sufian22/battlesnake/pkg/types"
)

func TestGameModeUnmarshalJSON(t *testing.T) {
	ruleset := models.Ruleset{}
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"name": "%s"}`, types.Solo)), &ruleset); err != nil {
		t.Error("it should be valid game mode")
	}

	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"name": "%s"}`, types.Squad)), &ruleset); err == nil {
		t.Error("it should be an invalid game mode")
	}
}

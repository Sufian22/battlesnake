package types_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"gitlab.com/etoshi/testingroom/battlesnake/pkg/models"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/types"
)

func TestSourceUnmarshalJSON(t *testing.T) {
	game := models.Game{}
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"source": "%s"}`, types.Custom)), &game); err != nil {
		t.Error("it should be valid source")
	}

	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"source": "%s"}`, types.Source("test"))), &game); err == nil {
		t.Error("it should be an invalid source")
	}
}

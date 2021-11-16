package types_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/sufian22/battlesnake/pkg/models"
	"github.com/sufian22/battlesnake/pkg/types"
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

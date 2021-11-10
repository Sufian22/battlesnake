package types

type Action string

const (
	START_GAME Action = "start-game"
	END_GAME   Action = "end-game"
	MOVE       Action = "move"
)

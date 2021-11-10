package errors

import "fmt"

type InvalidMovementErr struct{}

func (im *InvalidMovementErr) Error() string {
	return "invalid movement"
}

type InvalidSourceErr struct{}

func (is *InvalidSourceErr) Error() string {
	return "invalid source"
}

type InvalidModeErr struct{}

func (im *InvalidModeErr) Error() string {
	return "invalid game mode"
}

type InvalidHealthErr struct{}

func (ih *InvalidHealthErr) Error() string {
	return "invalid health value"
}

type InvalidCoordinateErr struct{}

func (ic *InvalidCoordinateErr) Error() string {
	return "invalid coordinate values"
}

type UnsupportedGameModeErr struct{}

func (ug *UnsupportedGameModeErr) Error() string {
	return "game mode not supported"
}

type UnknownGameModeErr struct{}

func (ug *UnknownGameModeErr) Error() string {
	return "unknown game mode"
}

type UnknownGameErr struct{}

func (ug *UnknownGameErr) Error(id string) string {
	return fmt.Sprintf("unknown game with id %s", id)
}

type GameAlreadyStartedErr struct{}

func (ga *GameAlreadyStartedErr) Error(id string) string {
	return fmt.Sprintf("game with id %s has already been started", id)
}

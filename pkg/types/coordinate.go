package types

type Coordinate struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

type Coordinates []Coordinate

func (c *Coordinate) IsValid() bool {
	return c.X >= 0 && c.Y >= 00
}

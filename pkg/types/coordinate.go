package types

type Coordinate struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

type Coordinates []Coordinate

func (c *Coordinate) IsValid(height, width int32) bool {
	return c.X >= 0 && c.Y >= 00 && c.X < width && c.Y < height
}

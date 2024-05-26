package dao

import "fmt"

type Position struct {
	X         int
	Y         int
	Direction string
}

func (p Position) String() string {
	return fmt.Sprintf("%v:(%v, %v )", p.Direction, p.X, p.Y)
}

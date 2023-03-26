package direction

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	s := [...]string{"North", "East", "South", "West"}
	return s[d]
}

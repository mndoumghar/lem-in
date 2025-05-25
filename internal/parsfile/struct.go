package parsfile

type Room struct {
	Name    string
	X, Y    int
	Links   []*Room
	IsStart bool
	Isend   bool
}

type Ant struct {
	ID   int
	Path []*Room
}

type Graph struct {
	Rooms map[string]*Room
	// Add AntCount if you want to store the number of ants in the graph.
	AntCount int
}

func (g *Graph) AddRoom(name string, x, y int, isStart, isEnd bool) {
	if g.Rooms == nil {
		g.Rooms = make(map[string]*Room)
	}
	g.Rooms[name] = &Room{
		Name:    name,
		X:       x,
		Y:       y,
		IsStart: isStart,
		Isend:   isEnd,
	}
}

func (g *Graph) AddLink(name1, name2 string) {
	room1 := g.Rooms[name1]
	room2 := g.Rooms[name2]

	if room1 != nil && room2 != nil {
		room1.Links = append(room1.Links, room2)
		room2.Links = append(room2.Links, room1)
	}
}
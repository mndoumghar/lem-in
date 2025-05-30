package algo

import (
	"fmt"
	"lemin/internal/parsfile"
)

// roomInPath checks if a room is already in the given path (to avoid cycles).
func roomInPath(path []*parsfile.Room, room *parsfile.Room) bool {
	for _, r := range path {
		if r.Name == room.Name {
			return true
		}
	}
	return false
}


func FindShortestPath(start *parsfile.Room, end *parsfile.Room) []*parsfile.Room {
	type state struct {
		Room *parsfile.Room
		Path []*parsfile.Room
	}
	visited := make(map[string]bool)
	queue := []state{{Room: start, Path: []*parsfile.Room{start}}}
	visited[start.Name] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Room == end {
			return current.Path
		}

		for _, neighbor := range current.Room.Links {
			if visited[neighbor.Name] || roomInPath(current.Path, neighbor) {
				continue
			}
			visited[neighbor.Name] = true
			newPath := append([]*parsfile.Room{}, current.Path...)
			newPath = append(newPath, neighbor)
			queue = append(queue, state{Room: neighbor, Path: newPath})
		}
	}
	return nil
}











// tedting ...

func SimulateAnts(ants int, paths [][]string) {
	var antList []Ant
	antID := 1

	for i := 0; i < ants; i++ {
		path := paths[i % len(paths)] 
		antList = append(antList, Ant{
			ID: antID,
			Path: path,
			Pos: 0, 
		})
		antID++
	}

	fmt.Println("Simulation starts:")
	for {
		done := true
		for i := 0; i < len(antList); i++ {
			if antList[i].Pos < len(antList[i].Path)-1 {
				antList[i].Pos++
				fmt.Printf("L%d-%s ", antList[i].ID, antList[i].Path[antList[i].Pos])
				done = false
			}
		}
		fmt.Println()
		if done {
			break
		}
	}
}

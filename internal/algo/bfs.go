package algo

import "lemin/internal/parsfile"

// roomInPath checks if a room is already in the given path (to avoid cycles).
func roomInPath(path []*parsfile.Room, room *parsfile.Room) bool {
	for _, r := range path {
		if r.Name == room.Name {
			return true
		}
	}
	return false
}

// FindShortestPath finds the shortest path from start to end using BFS.
// Returns the path as a slice of *parsfile.Room, or nil if no path is found.
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
package parsfile

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadLines(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	return lines, nil
}

func ParseFile(path string) (*Graph, int, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, 0, err
	}

	g := &Graph{Rooms: make(map[string]*Room)}

	var ants int
	var isStart, isEnd bool

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		if line == "" || (strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##")) {
			continue // skip commenter ##
		}

		if ants == 0 && isDigit(line) {
			ants, _ = strconv.Atoi(line)
			continue
		}

		if line == "##start" {
			isStart = true
			continue
		} else if line == "##end" {
			isEnd = true
			continue
		}

		if strings.Contains(line, " ") {
			parts := strings.Split(line, " ")
			if len(parts) != 3 {
				return nil, 0, fmt.Errorf("ERROR: invalid room format")
			}
			name := parts[0]
			x, _ := strconv.Atoi(parts[1])
			y, _ := strconv.Atoi(parts[2])

			g.AddRoom(name, x, y, isStart, isEnd)

			if isStart {
				g.Start = g.Rooms[name]
			}
			if isEnd {
				g.End = g.Rooms[name]
			}

			isStart, isEnd = false, false
		} else if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, 0, fmt.Errorf("ERROR: invalid link format")
			}
			g.AddLink(parts[0], parts[1])
		} else {
			return nil, 0, fmt.Errorf("ERROR: invalid data format")
		}
	}

	return g, ants, nil
}

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// 🧠 دالة DFS لإيجاد جميع المسارات
func FindAllPathsDFS(graph *Graph, startName, endName string) [][]string {
	var paths [][]string
	visited := make(map[string]bool)
	var currentPath []string

	var dfs func(room *Room)

	dfs = func(room *Room) {
		if visited[room.Name] {
			return
		}

		visited[room.Name] = true
		currentPath = append(currentPath, room.Name)

		if room.Name == endName {
			pathCopy := make([]string, len(currentPath))
			copy(pathCopy, currentPath)
			paths = append(paths, pathCopy)
		} else {
			for _, neighbor := range room.Links {
				dfs(neighbor)
			}
		}

		// backtrack
		currentPath = currentPath[:len(currentPath)-1]
		visited[room.Name] = false
	}

	startRoom := graph.Rooms[startName]
	dfs(startRoom)

	return paths
}

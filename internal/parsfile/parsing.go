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

		if line == "" || strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##") {
			continue // تجاهل التعاليق
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


package ants

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ants struct {
	AntNum         int
	RoomsWithCords map[string][]string
	Start          string
	End            string
	Links          []string
	Err            error
}



func ReadFile(filename string) Ants {
	var ant Ants

	data, err := os.ReadFile(filename)
	if err != nil {
		ant.Err = err
		return ant
	}
	fmt.Println(string(data) + "\n")

	rawLines := strings.Split(string(data), "\n")
	lines := make([]string, 0, len(rawLines))
	for _, l := range rawLines {
		lines = append(lines, strings.TrimSpace(l))
	}

	if len(lines) == 0 {
		ant.Err = fmt.Errorf("invalid data format, empty file")
		return ant
	}
	ant.AntNum, err = strconv.Atoi(lines[0])
	if err != nil {
		ant.Err = err
		return ant
	}
	if ant.AntNum <= 0 {
		ant.Err = fmt.Errorf("invalid data format, no ant in the start room")
		return ant
	}

	ant.RoomsWithCords = make(map[string][]string)
	ant.Links = []string{}
	roomNames := make(map[string]bool)    
	coordinates := make(map[string]bool)  

	i := 1
	for i < len(lines) {
		line := lines[i]

		if line == "" || (strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##")) {
			i++
			continue
		}

		if line == "##start" || line == "##end" {
			isStart := (line == "##start")
			if i+1 >= len(lines) {
				ant.Err = fmt.Errorf("invalid data format, missing %s room definition", map[bool]string{true: "start", false: "end"}[isStart])
				return ant
			}

			next := strings.Fields(lines[i+1])
			if len(next) != 3 {
				ant.Err = fmt.Errorf("invalid data format, invalid %s room", map[bool]string{true: "start", false: "end"}[isStart])
				return ant
			}
			name := next[0]
			xStr, yStr := next[1], next[2]
			X, errX := strconv.Atoi(xStr)
			Y, errY := strconv.Atoi(yStr)
			if errX != nil || errY != nil {
				ant.Err = fmt.Errorf("invalid data format, invalid coordinates")
				return ant
			}
			
			coordKey := fmt.Sprintf("%d,%d", X, Y)
			if strings.HasPrefix(name, "#") || strings.HasPrefix(name, "L") || roomNames[name] || coordinates[coordKey] {
				ant.Err = fmt.Errorf("invalid data format, duplicate or invalid room/%s", name)
				return ant
			}
			roomNames[name] = true
			coordinates[coordKey] = true
			ant.RoomsWithCords[name] = []string{strconv.Itoa(X), strconv.Itoa(Y)}
			if isStart {
				ant.Start = name
			} else {
				ant.End = name
			}
			i += 2
			continue
		}

		parts := strings.Fields(line)

		if len(parts) == 3 {
			name := parts[0]
			xStr, yStr := parts[1], parts[2]
			X, errX := strconv.Atoi(xStr)
			Y, errY := strconv.Atoi(yStr)
			if errX != nil || errY != nil {
				ant.Err = fmt.Errorf("invalid data format, invalid coordinates")
				return ant
			}
			coordKey := fmt.Sprintf("%d,%d", X, Y)
			if strings.HasPrefix(name, "#") || strings.HasPrefix(name, "L") || roomNames[name] || coordinates[coordKey] {
				ant.Err = fmt.Errorf("invalid data format, duplicate or invalid room/%s", name)
				return ant
			}
			roomNames[name] = true
			coordinates[coordKey] = true
			ant.RoomsWithCords[name] = []string{strconv.Itoa(X), strconv.Itoa(Y)}
			i++
			continue
		}

		if strings.Contains(parts[0], "-") && !strings.HasPrefix(parts[0], "#") && len(parts) == 1 {
			ant.Links = append(ant.Links, parts[0])
			i++
			continue
		}

		ant.Err = fmt.Errorf("invalid data format, unrecognized line: %q", line)
		return ant
	}

	if ant.Start == "" || ant.End == "" || ant.Start == ant.End {
		ant.Err = fmt.Errorf("invalid data format, missing start or end room")
		return ant
	}

	if !isBetween(filename) {
		ant.Err = fmt.Errorf("invalid data format, paths shouldn't be between ##start and ##end rooms")
		return ant
	}

	return ant
}

func isBetween(filename string) bool {
	fh, err := os.Open(filename)
	if err != nil {
		return false
	}
	defer fh.Close()

	inStartBlock := false
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "##start" {
			inStartBlock = true
			continue
		}
		if line == "##end" {
			inStartBlock = false
			continue
		}
		if inStartBlock && strings.Contains(line, "-") {
			return false
		}
	}
	return scanner.Err() == nil
}

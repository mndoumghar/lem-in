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

	file, err := os.ReadFile(filename)
	if err != nil {
		ant.Err = err
		return ant
	}
	fmt.Println(string(file) + "\n")
rawLines := strings.Split(string(file), "\n")
lines := make([]string, 0, len(rawLines))

for _, line := range rawLines {
    lines = append(lines, strings.TrimSpace(line))
}
	//ant.AntNum, err = strconv.Atoi(lines[0])
	ant.AntNum, err = strconv.Atoi(strings.TrimSpace(lines[0]))

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
		if lines[i] == "##start" {
			a := strings.Split(lines[i+1], " ")
			if len(a) != 3 {
				ant.Err = fmt.Errorf("invalid data format, invalid start room")
				return ant
			}
			ant.Start = strings.Split(lines[i+1], " ")[0]
			i++
		}
		if lines[i] == "##end" {
			a := strings.Split(lines[i+1], " ")
			if len(a) != 3 {
				ant.Err = fmt.Errorf("invalid data format, invalid start room")
				return ant
			}
			ant.End = strings.Split(lines[i+1], " ")[0]
			i++
		}
		if len(lines[i]) == 0 || strings.HasPrefix(lines[i], "#") {
			i++
			continue
		}
		parts := strings.Split(lines[i], " ")
		if len(parts) == 3 {
			room := parts[0]
			cords := parts[1:]
			X, err := strconv.Atoi(cords[0])
			if err != nil {
				ant.Err = fmt.Errorf("invalid data format, invalid coordinates X")
			}
			Y, err := strconv.Atoi(cords[1])
			if err != nil {
				ant.Err = fmt.Errorf("invalid data format, invalid coordinates Y")
				return ant
			}
			cords = []string{strconv.Itoa(X), strconv.Itoa(Y)}
			cordsKey := strings.Join(cords, ",")

			if strings.HasPrefix(room, "#") || strings.HasPrefix(room, "L") {
				ant.Err = fmt.Errorf("invalid data format, invalid room name")
				return ant
			}
			if roomNames[room] {
				ant.Err = fmt.Errorf("invalid data format, duplicate room")
				return ant
			}
			if coordinates[cordsKey] {
				ant.Err = fmt.Errorf("invalid data format, duplicate coordinates")
				return ant
			}
			roomNames[room] = true
			coordinates[cordsKey] = true
			ant.RoomsWithCords[room] = cords
		} else if len(parts) == 1 && strings.Contains(parts[0], "-") {
			ant.Links = append(ant.Links, parts[0])
		}

		i++
	}
	if ant.Start == "" || ant.End == "" || ant.Start == ant.End {
		ant.Err = fmt.Errorf("invalid data format, missing start or end room")
		return ant
	}

	if !isBetween(os.Args[1]) {
		ant.Err = fmt.Errorf("invalid data format, paths shouldn't be between ##start and ##end rooms")
		return ant
	}
	return ant
}

func isBetween(filename string) bool {
	file, err := os.Open(filename)
	if err != nil {
		return false
	}
	defer file.Close()

	inStartBlock := false
	inEndBlock := false
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "##start") {
			inStartBlock = true
		} else if strings.Contains(line, "##end") {
			inEndBlock = true
		}

		if inStartBlock && !inEndBlock {
			if strings.Contains(line, "-") {
				return false
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return false
	}
	return true
}

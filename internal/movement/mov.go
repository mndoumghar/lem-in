package movement

import (
	"fmt"
	"strings"
)

type Ant struct {
	id       int
	path     []string
	position int
}


type Path struct {
	id    int
	rooms []string
	ants  int
}


func SimulateAntMovement(availablePaths [][]string, antCount int) {
	allPaths := []Path{}
	for idx, rooms := range availablePaths {
		allPaths = append(allPaths, Path{id: idx + 1, rooms: rooms, ants: 0})
	}

	allAnts := AntAssignmentToPaths(allPaths, antCount)

	maxJourney := 0
	for _, path := range availablePaths {
		if len(path) > maxJourney {
			maxJourney = len(path)
		}
	}

	for currentStep := 0; currentStep < maxJourney+antCount; currentStep++ {
		active := false
		occupiedRooms := make(map[string]bool)
		inProgressPaths := make(map[string]bool)
		steps := []string{}

		for i := 0; i < len(allAnts); i++ {
			ant := &allAnts[i]

			if ant.position < len(ant.path)-1 {
				active = true
				ant.position++

				currentRoom := ant.path[ant.position]
				initialRoom := ant.path[0]

				if occupiedRooms[currentRoom] && inProgressPaths[initialRoom] {
					ant.position--
					continue
				}

				inProgressPaths[initialRoom] = true
				occupiedRooms[currentRoom] = true
				steps = append(steps, fmt.Sprintf("L%d-%s", ant.id, currentRoom))
			}
		}

		if len(steps) > 0 {
			fmt.Println(strings.Join(steps, " "))
		}

		
		if !active {
			break
		}
	}
}


func AntAssignmentToPaths(routePaths []Path, numAnts int) []Ant {
	antList := make([]Ant, numAnts)
	pathCount := len(routePaths)

	for i := 0; i < numAnts; i++ {
		minIdx := 0
		minWeight := len(routePaths[0].rooms) + routePaths[0].ants

		for j := 1; j < pathCount; j++ {
			weight := len(routePaths[j].rooms) + routePaths[j].ants
			if weight < minWeight {
				minIdx = j
				minWeight = weight
			}
		}

		routePaths[minIdx].ants++
		antList[i] = Ant{
			id:       i + 1,
			path:     routePaths[minIdx].rooms,
			position: -1,
		}
	}

	return antList
}

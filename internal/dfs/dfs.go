package dfs

import (
	"strings"
)

var (
	allPath      [][]string
	nowPath   []string
	visited = make(map[string]bool)
	room    = make(map[string][]string)
)

func Dfs(start, end string) {
	visited[start] = true
	nowPath = append(nowPath, start)

	if start == end {
		slice := []string{}
		slice = append(slice, nowPath...)
		allPath = append(allPath, slice)

	}
	for _, char := range room[start] {
		if !visited[char] {

			Dfs(char, end)
		}
	}

	visited[start] = false
	nowPath = nowPath[:len(nowPath)-1]
}

func SortPaths(paths [][]string) [][]string {
	for i := 0; i < len(paths); i++ {
		for j := i + 1; j < len(paths); j++ {
			if len(paths[i]) > len(paths[j]) {
				paths[i], paths[j] = paths[j], paths[i]
			}
		}
	}
	return paths
}

func FindPaths(start, end string, links []string) [][]string {
	buildRoomMap(links)
	Dfs(start, end)
	return allPath
}



func buildRoomMap(links []string) {
	for _, link := range links {
		rooms := strings.Split(link, "-")
		if len(rooms) != 2 {
			continue
		}
		
		room[rooms[0]] = append(room[rooms[0]], rooms[1])
		room[rooms[1]] = append(room[rooms[1]], rooms[0])
	}
}







func findMinScorePath(scores []int) (minIndex int) {
	minScore := -1
	for i, score := range scores {
		if score == -1 {
			continue
		}
		if minScore == -1 || score < minScore {
			minScore = score
			minIndex = i
		}
	}
	scores[minIndex] = -1
	return minIndex
}

func isValidPath(path []string, seen map[string]bool) bool {
	for _, room := range path {
		if seen[room] {
			return false
		}
		seen[room] = true
	}
	return true
}

func calculatePathScores(paths [][]string) []int {
	scores := make([]int, len(paths))
	for i, path := range paths {
		for _, room := range path {
			for _, otherPath := range paths {
				for _, otherRoom := range otherPath {
					if room == otherRoom {
						scores[i]++
					}
				}
			}
		}
	}
	return scores
}

func RemoveDuplicatePaths(paths [][]string) [][]string {
	scores := calculatePathScores(paths)
	result := make([][]string, 0, len(paths))
	seen := make(map[string]bool)

	for range paths {
		pathIndex := findMinScorePath(scores)
		innerPath := paths[pathIndex][1:]
		if isValidPath(innerPath[:len(innerPath)-1], seen) {
			result = append(result, innerPath)
		}
	}
	return result
}

package main

import (
	"fmt"

	"lemin/internal/dfs"
	"lemin/internal/movement"


	"lemin/internal/ants"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <filename>")
		return
	}
	nmla := ants.ReadFile(os.Args[1])

	if nmla.Err != nil {
		fmt.Printf("Error: %v\n", nmla.Err)
		return
	}

	paths := dfs.FindPaths(nmla.Start, nmla.End, nmla.Links)

	paths = dfs.RemoveDuplicatePaths(paths)

	if len(paths) == 0 {
		fmt.Println("ERROR: invalid data format, no path found")
		return
	}
	fmt.Println(paths)

	movement.SimulateAntMovement(paths, nmla.AntNum)

}

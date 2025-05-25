package main

import (
	"fmt"	
	"lemin/internal/parsfile"
)


// File



func main() {

	
	g, ants, err := parsfile.ParseFile("file1.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Ant", ants)
	for name, room := range g.Rooms {

		fmt.Printf("Room %s : (%d, %d) -> ", name, room.X, room.Y)
		for _, link := range room.Links {

			fmt.Printf("%s ", link.Name)
			
		}
		fmt.Println("")
		
	}




/*
	g := &Graph{}

	g.AddRoom("A", 1, 2, true, false) // start
	g.AddRoom("B", 3, 4, false, false)
	g.AddRoom("C", 5, 6, false, true) // End

	g.AddLink("A", "B")
	g.AddLink("B", "C")

	for name, room := range g.Rooms {

		fmt.Printf("Room %s is connected To : ", name)
		for _, link := range room.Links {
			fmt.Printf("%s ", link.Name)
		}
		fmt.Println("")
		
		}
		*/

}

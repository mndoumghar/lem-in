package main

import (
    "fmt"
    "log"
    "os"

    "lemin/internal/parsfile"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run . <input_file>")
        return
    }

    path := os.Args[1]

    // تحليل الملف
    graph, ants, err := parsfile.ParseFile(path)
    if err != nil {
        log.Fatalf("Error parsing file: %v\n", err)
    }

    fmt.Printf("Ants: %d\n", ants)

    // طباعة الغرف والروابط (اختياري للتأكد)
    for name, room := range graph.Rooms {
        fmt.Printf("Room %s : (%d, %d) -> ", name, room.X, room.Y)
        for _, link := range room.Links {
            fmt.Printf("%s ", link.Name)
        }
        fmt.Println()
    }

    // استعمال DFS لإيجاد جميع الطرق

    paths := parsfile.FindAllPathsDFS(graph, graph.Start.Name, graph.End.Name)
	
    fmt.Println("All paths from start to end:")
    for _, path := range paths {
        fmt.Println(path)
    }
}

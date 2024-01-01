# lem-in

A high-performance Go solution for the ant colony pathfinding problem. This program finds optimal paths for ants to traverse through a colony from source to destination while minimizing total movement steps.

##  Description

Lem-in solves the classic algorithmic challenge of finding multiple non-overlapping paths in a graph to efficiently move a colony of ants from a start room to an end room. It uses advanced graph algorithms to optimize ant movement and minimize the total number of steps required.

**Key Concept:** Given a number of ants and a map of interconnected rooms, find the fastest way to move all ants from the start room to the end room without any two ants being in the same room or tunnel at the same time.

##  Features

- **Graph Parsing**: Efficiently parses colony maps with rooms and tunnel connections
- **Pathfinding**: Implements DFS-based pathfinding algorithms for optimal routing
- **Ant Simulation**: Simulates realistic ant movement through multiple paths simultaneously
- **Performance Optimized**: Minimizes total moves required for all ants
- **Error Handling**: Robust validation of input files and graph structure
- **Modular Design**: Separated concerns for ants, pathfinding (DFS), and movement logic

## 🔧 Requirements

- **Go** 1.16 or higher
- **OS**: Windows, Linux, or macOS

## 📥 Installation

1. Clone or navigate to the project directory:

```bash
cd c:\Users\NDOUMGHAR\readme\lem-in
```

2. Ensure Go is installed:

```bash
go version
```

3. Initialize Go modules (if not already done):

```bash
go mod init lem-in
```

## 🚀 Quick Start

Run the program with a map file:

```bash
    cd lem-in
    go run .\cmd\main  .\data.txt
```

Or build and execute:

```bash
go build -o lem-in ./cmd/main
./lem-in < map.txt
```

##  Input Format

The input file should contain:

```
<number_of_ants>
<room_name> <x_coordinate> <y_coordinate>
##start
<starting_room_name> <x> <y>
##end
<ending_room_name> <x> <y>
<room_name> <x> <y>
<room1>-<room2>
<room3>-<room4>
...
``

### Example Input:

```
10
room1 0 0
room2 1 1
##start
anthill 2 2
##end
exit 3 3
room1-room2
room2-anthill
anthill-exit
```

## Output

The program outputs ant movements in turns:

```
L1-room1 L2-room2 L3-room1
L1-room2 L2-anthill L3-room2
L1-anthill L2-exit L3-anthill
L1-exit L3-room2
L3-anthill
L3-exit
```

##  Project Structure

```
lem-in/
├── cmd/
│   └── main/
│       └── main.go          # Entry point
├── internal/
│   ├── ants/                # Ant management and logic
│   ├── dfs/                 # Depth-First Search pathfinding
│   └── movement/            # Ant movement simulation
├── README.md
├── go.mod
└── go.sum
```

### Module Descriptions

- **`cmd/main/`**: Entry point of the application. Handles input parsing and orchestrates the algorithm.
- **`internal/ants/`**: Manages ant data structures, properties, and ant-specific operations.
- **`internal/dfs/`**: Implements Depth-First Search algorithm for pathfinding through the colony graph.
- **`internal/movement/`**: Simulates and coordinates ant movement along discovered paths.

##  Algorithm Details

- Uses **Depth-First Search (DFS)** for pathfinding through the colony
- Implements **Multi-path routing** for optimal ant distribution
- Handles **concurrent movement** of multiple ants simultaneously
- Optimizes for **minimum total steps** across all ants

##  Contributing


- Our development process
- How to propose bugfixes and improvements
- How to build and test your changes
- Code of Conduct

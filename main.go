package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cell struct {
	x, y int64
}

type State map[Cell]bool

func parseInt64(s string) (int64, error) {
	return strconv.ParseInt(strings.TrimSpace(s), 10, 64)
}

func readInput() (State, error) {
	state := make(State)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}
		x, err := parseInt64(parts[0])
		if err != nil {
			return nil, err
		}
		y, err := parseInt64(parts[1])
		if err != nil {
			return nil, err
		}
		state[Cell{x, y}] = true
	}
	return state, scanner.Err()
}

func getNeighbors(c Cell) []Cell {
	return []Cell{
		{x: c.x - 1, y: c.y - 1},
		{x: c.x - 1, y: c.y},
		{x: c.x - 1, y: c.y + 1},
		{x: c.x, y: c.y - 1},
		{x: c.x, y: c.y + 1},
		{x: c.x + 1, y: c.y - 1},
		{x: c.x + 1, y: c.y},
		{x: c.x + 1, y: c.y + 1},
	}
}

func countLiveNeighbors(state State, c Cell) int {
	count := 0
	for _, neighbor := range getNeighbors(c) {
		if state[neighbor] {
			count++
		}
	}
	return count
}

func nextGeneration(current State) State {
	next := make(State)

	candidates := make(map[Cell]bool)
	for cell := range current {
		candidates[cell] = true
		for _, neighbor := range getNeighbors(cell) {
			candidates[neighbor] = true
		}
	}

	for cell := range candidates {
		count := countLiveNeighbors(current, cell)
		alive := current[cell]
		if alive && (count == 2 || count == 3) {
			next[cell] = true
		} else if !alive && count == 3 {
			next[cell] = true
		}
	}
	return next
}

func printState(state State) {
	fmt.Println("#Life 1.06")
	for cell := range state {
		fmt.Printf("%d %d\n", cell.x, cell.y)
	}
}

func main() {
	state, err := readInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading inpu: %v\n", err)
		os.Exit(1)
	}

	for i := 0; i < 10; i++ {
		state = nextGeneration(state)
	}
	printState(state)
}

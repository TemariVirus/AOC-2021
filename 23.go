package main

import (
	"math"
	"strings"
)

type Amphipod struct {
	Type byte
	Pos  Point
}

var AMPHIPOD_COSTS = [5]float64{0, 1, 10, 100, 1000}

func (a Amphipod) index() int {
	index := a.Pos.X
	if a.Pos.Y > 0 {
		index = a.Pos.X/2 + a.Pos.Y*4 + 6
	}
	return index
}

func (a Amphipod) isSolved(rooms [4][]byte) bool {
	if a.Type*2 != byte(a.Pos.X) {
		return false
	}

	room := rooms[a.Type-1]
	for y := a.Pos.Y - 1; y < len(room); y++ {
		if room[y] != a.Type {
			return false
		}
	}
	return true
}

func parseInput23(input string) []Amphipod {
	lines := strings.Split(input, "\n")
	amphipods := []Amphipod{}
	for y, line := range lines[2:4] {
		for x := 3; x <= 9; x += 2 {
			amphipods = append(amphipods, Amphipod{Type: line[x] - 'A' + 1, Pos: Point{X: x - 1, Y: y + 1}})
		}
	}
	return amphipods
}

func makeAmphipodRooms(amphipods []Amphipod) [4][]byte {
	room_size := len(amphipods) / 4
	rooms := [4][]byte{
		make([]byte, room_size),
		make([]byte, room_size),
		make([]byte, room_size),
		make([]byte, room_size),
	}

	for _, amphipod := range amphipods {
		if amphipod.Pos.Y == 0 {
			continue
		}
		rooms[(amphipod.Pos.X/2)-1][amphipod.Pos.Y-1] = amphipod.Type
	}
	return rooms
}

var hashAmphimods = func() func(amphipods []Amphipod) uint64 {
	powers_of_5 := []uint64{1}

	return func(amphipods []Amphipod) uint64 {
		hash := uint64(0)
		for _, amphipod := range amphipods {
			index := amphipod.index()
			for len(powers_of_5) <= index {
				powers_of_5 = append(powers_of_5, powers_of_5[len(powers_of_5)-1]*5)
			}
			hash += powers_of_5[index] * uint64(amphipod.Type)
		}
		return hash
	}
}()

func solution23Part1(input string) float64 {
	amphipods := parseInput23(input)
	rooms := makeAmphipodRooms(amphipods)

	return findLeastEnergy(amphipods, [11]bool{}, rooms, map[uint64]float64{})
}

func findLeastEnergy(amphipods []Amphipod, hallway [11]bool, rooms [4][]byte, seen map[uint64]float64) float64 {
	const MIN_X, MAX_X = 0, 10

	hash := hashAmphimods(amphipods)
	if energy, ok := seen[hash]; ok {
		return energy
	}

	min_energy := math.Inf(1)
	seen[hash] = math.Inf(1)
	mobile_amphs := []int{}
	solvedCount := 0
	for i, amphipod := range amphipods {
		// No need to move
		if amphipod.isSolved(rooms) {
			solvedCount++
			continue
		}

		// In room
		if amphipod.Pos.Y != 0 {
			room := rooms[(amphipod.Pos.X/2)-1]
			obstructed := false
			for y := 0; y < amphipod.Pos.Y-1; y++ {
				if room[y] != 0 {
					obstructed = true
					break
				}
			}

			if !obstructed {
				mobile_amphs = append(mobile_amphs, i)
			}

			continue
		}

		// From hallway to room
		room := rooms[amphipod.Type-1]
		target_x := int(amphipod.Type * 2)
		target_y := len(room)
		left, right := min(amphipod.Pos.X, target_x), max(amphipod.Pos.X, target_x)

		obstructed := false
		for x := left + 1; !obstructed && x < right; x++ {
			if hallway[x] {
				obstructed = true
			}
		}
		for y := target_y - 1; !obstructed && y >= 0; y-- {
			if room[y] == 0 {
				target_y = y + 1
				break
			}
			if room[y] != amphipod.Type {
				obstructed = true
			}
		}
		for y := 0; !obstructed && y < target_y-1; y++ {
			if room[y] != 0 {
				obstructed = true
			}
		}
		if obstructed {
			continue
		}

		amphipods[i].Pos = Point{target_x, target_y}
		hallway[amphipod.Pos.X] = false
		room[target_y-1] = amphipod.Type
		energy := findLeastEnergy(amphipods, hallway, rooms, seen)
		room[target_y-1] = 0
		// hallway[amphipod.Pos.X] = true // Not needed
		amphipods[i].Pos = amphipod.Pos

		dist := absInt(target_x-amphipod.Pos.X) + target_y
		energy += AMPHIPOD_COSTS[amphipod.Type] * float64(dist)
		seen[hash] = energy
		return energy
	}

	if solvedCount == len(amphipods) {
		seen[hash] = 0
		return 0
	}

	// From room to hallway
	for _, i := range mobile_amphs {
		amphipod := amphipods[i]
		cell := &rooms[(amphipod.Pos.X/2)-1][amphipod.Pos.Y-1]
		for x := amphipod.Pos.X - 1; x >= MIN_X; x-- {
			if x == 2 || x == 4 || x == 6 || x == 8 {
				continue
			}
			if hallway[x] {
				break
			}

			amphipods[i].Pos = Point{x, 0}
			*cell = 0
			hallway[x] = true
			energy := findLeastEnergy(amphipods, hallway, rooms, seen)
			hallway[x] = false
			*cell = amphipod.Type
			amphipods[i].Pos = amphipod.Pos

			dist := amphipod.Pos.X - x + amphipod.Pos.Y
			energy += AMPHIPOD_COSTS[amphipod.Type] * float64(dist)
			min_energy = math.Min(min_energy, energy)
		}
		for x := amphipod.Pos.X + 1; x <= MAX_X; x++ {
			if x == 2 || x == 4 || x == 6 || x == 8 {
				continue
			}
			if hallway[x] {
				break
			}

			amphipods[i].Pos = Point{x, 0}
			*cell = 0
			hallway[x] = true
			energy := findLeastEnergy(amphipods, hallway, rooms, seen)
			hallway[x] = false
			*cell = amphipod.Type
			amphipods[i].Pos = amphipod.Pos

			dist := x - amphipod.Pos.X + amphipod.Pos.Y
			energy += AMPHIPOD_COSTS[amphipod.Type] * float64(dist)
			min_energy = math.Min(min_energy, energy)
		}
	}

	seen[hash] = min_energy
	return min_energy
}

func solution23Part2(input string) float64 {
	amphipods := parseInput23(input)
	for i := 4; i < 8; i++ {
		amphipods[i].Pos.Y = 4
	}
	amphipods = append(amphipods, Amphipod{Type: 4, Pos: Point{X: 2, Y: 2}})
	amphipods = append(amphipods, Amphipod{Type: 3, Pos: Point{X: 4, Y: 2}})
	amphipods = append(amphipods, Amphipod{Type: 2, Pos: Point{X: 6, Y: 2}})
	amphipods = append(amphipods, Amphipod{Type: 1, Pos: Point{X: 8, Y: 2}})
	amphipods = append(amphipods, Amphipod{Type: 4, Pos: Point{X: 2, Y: 3}})
	amphipods = append(amphipods, Amphipod{Type: 2, Pos: Point{X: 4, Y: 3}})
	amphipods = append(amphipods, Amphipod{Type: 1, Pos: Point{X: 6, Y: 3}})
	amphipods = append(amphipods, Amphipod{Type: 3, Pos: Point{X: 8, Y: 3}})

	rooms := makeAmphipodRooms(amphipods)

	return findLeastEnergy(amphipods, [11]bool{}, rooms, map[uint64]float64{})
}

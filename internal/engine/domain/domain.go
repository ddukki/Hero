package domain

import "math/rand"

// Defines a state that the engine can be in. The state can be running in the
// background. Every state updates itself based on the engine clock.
type State interface {
	// Update refreshes the state as the engine clock ticks forward.
	Update()
}

// Clock provides regular tick intervals for the engine to update its states.
type Clock interface {
	// Ticks the clock forward and updates any variables within the clock.
	Tick()
}

// RollD generates a random integer between 1 and the number supplied. The
// distribution is uniform.
func RollD(d int) int {
	return rand.Intn(d) + 1
}

// RollND generates a random integer between 1 and 'd' with a uniform
// distribution 'n' times and returns the sum.
func RollND(n, d int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += RollD(d)
	}

	return sum
}

// RollBaseStat generates a random integer between 1 and 6 with a uniform
// distribution 4 times and returns the sum of the highest 3 values.
func RollBaseStat() int {
	sum := 0
	lowest := 7
	for i := 0; i < 4; i++ {
		r := RollD(6)
		sum += r
		if lowest > r {
			lowest = r
		}
	}
	return sum - lowest
}

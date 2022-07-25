package domain

import (
	"fmt"

	"theduckymonk.org/hero/internal/engine/domain"
)

type Number interface {
	uint | uint32 | uint64 | int | int32 | int64 | float32 | float64
}

// Stats contains the values needed to determine statistical attributes of an
// entity.
type Stats struct {
	// physical stats
	Strength     int
	Constitution int
	Dexterity    int

	// soft stats
	Charisma     int
	Wisdom       int
	Intelligence int
}

func NewRandomStats() *Stats {
	return &Stats{
		Strength:     domain.RollBaseStat(),
		Constitution: domain.RollBaseStat(),
		Dexterity:    domain.RollBaseStat(),
		Charisma:     domain.RollBaseStat(),
		Wisdom:       domain.RollBaseStat(),
		Intelligence: domain.RollBaseStat(),
	}
}

// String returns a compact string version of the stats block.
func (s *Stats) String() string {
	return fmt.Sprintf("Str: %d, Con: %d, Dex: %d, Cha: %d, Wis: %d, Int: %d",
		s.Strength, s.Constitution, s.Dexterity, s.Charisma, s.Wisdom, s.Intelligence)
}

// Entity is the atomic building block of anything that has a stat block.
type Entity interface {
	// GetStats returns the raw stats as-is from the entity.
	Stats() *Stats

	// UpdateStats updates the currently stored stats based on the characteristics
	// and specific attributes of the entity.
	UpdateStats()
}

// DerivedStats provide calculation and generation of stat values derived from
// the raw characteristics of an entity.
type DerivedStats interface {
	// GenerateStat calculates a new Stat value based on the given stats. For
	// example, a black-smithing stat might be generated from strength, dexterity,
	// and intelligenece.
	GenerateStat(Stats) int
}

// RangedStat defines an interface for a stat with a value that is variable
// within a range.
type RangedStat[T Number] interface {
	// GetMax returns the max value of the stat.
	Max() T

	// GetMin returns the minimum value of the stat.
	Min() T

	// GetValue returns the current value of the stat.
	Value() T
}

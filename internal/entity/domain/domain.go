package domain

import (
	"fmt"

	"github.com/ddukki/Hero/internal/engine/domain"
)

// StatEnum defines an enumeration for the different types of stat values.
type StatEnum int

const (
	STR StatEnum = iota
	CON
	DEX
	CHA
	INT
	WIS

	_statlim
)

// A range of different numeric types.
type Number interface {
	uint | uint32 | uint64 | int | int32 | int64 | float32 | float64
}

// Stats contains the values needed to determine statistical attributes of an
// entity.
type Stats struct {
	Conds map[ConditionEnum]Condition
	Base  map[StatEnum]int
}

// GetStatEnums returns  an exhaustive list of the stat enums.
func GetStatEnums() []StatEnum {
	se := make([]StatEnum, _statlim)
	for e := StatEnum(0); e < _statlim; e++ {
		se[e] = e
	}
	return se
}

// NewRandomBaseStats generates new random stats based on domain.RollBaseStat.
func NewRandomBaseStats() *Stats {
	s := &Stats{Base: make(map[StatEnum]int)}

	for e := StatEnum(0); e < _statlim; e++ {
		s.Base[e] = domain.RollBaseStat()
	}

	return s
}

// String returns a compact string version of the stats block.
func (s *Stats) String() string {
	b := s.Base
	return fmt.Sprintf("Str: %d, Con: %d, Dex: %d, Cha: %d, Wis: %d, Int: %d",
		b[STR], b[CON], b[DEX], b[CHA], b[WIS], b[INT])
}

// ConditionEnum defines an enumeration for the different types of condition
// ranged stats.
type ConditionEnum int

const (
	HP ConditionEnum = iota
	MP
	SP

	_condlim
)

// GetConditionEnums returns a full list of all condition enums.
func GetConditionEnums() []ConditionEnum {
	ce := make([]ConditionEnum, _condlim)
	for e := ConditionEnum(0); e < _condlim; e++ {
		ce[e] = e
	}

	return ce
}

var _ RangedStat[int] = (*Condition)(nil)

// Condition defines the condition of whatever contains it. The condition can
// deteriorate, indicated by the current value of the condition.
type Condition struct {
	max int
	min int
	val int
}

// Max implements RangedStat.
func (c *Condition) Max() int {
	return c.max
}

// Min implements RangedStat.
func (c *Condition) Min() int {
	return c.min
}

// Value implements RangedStat.
func (c *Condition) Value() int {
	return c.val
}

// Set implements RangedStat.
func (c *Condition) Set(v int) {
	if v < c.min {
		v = c.min
	} else if v > c.max {
		v = c.max
	}

	c.val = v
}

// Set implements RangedStat.
func (c *Condition) SetMin(min int) {
	if min >= c.max {
		min = c.max - 1
	}
	c.min = min
}

// Set implements RangedStat.
func (c *Condition) SetMax(max int) {
	if max <= c.min {
		max = c.min + 1
	}
	c.max = max
}

// SetMinMax implements RangedStat.
func (c *Condition) SetMinMax(min, max int) {
	if max <= min {
		max = min + 1
	}
	c.min = min
	c.max = max
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

	// Value returns the current value of the stat. If the value is out of bounds,
	// the value will be clipped to the min or max.
	Value() T

	// Set sets the value of the ranged stat.
	Set(T)

	// SetMax sets the value of the maximum of the ranged stat.
	SetMax(T)

	// SetMin sets the value of the minimum of the ranged stat.
	SetMin(T)

	// SetMinMax sets the values of the minimum and maximum of the ranged stat.
	SetMinMax(T, T)
}

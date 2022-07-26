package character

import (
	"fmt"
	"math/rand"

	"github.com/ddukki/Hero/internal/entity/domain"
	"github.com/ddukki/Hero/internal/lang"
)

// Character defines the properties and characteristics of a particular
// character.
type Character struct {
	conditions       map[domain.ConditionEnum]*domain.Condition
	GivenName        string
	GenerationalName string
	KnownLangs       []*lang.Language
	baseStats        *domain.Stats
	education        map[string]Learned
}

// NewCharacter creates a new character with the given attributes.
func NewCharacter(
	stats *domain.Stats,
	langs []*lang.Language,
	given string,
	generational string,
) *Character {
	return &Character{
		GivenName:        given,
		GenerationalName: generational,
		KnownLangs:       langs,
		baseStats:        stats,
		education:        make(map[string]Learned),
	}
}

// NewRandomCharacter creates a new character with randomized attributes.
func NewRandomCharacter() *Character {

	// Get a random language.
	nLang := len(lang.DB.List)
	rl := rand.Intn(nLang)
	l := lang.DB.List[rl]
	languages := []*lang.Language{l}

	given := l.RulesByName[lang.RULE_GIVEN].Generate()
	generational := l.RulesByName[lang.RULE_GENERATIONAL].Generate()

	return NewCharacter(domain.NewRandomStats(), languages, given, generational)
}

// UpdateConditions updates the conditions based on the stats of the character.
func (c *Character) UpdateConditions() {
	panic("not implemented")
}

// String generates a string that gives the minimal information that defines
// this character.
func (c *Character) String() string {
	return fmt.Sprintf("Name: %s %s, baseStats: %s",
		c.GivenName, c.GenerationalName, c.baseStats.String())
}

// GetStat returns the desired stat, including any modifiers applied.
func (c *Character) GetStat(se domain.StatEnum) int {
	stat, ok := c.baseStats.Base[se]
	if !ok {
		return 0
	}
	for _, s := range c.education {
		stat += s.GetModifier(se)
	}
	return stat
}

// GetCondition returns the desired condition.
func (c *Character) GetCondition(ce domain.ConditionEnum) *domain.Condition {
	cond, ok := c.conditions[ce]
	if !ok {
		cond = &domain.Condition{}
		cond.SetMax(1)
		cond.SetMin(0)
		cond.Set(1)
		return cond
	}

	return cond
}

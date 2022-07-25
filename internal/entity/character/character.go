package character

import (
	"fmt"

	"theduckymonk.org/hero/internal/entity/domain"
	"theduckymonk.org/hero/internal/lang"
)

type Character struct {
	GivenName        string
	GenerationalName string
	KnownLangs       []*lang.Language
	baseStats        *domain.Stats
}

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
	}
}

func NewRandomCharacter() *Character {
	l := lang.LangDB.Languages[0]
	languages := []*lang.Language{l}

	given := l.RulesByName[lang.RULE_GIVEN].Generate()
	generational := l.RulesByName[lang.RULE_GENERATIONAL].Generate()

	return NewCharacter(domain.NewRandomStats(), languages, given, generational)
}

func (c *Character) String() string {
	return fmt.Sprintf("Name: %s %s, baseStats: %s",
		c.GivenName, c.GenerationalName, c.baseStats.String())
}

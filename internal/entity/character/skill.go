package character

import "theduckymonk.org/hero/internal/entity/domain"

type SkillDB struct {
	SkillNames []string
	Skills     map[string]*Skill
}

type Skill struct {
	name string
	mods domain.Stats
}

// GetName retrieves the name of the skill proficiency.
func (s *Skill) GetName() string {
	return s.name
}

// GetModifiers retrieves the modifiers of the base stats that the proficiency
// provides.
func (s *Skill) GetModifiers() domain.Stats {
	return s.mods
}

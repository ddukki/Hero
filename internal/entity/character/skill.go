package character

import "github.com/ddukki/Hero/internal/entity/domain"

var DB *db

func init() {
	DB = &db{}
}

// db is a database containing all skills.
type db struct {
	SkillNames []string
	Skills     map[string]*Skill
}

// Skill defines a learnable skill.
type Skill struct {
	name string
	mods map[domain.StatEnum]int
}

// Learned defines a level for a learned skill and provides access to
// modifier values.
type Learned struct {
	base  *Skill
	level int
}

// GetName retrieves the name of the skill proficiency.
func (s *Skill) GetName() string {
	return s.name
}

// GetModifiers retrieves the modifiers of the base stats that the proficiency
// provides.
func (s *Skill) GetModifiers() map[domain.StatEnum]int {
	return s.mods
}

// GetLevel returns the level of the learned skill.
func (l *Learned) GetLevel() int {
	return l.level
}

// LevelUp increments the level counter.
func (l *Learned) LevelUp() {
	l.level++
}

// GetModifier returns a modified value for the stat with the level applied. If
// there is no stat value for the enum, a zero is returned.
func (l *Learned) GetModifier(se domain.StatEnum) int {
	mod, ok := l.base.mods[se]
	if !ok {
		return 0
	}
	return mod * l.level
}

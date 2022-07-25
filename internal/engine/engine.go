package engine

import "theduckymonk.org/hero/internal/engine/domain"

// Engine is the main driving force of the program.
type Engine struct {
	States []*domain.State
	Clock  domain.Clock
}

func NewEngine() *Engine {
	return &Engine{
		States: make([]*domain.State, 0),
	}
}

var _ domain.Clock = (*clock)(nil)

type clock struct {
	FPS int
}

func NewClock() domain.Clock {
	return &clock{
		FPS: 60,
	}
}

func (c *clock) Tick() {

}

// Subscribe adds the state as a subscriber to the engine. When the engine's
// clock ticks, all states will update.
func (e *Engine) Subscribe(s *domain.State) {
	e.States = append(e.States, s)
}

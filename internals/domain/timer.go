package domain

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Timer struct {
	currentTicks int
	targetTicks  int
}

func NewTimer(d time.Duration) *Timer {
	return &Timer{
		currentTicks: 0,
		targetTicks:  int(d.Milliseconds()) * ebiten.TPS() / 1000,
	}
}

func (t *Timer) Update() {
	if t.currentTicks < t.targetTicks {
		t.currentTicks++
	}
}

func (t *Timer) DecrementTargetTicks(gameSpeed int) {
	if t.TargetTicks() > 30 {
		t.targetTicks = t.targetTicks - gameSpeed
	} else {
		t.targetTicks = 30
	}
}

func (t *Timer) TargetTicks() int {
	return t.targetTicks
}

func (t *Timer) IsReady() bool {
	return t.currentTicks >= t.targetTicks
}

func (t *Timer) Reset() {
	t.currentTicks = 0
}

package main

import (
	_ "image/png"
	"slices"
	"time"

	"github.com/andersbloch/game/internals/domain"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player        domain.Player
	meteors       []*domain.Meteor
	bullets       []*domain.Bullet
	attackTimer   *domain.Timer
	shootCooldown *domain.Timer
}

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

func (g *Game) Update() error {
	if err := g.player.Update(); err != nil {
		return err
	}
	g.attackTimer.Update()
	if g.attackTimer.IsReady() {
		g.attackTimer.Reset()

		m := domain.NewMeteor(ScreenWidth, ScreenHeight)
		g.meteors = append(g.meteors, m)
	}
	g.shootCooldown.Update()
	if g.shootCooldown.IsReady() && ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.shootCooldown.Reset()
		g.bullets = append(g.bullets, domain.NewBullet(g.player.ShipCenter().X, g.player.ShipCenter().Y, g.player.Rotation()))
	}
	for i, m := range g.meteors {
		for j, b := range g.bullets {
			if m.IsColliding(b) {
				g.meteors = slices.Delete(g.meteors, i, i+1)
				g.bullets = slices.Delete(g.bullets, j, j+1)
				m.BlowUp()
			}
		}
	}

	for _, m := range g.meteors {
		m.Update()
	}
	for _, b := range g.bullets {
		b.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, m := range g.meteors {
		m.Draw(screen)
	}
	for _, b := range g.bullets {
		b.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	g := &Game{
		bullets:       []*domain.Bullet{},
		player:        *domain.NewPlayer(ScreenWidth, ScreenHeight),
		attackTimer:   domain.NewTimer(5 * time.Second),
		shootCooldown: domain.NewTimer(500000000), // 0.5 second
		meteors:       []*domain.Meteor{domain.NewMeteor(ScreenWidth, ScreenHeight)},
	}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}

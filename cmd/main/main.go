package main

import (
	"bytes"
	"fmt"
	_ "image/png"
	"time"

	"github.com/andersbloch/game/internals/domain"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	player        domain.Player
	meteors       []*domain.Meteor
	bullets       []*domain.Bullet
	attackTimer   *domain.Timer
	shootCooldown *domain.Timer
	score int
	fontSource *text.GoTextFaceSource
	fontFace *text.GoTextFace
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
		m := domain.NewMeteor(ScreenWidth, ScreenHeight, g.score + 1)
		g.meteors = append(g.meteors, m)
	}
	g.shootCooldown.Update()
	if g.shootCooldown.IsReady() && ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.shootCooldown.Reset()
		g.bullets = append(g.bullets, domain.NewBullet(g.player.ShipCenter().X, g.player.ShipCenter().Y, g.player.Rotation()))
	}

	meteorToRemove := 1000
	
	for i, m := range g.meteors {
		if m.IsColliding(g.player.Bounds()) {
			g.player.BlowUp()
		}
		for _, b := range g.bullets {
			if m.IsColliding(b.Bounds()) {
				meteorToRemove = i
				g.score = g.score + 1
				g.attackTimer.DecrementTargetTicks(g.score)
 				m.BlowUp()
			}
		}
	}

	if meteorToRemove < 1000 {
		g.meteors = append(g.meteors[:meteorToRemove], g.meteors[meteorToRemove+1:]...)
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
	
	for _, m := range g.meteors {
		m.Draw(screen)
	}
	for _, b := range g.bullets {
		b.Draw(screen)
	}
	g.player.Draw(screen)
	txtOp := &text.DrawOptions{}
    // Start drawing at the top center of the screen.
    txtOp.GeoM.Translate(5, 5)
    // By default, the text is white. We can call ScaleWithColor to specify a different color.
    //colorGreen := color.RGBA{0, 255, 0, 255}
    //txtOp.ColorScale.ScaleWithColor(colorGreen)
	text.Draw(screen, fmt.Sprintf("Score: %d, ticks: %d", g.score, g.attackTimer.TargetTicks()), g.fontFace, txtOp)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	fontSource, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		panic(err)
	}

	g := &Game{
		bullets:       []*domain.Bullet{},
		player:        *domain.NewPlayer(ScreenWidth, ScreenHeight),
		attackTimer:   domain.NewTimer(5 * time.Second),
		shootCooldown: domain.NewTimer(100000000), // 0.5 second
		meteors:       []*domain.Meteor{domain.NewMeteor(ScreenWidth, ScreenHeight, 1)},
		fontSource: fontSource,
		fontFace: &text.GoTextFace{
			Source: fontSource,
			Size:   16,
		},
	}

	err = ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}

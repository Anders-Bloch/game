package domain

import (
	"image"
	"math"

	projectRoot "github.com/andersbloch/game"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vector struct {
	X float64
	Y float64
}

func (v Vector) Normalize() Vector {
	length := math.Sqrt(v.X*v.X + v.Y*v.Y)
	if length == 0 {
		return Vector{X: 0, Y: 0}
	}
	return Vector{
		X: v.X / length,
		Y: v.Y / length,
	}
}

type Player struct {
	position Vector
	sprite   *ebiten.Image
	rotation float64
}

func NewPlayer(ScreenWidth, ScreenHeight float64) *Player {
	sprite := playerSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos := Vector{
		X: ScreenWidth/2 - halfW,
		Y: ScreenHeight/2 - halfH,
	}

	return &Player{
		position: pos,
		sprite:   sprite,
		rotation: 0,
	}
}

var playerSprite = projectRoot.MustLoadImage("assets/ship_A.png")

func (p *Player) Rotation() float64 {
	return p.rotation
}

func (p *Player) Bounds() image.Rectangle {
	return p.sprite.Bounds()
}

func (p *Player) ShipCenter() Vector {
	bounds := p.sprite.Bounds()
	shipCenterX := p.position.X + float64(bounds.Dx())/2
	shipCenterY := p.position.Y + float64(bounds.Dy())/2
	return Vector{X: shipCenterX, Y: shipCenterY}
}

func (p *Player) Update() error {
	speed := math.Pi / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += speed
	}
	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, op)
}

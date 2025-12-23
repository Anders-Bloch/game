package domain

import (
	"image"
	"math"

	projectRoot "github.com/andersbloch/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	position Vector
	movement Vector
	rotation float64
	sprite   *ebiten.Image
}

var BulletSprite = projectRoot.MustLoadImage("assets/effect_purple.png")

func NewBullet(shipX, shipY float64, rotation float64) *Bullet {
	sprite := BulletSprite

	// Randomized velocity
	velocity := 1.0

	// Direction is based on rotation
	rotationRad := rotation - math.Pi/2 // Adjusting for sprite orientation
	direction := Vector{
		X: math.Cos(rotationRad),
		Y: math.Sin(rotationRad),
	}

	// Normalize the vector — get just the direction without the length
	normalizedDirection := direction.Normalize()

	// Multiply the direction by velocity
	movement := Vector{
		X: normalizedDirection.X * velocity,
		Y: normalizedDirection.Y * velocity,
	}

	return &Bullet{
		position: Vector{X: shipX, Y: shipY},
		movement: movement,
		sprite:   sprite,
		rotation: rotation,
	}
}

func (b *Bullet) Bounds() image.Rectangle {
	return b.sprite.Bounds()
}

func (b *Bullet) Update() error {
	b.position.X += b.movement.X
	b.position.Y += b.movement.Y
	return nil
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.1, 0.1)
	op.GeoM.Rotate(b.rotation)
	op.GeoM.Translate(b.position.X, b.position.Y)
	screen.DrawImage(b.sprite, op)
}

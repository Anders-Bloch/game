package domain

import (
	"image/color"
	"math"

	projectRoot "github.com/andersbloch/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Bullet struct {
	position Vector
	movement Vector
	rotation float64
	sprite   *ebiten.Image
	boundary Circle
}

var BulletSprite = projectRoot.MustLoadImage("assets/effect_purple.png")

func NewBullet(shipX, shipY float64, rotation float64) *Bullet {
	sprite := BulletSprite

	// Randomized velocity
	velocity := 2.5

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
		boundary: Circle{
			X: shipX,
			Y: shipY+2, 
			Radius: 4.0,
		},
	}
}

func (b *Bullet) Bounds() Circle {
	return b.boundary
}

func (b *Bullet) Update() error {
	b.position.X += b.movement.X
	b.position.Y += b.movement.Y
	b.boundary.X = b.position.X
	b.boundary.Y = b.position.Y + 2
	return nil
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.1, 0.1)
	op.GeoM.Rotate(b.rotation)
	op.GeoM.Translate(b.position.X, b.position.Y)

    strokeWidth := 2.0
    clr := color.RGBA{0, 255, 0, 255} // Green
    
    vector.StrokeCircle(screen, float32(b.boundary.X), float32(b.boundary.Y), float32(b.boundary.Radius), float32(strokeWidth), clr, true)
}

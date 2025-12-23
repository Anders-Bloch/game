package domain

import (
	"fmt"
	"image/color"
	"math"
	"math/rand/v2"

	projectRoot "github.com/andersbloch/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var MeteorSprite = projectRoot.MustLoadImage("assets/meteor_small.png")

type Meteor struct {
	position      Vector
	movement      Vector
	sprite        *ebiten.Image
	rotationSpeed float64
	rotation      float64
	boundary 	  Circle
}

func (m *Meteor) BlowUp() {
	fmt.Println("Blow up!")
	// Placeholder for blow-up logic
}

func (m *Meteor) IsColliding(c Circle) bool {
	result := m.boundary.Intersect(c)
	if result {
		fmt.Printf("IsColliding: %v \n", result)

	}
	return result
}

func NewMeteor(ScreenWidth, ScreenHeight float64, gameSpeed int) *Meteor {
	sprite := MeteorSprite

	// Figure out the target position — the screen center, in this case
	target := Vector{
		X: ScreenWidth / 2,
		Y: ScreenHeight / 2,
	}

	// The distance from the center the meteor should spawn at — half the width
	r := ScreenWidth / 2.0

	// Pick a random angle — 2π is 360° — so this returns 0° to 360°
	angle := rand.Float64() * 2 * math.Pi

	// Figure out the spawn position by moving r pixels from the target at the chosen angle
	pos := Vector{
		X: target.X + math.Cos(angle)*r,
		Y: target.Y + math.Sin(angle)*r,
	}

	// Randomized velocity
	velocity := (0.25 + rand.Float64()*1.5) + (float64(gameSpeed)/10)

	// Direction is the target minus the current position
	direction := Vector{
		X: target.X - pos.X,
		Y: target.Y - pos.Y,
	}

	// Normalize the vector — get just the direction without the length
	normalizedDirection := direction.Normalize()

	// Multiply the direction by velocity
	movement := Vector{
		X: normalizedDirection.X * velocity,
		Y: normalizedDirection.Y * velocity,
	}

	rotationSpeed := -0.02 + rand.Float64()*0.04

	return &Meteor{
		position:      pos,
		movement:      movement,
		sprite:        sprite,
		rotationSpeed: rotationSpeed,
		boundary: Circle{
			X: pos.X+65,
			Y: pos.Y+65, 
			Radius: 35.0,
		},
	}
}

func (m *Meteor) Intersect(c Circle) bool {
	return m.Intersect(c)
}

func (m *Meteor) Update() error {
	m.position.X += m.movement.X
	m.position.Y += m.movement.Y
	m.rotation += m.rotationSpeed
	m.boundary.X = m.position.X + 65
	m.boundary.Y = m.position.Y + 65
	return nil
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	bounds := m.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(m.rotation)
	op.GeoM.Translate(halfW, halfH)
	op.GeoM.Translate(m.position.X, m.position.Y)
	
    strokeWidth := 2.0
    clr := color.RGBA{0, 255, 0, 255} // Green
    vector.StrokeCircle(screen, float32(m.boundary.X), float32(m.boundary.Y), float32(m.boundary.Radius), float32(strokeWidth), clr, true)


	screen.DrawImage(m.sprite, op)
}

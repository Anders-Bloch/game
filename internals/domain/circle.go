package domain

import "math"

type Circle struct {
	X, Y, Radius float64
}

// Intersect checks if two circles intersect
func (c *Circle) Intersect(c2 Circle) bool {
	dx := c.X - c2.X
	dy := c.Y - c2.Y
	distanceSquared := dx*dx + dy*dy
	radiusSum := c.Radius + c2.Radius
	
	// Avoid SQRT for performance
	return distanceSquared <= radiusSum*radiusSum && 
	       distanceSquared >= math.Pow(c.Radius-c2.Radius, 2)
}
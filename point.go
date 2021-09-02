package main

import (
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type Point struct {
	Fixed       bool
	Position    pixel.Vec
	OldPosition pixel.Vec
	Color       color.RGBA
}

func NewPoint(x, y float64, f bool) *Point {
	pos := pixel.Vec{
		X: x,
		Y: y,
	}
	return &Point{
		Fixed:       f,
		Position:    pos,
		OldPosition: pos,
		Color:       colornames.Orange,
	}
}

func (p *Point) Distance(other *Point) float64 {
	diff := p.Position.Sub(other.Position)
	dist := math.Sqrt(diff.X*diff.X + diff.Y*diff.Y)
	return dist
}

func (p *Point) Update() {
	if p.Fixed {
		return
	}
	// Retrieve velocity as delta of space
	vel := p.Position.Sub(p.OldPosition)
	// Apply friction
	vel = vel.Scaled(Friction)
	// fmt.Println(p.Position, p.OldPosition, vel)
	p.OldPosition = p.Position
	// Apply velocity
	p.Position = p.Position.Add(vel)
	// Apply gravity
	p.Position = p.Position.Sub(pixel.V(0, Gravity))

	// Check bounds
	if p.Position.X > windowWidth {
		p.Position.X = windowWidth
		p.OldPosition.X = p.Position.X + vel.X
	} else if p.Position.X < 0 {
		p.Position.X = 0
		p.OldPosition.X = p.Position.X + vel.X
	}
	if p.Position.Y > windowHeight {
		p.Position.Y = windowHeight
		p.OldPosition.Y = p.Position.Y + vel.Y
	} else if p.Position.Y < 0 {
		p.Position.Y = 0
		p.OldPosition.Y = p.Position.Y + vel.Y
	}
}

func (p *Point) Draw(imd *imdraw.IMDraw) {
	imd.Color = p.Color
	imd.Push(pixel.V(p.Position.X, p.Position.Y))
	imd.Circle(5, 0)
}

package component

import (
  "github.com/veandco/go-sdl2/sdl"
  "input"
)

const MAX = 100

type Component int
type Mask Component

const (
  COMPONENT_NONE     = 0
  COMPONENT_POSITION = 1 << 0
  COMPONENT_VELOCITY = 1 << 1
  COMPONENT_CONTROL  = 1 << 2
  COMPONENT_BOX      = 1 << 3
)

type Vector2d struct {
  X float64
  Y float64
}

type Box struct {
	X float64
	Y float64
	W float64
	H float64
}

func (a *Box) Intersects(b *Box) bool {
	if a.X >= b.X + b.W || b.X >= a.X + a.W ||
	   a.Y >= b.Y + b.H || b.Y >= a.Y + a.H {
		return false
	}

	return true
}

func (b *Box) ToRect() sdl.Rect {
	return sdl.Rect{int32(b.X), int32(b.Y), int32(b.W), int32(b.H)}
}

func (v Vector2d) ToInt32() (int32, int32) {
  return int32(v.X), int32(v.Y)
}

type Collection struct {

  Mask[MAX] Mask

  Position[MAX] Vector2d
  Velocity[MAX] Vector2d
  Intersections[MAX] Vector2d
  Box[MAX] Box

  Walls[100] Box
}

func (c *Collection) Alloc() int {
  for i := 0; i < MAX; i++ {
    if c.Mask[i] == 0 {
      return i
    }
  }

  return MAX
}

func (c *Collection) CreateWall(x, y, w, h float64) {
  for i, wall := range c.Walls {
    if wall.W == 0 || wall.H == 0 {
      c.Walls[i] = Box{x, y, w, h}
      return
    }
  }
}

func (c *Collection) Set(pos int, mask Mask) {
  c.Mask[pos] = mask
}

func (c *Collection) Update(i *input.Input) {
  UpdateControl(c, i)
  UpdateMoving(c)
}

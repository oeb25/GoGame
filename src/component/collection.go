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
)

type Vector2d struct {
  X float64
  Y float64
}

func (v Vector2d) ToInt32() (int32, int32) {
  return int32(v.X), int32(v.Y)
}

type Collection struct {

  Mask[MAX] Mask

  Position[MAX] Vector2d
  Velocity[MAX] Vector2d
  Intersections[MAX] Vector2d


  Walls[100] sdl.Rect
}

func (c *Collection) Alloc() int {
  for i := 0; i < MAX; i++ {
    if c.Mask[i] == 0 {
      return i
    }
  }

  return MAX
}

func (c *Collection) CreateWall(x, y, w, h int32) {
  for i, wall := range c.Walls {
    if wall.W == 0 || wall.H == 0 {
      c.Walls[i] = sdl.Rect{x, y, w, h}
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
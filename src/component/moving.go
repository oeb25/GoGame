package component

import (
  "github.com/veandco/go-sdl2/sdl"
)

const MASK_MOVING = COMPONENT_POSITION | COMPONENT_VELOCITY

func UpdateMoving(c *Collection) {
  for i, mask := range c.Mask {
    if (mask & MASK_MOVING) != MASK_MOVING {
      continue
    }

    pos := &c.Position[i]
    vel := &c.Velocity[i]
    intersections := &c.Intersections[i]

    intersections.X = 0
    intersections.Y = 0

    rect := sdl.Rect{0, 0, 100, 100}

    pos.X += vel.X
    rect.X = int32(pos.X)
    rect.Y = int32(pos.Y)

    for _, wall := range c.Walls {
      d, hit := rect.Intersect(&wall)

      if hit {
        if d.X == wall.X {
          pos.X -= float64(d.W)
          intersections.X = 1
        } else {
          pos.X = float64(d.X + d.W)
          intersections.X = -1
        }
      }
    }

    pos.Y += vel.Y
    rect.X = int32(pos.X)
    rect.Y = int32(pos.Y)

    for _, wall := range c.Walls {
      d, hit := rect.Intersect(&wall)

      if hit {
        if d.Y == wall.Y {
          pos.Y -= float64(d.H)
          intersections.Y = 1
        } else {
          pos.Y = float64(d.Y + d.H)
          intersections.Y = -1
        }
      }
    }
  }
}
package component

import (
  "input"
  "github.com/veandco/go-sdl2/sdl"
  // "fmt"
)

const MASK_CONTROL = COMPONENT_VELOCITY | COMPONENT_CONTROL

func UpdateControl(c *Collection, inp *input.Input) {
  for i, mask := range c.Mask {
    if (mask & MASK_CONTROL) != MASK_CONTROL {
      continue
    }

    if inp.Keyboard.Held(sdl.K_RIGHT) {
      c.Velocity[i].X += 1
    }
    if inp.Keyboard.Held(sdl.K_LEFT) {
      c.Velocity[i].X -= 1
    }

    if inp.Keyboard.Press(sdl.K_SPACE) && c.Intersections[i].Y == 1 {
    	c.Velocity[i].Y = -8
    }
  }
}

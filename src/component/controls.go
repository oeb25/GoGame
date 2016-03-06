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

    c.Velocity[i] = Vector2d(inp.Controller.Left.Scale(3))

    if inp.Keyboard.Held(sdl.K_RIGHT) {
      c.Velocity[i].X += 3
    }
    if inp.Keyboard.Held(sdl.K_LEFT) {
      c.Velocity[i].X -= 3
    }

    if inp.Keyboard.Held(sdl.K_DOWN) {
      c.Velocity[i].Y += 3
    }
    if inp.Keyboard.Held(sdl.K_UP) {
      c.Velocity[i].Y -= 3
    }
  }
}
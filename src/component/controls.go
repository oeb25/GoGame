package component

import (
  "input"
)

const MASK_CONTROL = COMPONENT_VELOCITY | COMPONENT_CONTROL

func UpdateControl(c *Collection, inp *input.Input) {
  for i, mask := range c.Mask {
    if (mask & MASK_CONTROL) != MASK_CONTROL {
      continue
    }

    c.Velocity[i] = Vector2d(inp.Controller.Left.Scale(3))
  }
}
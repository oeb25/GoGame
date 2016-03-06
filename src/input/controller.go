package input

import (
  "github.com/veandco/go-sdl2/sdl"
  "math"
)

type Axis struct {
  X float64
  Y float64
}

func (a Axis) Scale(s float64) Axis {
  return Axis{ a.X * s, a.Y * s }
}

func (a *Axis) Deadzone(amt float64) {
  if math.Abs(a.X) < amt {
    a.X = 0
  }
  if math.Abs(a.Y) < amt {
    a.Y = 0
  }
}

type Controller struct {
  c *sdl.GameController

  Active bool

  A State
  B State
  X State
  Y State

  BACK State
  START State

  Left Axis
  Right Axis

  deadzone float64
}

func (c *Controller) HandleEvent(t *sdl.ControllerButtonEvent) {
  switch t.Button {
  case sdl.CONTROLLER_BUTTON_A:
    if t.State == 1 {
      c.A = HELD | PRESS
    } else {
      c.A = RELEASE
    }
  case sdl.CONTROLLER_BUTTON_B:
    if t.State == 1 {
      c.B = HELD | PRESS
    } else {
      c.B = RELEASE
    }
  case sdl.CONTROLLER_BUTTON_X:
    if t.State == 1 {
      c.X = HELD | PRESS
    } else {
      c.X = RELEASE
    }
  case sdl.CONTROLLER_BUTTON_Y:
    if t.State == 1 {
      c.Y = HELD | PRESS
    } else {
      c.Y = RELEASE
    }

  case sdl.CONTROLLER_BUTTON_BACK:
    if t.State == 1 {
      c.BACK = HELD | PRESS
    } else {
      c.BACK = RELEASE
    }

  case sdl.CONTROLLER_BUTTON_START:
    if t.State == 1 {
      c.START = HELD | PRESS
    } else {
      c.START = RELEASE
    }
  }
}

func (c *Controller) Update() {
  // c.A = c.c.GetButton(sdl.CONTROLLER_BUTTON_A) == 1
  // c.B = c.c.GetButton(sdl.CONTROLLER_BUTTON_B) == 1
  // c.X = c.c.GetButton(sdl.CONTROLLER_BUTTON_X) == 1
  // c.Y = c.c.GetButton(sdl.CONTROLLER_BUTTON_Y) == 1

  // c.Back = c.c.GetButton(sdl.CONTROLLER_BUTTON_BACK) == 1
  // c.Start = c.c.GetButton(sdl.CONTROLLER_BUTTON_START) == 1

  c.Left.X = float64(c.c.GetAxis(sdl.CONTROLLER_AXIS_LEFTX)) / 32768
  c.Left.Y = float64(c.c.GetAxis(sdl.CONTROLLER_AXIS_LEFTY)) / 32768
  c.Right.X = float64(c.c.GetAxis(sdl.CONTROLLER_AXIS_RIGHTX)) / 32768
  c.Right.Y = float64(c.c.GetAxis(sdl.CONTROLLER_AXIS_RIGHTY)) / 32768

  c.Left.Deadzone(c.deadzone)
  c.Right.Deadzone(c.deadzone)
}

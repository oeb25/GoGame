package input

import (
  "github.com/veandco/go-sdl2/sdl"
)

type Input struct {
  Controller Controller
}

func Init() Input {
  sdl.Init(sdl.INIT_GAMECONTROLLER)

  i := Input{}

  i.Controller.deadzone = 0.16
  i.Controller.c = sdl.GameControllerOpen(0)

  return i
}

func (i *Input) Update() bool {
  quit := false
  c := &i.Controller

  c.A = c.A & HELD
  c.B = c.B & HELD
  c.X = c.X & HELD
  c.Y = c.Y & HELD

  c.BACK = c.BACK & HELD
  c.START = c.START & HELD

  for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
    switch t := event.(type) {
    case *sdl.QuitEvent:
      quit = true;
    case *sdl.ControllerButtonEvent:
      c.HandleEvent(t)
    }
  }

  c.Update()

  return quit
}

func (i *Input) Destroy() {
  // maybe someday this will do something
}
package input

import (
  "github.com/veandco/go-sdl2/sdl"
)

type State int

const (
  NONE    = 0
  HELD    = 1 << 0
  PRESS   = 1 << 1
  RELEASE = 1 << 2
)


type Input struct {
  Controller Controller
  Keyboard Keyboard
}

func Init() Input {
  sdl.Init(sdl.INIT_GAMECONTROLLER)

  i := Input{}

  i.Controller.deadzone = 0.16
  i.Controller.c = sdl.GameControllerOpen(0)

  i.Keyboard.keys = make(map[sdl.Keycode]State)

  return i
}

func (i *Input) Update() bool {
  quit := false
  c := &i.Controller
  k := &i.Keyboard

  c.A = c.A & HELD
  c.B = c.B & HELD
  c.X = c.X & HELD
  c.Y = c.Y & HELD

  c.BACK = c.BACK & HELD
  c.START = c.START & HELD

  for key := range k.keys {
    k.keys[key] = k.keys[key] & HELD
  }

  for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
    switch t := event.(type) {
    case *sdl.QuitEvent:
      quit = true;
    case *sdl.ControllerButtonEvent:
      c.HandleEvent(t)
    case *sdl.KeyDownEvent:
      if t.Repeat == 0 {
        k.keys[t.Keysym.Sym] = PRESS | HELD
      }
    case *sdl.KeyUpEvent:
      k.keys[t.Keysym.Sym] = RELEASE
    }
  }

  c.Update()

  return quit
}

func (i *Input) Destroy() {
  // maybe someday this will do something
}
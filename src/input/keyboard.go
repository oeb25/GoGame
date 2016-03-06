package input

import (
  "github.com/veandco/go-sdl2/sdl"
)

type Keyboard struct {
  keys map[sdl.Keycode]State
}

func (k *Keyboard) Press(keycode sdl.Keycode) bool {
  return (k.keys[keycode] & PRESS) == PRESS
}

func (k *Keyboard) Release(keycode sdl.Keycode) bool {
  return (k.keys[keycode] & RELEASE) == RELEASE
}

func (k *Keyboard) Held(keycode sdl.Keycode) bool {
  return (k.keys[keycode] & HELD) == HELD
}
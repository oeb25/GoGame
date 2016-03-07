package main

import (
  "graphics"
  "input"
  "component"
  "github.com/veandco/go-sdl2/sdl"
  "fmt"
)

func main() {
  g := graphics.Init()
  defer g.Destroy()

  i := input.Init()
  defer i.Destroy()

  chili := g.LoadTexture("chili.png")

  c := &i.Controller

  world := component.Collection{}

  player := world.Alloc()
  world.Set(player, component.MASK_MOVING | component.MASK_CONTROL)

  world.Box[player] = component.Box{0, 0, 100, 100}

  crate := world.Alloc()

  world.CreateWall(10, 400, 1000, 10)

  fmt.Println(player)
  fmt.Println(crate)

  quit := false

  for !quit {
    if i.Update() ||
      (c.BACK & input.PRESS) == input.PRESS ||
      i.Keyboard.Press(sdl.K_ESCAPE) {
      quit = true
    }

    world.Update(&i)

    // fmt.Println(world.Intersections[player])

    x, y := world.Position[player].ToInt32()

    g.Clear()

    for _, wall := range world.Walls {
    	rect := wall.ToRect()

    	g.DrawRect(&rect)
    }

    g.DrawTexture(chili, x, y)
    g.Flip()
  }
}

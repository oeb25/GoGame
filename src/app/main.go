package main

import (
  "graphics"
  "input"
  "component"
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

  crate := world.Alloc()

  world.CreateWall(100, 100, 10, 10)

  fmt.Println(player)
  fmt.Println(crate)

  quit := false

  for !quit {
    if i.Update() || (c.BACK & input.PRESS) == input.PRESS {
      quit = true
    }

    world.Update(&i)

    x, y := world.Position[player].ToInt32()

    g.Clear()
    g.DrawTexture(chili, x, y)
    g.Flip()
  }
}
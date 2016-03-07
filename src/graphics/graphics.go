package graphics

import (
  "fmt"
  "github.com/veandco/go-sdl2/sdl"
  "github.com/veandco/go-sdl2/sdl_image"
  "os"
  "path"
  "path/filepath"
)

type Graphics struct {
  Renderer *sdl.Renderer
  Window *sdl.Window

  textures map[string]*sdl.Texture

  C *sdl.GameController

  resPath string
}

func Init() Graphics {
  var graphics Graphics

  dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

  graphics.resPath = path.Join(dir, "../Resources")

  sdl.Init(sdl.INIT_EVERYTHING)

  graphics.C = sdl.GameControllerOpen(0)

  window, err := sdl.CreateWindow("Game",
    sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
  if err != nil {
    panic(err)
  }

  renderer, err := sdl.CreateRenderer(window, -1, 0)
  if err != nil {
    panic(err)
  }

  graphics.Window = window
  graphics.Renderer = renderer

  graphics.textures = make(map[string]*sdl.Texture)

  return graphics
}

func (g *Graphics) LoadTexture(src string) *sdl.Texture {
  if g.textures[src] == nil {
    surface, err := img.Load(path.Join(g.resPath, src))
    if err != nil {
      fmt.Println(path.Join(g.resPath, src))
      panic("Image could not be found/loaded.")
    }
    defer surface.Free()

    texture, err := g.Renderer.CreateTextureFromSurface(surface)
    if err != nil {
      panic(err)
    }

    g.textures[src] = texture
  }

  return g.textures[src]
}

func (g *Graphics) Clear() {
	g.Renderer.SetDrawColor(255, 255, 255, 255)
  g.Renderer.Clear()
}

func (g *Graphics) DrawTexture(tex *sdl.Texture, x, y int32) {
  src := sdl.Rect{0, 0, 100, 100}
  rect := sdl.Rect{x, y, 100, 100}

  g.Renderer.Copy(tex, &src, &rect)
}

func (g *Graphics) DrawRect(rect *sdl.Rect) {
	g.Renderer.SetDrawColor(100, 100, 255, 255)
	g.Renderer.FillRect(rect)
}

func (g *Graphics) Flip() {
  g.Renderer.Present()
  sdl.Delay(16)
}

func (g *Graphics) Debug() {
  fmt.Printf("%+v\n", g)
}

func (g *Graphics) Destroy() {
  defer g.Window.Destroy()
  defer g.Renderer.Destroy()

  g.C.Close()

  for i := range g.textures {
    g.textures[i].Destroy()
  }
}

package main

import (
	"image/color"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

var MAP [][]int = [][]int{
	{1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 0, 0, 1, 1},
	{1, 0, 0, 1, 0, 0, 1},
	{1, 0, 1, 1, 1, 0, 1},
	{1, 1, 0, 0, 0, 1, 1},
	{1, 1, 1, 1, 1, 1, 1}}

type vec2F struct {
	X, Y float64
}

type vec2D struct {
	X, Y int32
}

func rotate(dir float64, vec vec2F, dis float64) vec2F {
	var rotatedVec vec2F

	// rotatedVec.X = vec.X*math.Cos(deg) - vec.Y*math.Sin(deg)
	// rotatedVec.Y = vec.X*math.Sin(deg) + vec.Y*math.Cos(deg)
	rotatedVec.X = vec.X + (dis * math.Cos(dir*math.Pi/180))
	rotatedVec.Y = vec.Y + (dis * math.Sin(dir*math.Pi/180))

	return rotatedVec
}

func genRays(vecPlayer vec2F) [1280]float64 {
	var offset float64 = float64(60) / float64(1280)
	var x float64 = 0
	var rayList [1280]float64
	for i := 0; i < 1280; i++ {
		tmp := 0.0
		for {
			currRay := rotate(x, vecPlayer, tmp)
			tmp += 0.001
			vec := vec2D{X: int32(currRay.X), Y: int32(currRay.Y)}
			if MAP[vec.Y][vec.X] != 0 {
				rayList[i] = math.Sqrt(math.Pow(math.Abs(vecPlayer.X-currRay.X), 2) + math.Pow(math.Abs(vecPlayer.Y-currRay.Y), 2))
				break
			}
		}
		x += offset
	}
	return rayList
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	screen := vec2D{X: 1280, Y: 720}
	player := vec2F{X: 3, Y: 4}

	rays := genRays(player)

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screen.X, screen.Y, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)
	surface.Set(0, 0, color.RGBA{R: 255, G: 0, B: 255, A: 100})
	for i, ray := range rays {
		distance := int(math.Floor(720.0 / (ray * math.Cos((30)*math.Pi/180))))
		for y := int(720/2 - distance/2); y < int(720/2+distance/2); y++ {
			surface.Set(i, y, color.RGBA{R: 255, G: 0, B: 255, A: 100})
		}
		//surface.Set(i, int(720/2-distance/2), color.RGBA{R: 255, G: 0, B: 255, A: 100})

	}
	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			print()
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}
}

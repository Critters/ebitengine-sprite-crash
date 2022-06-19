package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
}

//var assets map[string]*img = make(map[string]*img)
var r float64 = 0
var scale, scaleTarget float64 = 1, 1
var renderWidth, renderHeight int = 800, 800
var scaleMin, scaleMax float64 = float64(renderWidth) / 2056, 1024 / float64(renderWidth)

var B *ebiten.Image = ebiten.NewImage(2056, 2056)
var op *ebiten.DrawImageOptions = &ebiten.DrawImageOptions{}

func (g *Game) Update() error {
	r += 0.06
	scaleMin = 1 //4000 / float64(renderHeight)
	scaleMax = 2 //4000 / float64(renderHeight)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op.GeoM.Reset()

	x, y := ebiten.CursorPosition()
	_, wheel := ebiten.Wheel()
	scaleTarget += wheel / 1000
	scaleTarget = math.Max(math.Min(scaleTarget, scaleMax), scaleMin)

	if scale < scaleTarget {
		scale += math.Abs(scale-scaleTarget) / 10
	} else {
		scale -= math.Abs(scale-scaleTarget) / 10
	}

	B.Clear()
	B.DrawImage(LoadImage("./img/map_2056.png"), op)

	var tmpImg *ebiten.Image = LoadImage("./img/godragon_32.png")
	for i := 0; i < 100; i++ {
		for n := 0; n < 50; n++ {
			op.GeoM.Reset()
			op.GeoM.Translate(-16, -16)
			op.GeoM.Rotate(r + float64(i) + float64(n))
			op.GeoM.Translate(float64(i*20)+10, float64(n*40)+20)
			B.DrawImage(tmpImg, op)
		}
	}

	op.GeoM.Reset()
	op.GeoM.Translate((-float64(x)/float64(renderWidth))*2056, (-float64(y)/float64(renderHeight))*2056)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(float64(x), float64(y))

	screen.DrawImage(B, op)

	ebitenutil.DebugPrint(screen, fmt.Sprint(ebiten.CurrentFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	renderWidth = 800
	renderHeight = 800
	return renderWidth, renderHeight
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(renderWidth, renderHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

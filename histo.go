package main

import (
	"fmt"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

type bar struct {
	Value       int
	isComparing bool
	imd         *imdraw.IMDraw
}

func (histo *bar) draw(target pixel.Target, pos int) {
	w := float64(windowWidth) / float64(size)
	h1 := float64(windowHeight) / (float64(size) + 1) * float64(histo.Value+1)
	x0 := w * float64(pos)
	x1 := x0 + w - 1
	c := koColor

	if pos == histo.Value {
		c = okColor
	}

	if histo.isComparing {
		c = ckColor
	}

	imd := histo.imd

	imd.Clear()
	imd.Reset()
	imd.Color = c
	imd.Push(pixel.V(x0, 0))
	imd.Push(pixel.V(x1, h1))
	imd.Rectangle(0)

	imd.Draw(target)

}

var histogram []bar
var statsText *text.Text
var debugText *text.Text

func initHisto() {
	histogram = make([]bar, size)
	for i := 0; i < size; i++ {
		histogram[i] = bar{Value: i, imd: imdraw.New(nil)}
	}

	face := basicfont.Face7x13
	atlas := text.NewAtlas(face, text.ASCII)
	statsText = text.New(pixel.V(10, 600), atlas)
	statsText.Color = colornames.Aqua
	debugText = text.New(pixel.V(200, 600), atlas)
	debugText.Color = colornames.Aqua
}

func drawHisto(target pixel.Target) {
	for i, j := range globalData.array {
		histogram[j].draw(target, i)
	}
	drawStats(target)
}

func drawStats(target pixel.Target) {
	statsText.Clear()
	statsText.WriteString(fmt.Sprintf("%v\n array size: %v\n tests: %v\n swaps: %v", algorithm, size, numOfComparison, numOfSwaps))
	if dataProcessed {
		statsText.WriteString(fmt.Sprintf("\n time: %v\n\n(press ESC to quit)", elapsedTime))
	}
	statsText.Draw(target, pixel.IM.Moved(pixel.V(0, -20)))
}

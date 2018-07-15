package main

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type algo string

var (
	win                                *pixelgl.Window
	windowWidth, windowHeight          float64
	algorithm                          algo
	size                               int
	speed                              int
	koColor, bgColor, okColor, ckColor pixel.RGBA
	runningTime                        time.Time
	elapsedTime                        time.Duration
	debug                              bool
	startDelay                         int
	dataAsc, dataDesc                  bool
	dataRndSeed                        int64
)

const (
	quickSort     algo = "quick_sort"
	bubbleSort    algo = "bubble_sort"
	bubbleSort2   algo = "bubble_sort2"
	insertionSort algo = "insertion_sort"
	selectionSort algo = "selection_sort"
	shellSort     algo = "shell_sort"
	cocktailSort  algo = "cocktail_sort"
	mergeSort     algo = "merge_sort"
)

var algos map[algo]sorter

func init() {
	algos = make(map[algo]sorter)
	algos[quickSort] = quickSortAlgo{}
	algos[bubbleSort] = bubbleSortAlgo{}
	algos[bubbleSort2] = bubbleSort2Algo{}
	algos[insertionSort] = insertionSortAlgo{}
	algos[selectionSort] = selectionSortAlgo{}
	algos[shellSort] = shellSortAlgo{}
	algos[cocktailSort] = cocktailSortAlgo{}
	algos[mergeSort] = mergeSortAlgo{}

	koColor = pixel.RGB(1, 0, 0)
	bgColor = pixel.RGB(0, 0, 0)
	okColor = pixel.RGB(0, 1, 0)
	ckColor = pixel.RGB(1, 1, 0)
}

func main() {
	parseArguments()
	initHisto()

	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:     fmt.Sprintf("Visual Sort [%v]", algorithm),
		Bounds:    pixel.R(0, 0, float64(windowWidth), float64(windowHeight)),
		VSync:     true,
		Resizable: true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	event := time.Tick(time.Duration(speed) * time.Millisecond)
	newData(event)

	var algo sorter
	var found bool

	if algo, found = algos[algorithm]; !found {
		panic("Algorithm not found")
	}

	go func() {
		runningTime = time.Now()
		algo.sort(&data)
		data.endOfWork()
	}()

	time.Sleep(time.Duration(startDelay) * time.Second)

	for !win.Closed() {
		if data.dataProcessed {
			time.Sleep(50 * time.Millisecond)
			if win.Pressed(pixelgl.KeyEscape) {
				break
			}
		}
		win.Clear(bgColor)
		drawHisto(win)
		win.Update()
	}
}

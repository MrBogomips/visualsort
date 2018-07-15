package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

func parseArguments() {
	showError := func(msg string) {
		fmt.Println(msg)
		flag.Usage()
		os.Exit(1)
	}
	helpAlgos := func() string {
		as := make([]string, len(algos))
		i := 0
		for k := range algos {
			as[i] = fmt.Sprintf("%q", string(k))
			i++
		}
		sort.Strings(as)
		return strings.Join(as, ", ")
	}

	var showHelp bool
	var _algorithm string
	var _windowWidth int
	var _windowHeight int

	flag.BoolVar(&showHelp, "help", false, "Show usage info")
	flag.BoolVar(&showHelp, "h", false, "Show usage info")
	flag.IntVar(&_windowWidth, "width", 800, "Window's width")
	flag.IntVar(&_windowHeight, "height", 600, "Window's height")
	flag.IntVar(&size, "size", 100, "Size of the array to sort")
	flag.IntVar(&speed, "speed", 2, "Step duration in milliseconds")
	flag.BoolVar(&debug, "debug", false, "Enable debug info")
	flag.StringVar(&_algorithm, "algo", string(bubbleSort), "One of: "+helpAlgos())

	flag.Parse()

	if showHelp {
		flag.Usage()
		os.Exit(0)
	}

	if _windowWidth <= 0 {
		showError("Window's width must be positive")
	}
	windowWidth = float64(_windowWidth)
	if _windowHeight <= 0 {
		showError("Window0s height must be positive")
	}
	if size <= 1 {
		showError("Size bust be at least 2")
	}
	windowHeight = float64(_windowHeight)
	algorithm = algo(_algorithm)
	if _, found := algos[algorithm]; !found {
		showError(fmt.Sprintf("Unknown algorithm %q", _algorithm))
	}

}

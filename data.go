package main

import (
	"log"
	"math/rand"
	"time"
)

type Data struct {
	array           []int
	sync            <-chan time.Time
	numOfComparison int
	numOfSwaps      int
	dataProcessed   bool
}

var data Data

func (data *Data) Swap(i, j int) {
	<-data.sync
	setTestedElements(i, j)
	data.numOfSwaps++
	data.array[i], data.array[j] = data.array[j], data.array[i]
}

func (data *Data) Less(i, j int) bool {
	<-data.sync
	setTestedElements(i, j)

	if debug {
		log.Printf("(i, j) = (%v, %v)\n", i, j)
	}

	data.numOfComparison++
	return data.array[i] < data.array[j]
}

// Length returns the length of the underlying data
func (data *Data) Len() int {
	return len(data.array)
}

var prevI, prevJ int
var firstComparison = true

func setTestedElements(i, j int) {
	if firstComparison {
		firstComparison = false
		prevI, prevJ = i, j
	}
	// reset status of previus compared elements
	histogram[data.array[prevI]].isComparing = false
	histogram[data.array[prevJ]].isComparing = false
	histogram[data.array[i]].isComparing = true
	histogram[data.array[j]].isComparing = true
	prevI, prevJ = i, j
}

func (data *Data) endOfWork() {
	histogram[data.array[prevI]].isComparing = false
	histogram[data.array[prevJ]].isComparing = false
	data.dataProcessed = true
	elapsedTime = time.Since(runningTime)
}

func newData(sync <-chan time.Time) {
	rand.Seed(time.Now().UnixNano())
	data.array = rand.Perm(size)
	data.sync = sync
}

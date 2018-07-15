package main

import (
	"log"
	"math/rand"
	"time"
)

var (
	sync            <-chan time.Time
	numOfComparison int
	numOfSwaps      int
	dataProcessed   bool
)

// Data is the interface to the sorting algorithms
type Data struct {
	array    []int
	min, max int // For split and merge operations
}

var globalData Data

func (data *Data) Split(i int) (left, right *Data) {
	left = &Data{
		array: data.array[:i],
		min:   data.min,
		max:   i,
	}
	right = &Data{
		array: data.array[i:],
		min:   i,
		max:   data.max,
	}
	log.Println("Split> i:", i, "data.min:", data.min, "data.max:", data.max)
	return left, right
}

func (data *Data) Merge(l, r *Data) *Data {
	left := make([]int, l.Len())
	right := make([]int, r.Len())
	results := make([]int, 0)
	copy(left, l.array)
	copy(right, r.array)

	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				numOfComparison++
				results = append(results, left[0])
				left = left[1:len(left)]
			} else {
				results = append(results, right[0])
				right = right[1:len(right)]
			}
		} else if len(left) > 0 {
			results = append(results, left[0])
			left = left[1:len(left)]
		} else if len(right) > 0 {
			results = append(results, right[0])
			right = right[1:len(right)]
		}
	}

	log.Println("Merge> l.min:", l.min, "r.max:", r.max)
	d := &Data{
		array: globalData.array[l.min:r.max],
		min:   l.min,
		max:   r.max,
	}
	copy(d.array, results)
	return d
}

func (data *Data) Swap(i, j int) {
	<-sync
	setTestedElements(i, j)
	numOfSwaps++
	data.array[i], data.array[j] = data.array[j], data.array[i]
}

func (data *Data) Less(i, j int) bool {
	<-sync
	setTestedElements(i, j)

	if debug {
		log.Printf("(i, j) = (%v, %v)\n", i, j)
	}

	numOfComparison++
	return data.array[i] < data.array[j]
}

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
	histogram[globalData.array[prevI]].isComparing = false
	histogram[globalData.array[prevJ]].isComparing = false
	histogram[globalData.array[i]].isComparing = true
	histogram[globalData.array[j]].isComparing = true
	prevI, prevJ = i, j
}

func (data *Data) endOfWork() {
	histogram[data.array[prevI]].isComparing = false
	histogram[data.array[prevJ]].isComparing = false
	dataProcessed = true
	elapsedTime = time.Since(runningTime)
}

func newData(s <-chan time.Time) {
	if dataAsc {
		globalData.array = make([]int, size)
		for i := 0; i < size; i++ {
			globalData.array[i] = i
		}
	} else if dataDesc {
		globalData.array = make([]int, size)
		for i := 0; i < size; i++ {
			globalData.array[i] = size - i - 1
		}
	} else {
		if dataRndSeed > 0 {
			rand.Seed(dataRndSeed)
		} else {
			rand.Seed(time.Now().UnixNano())
		}
		globalData.array = rand.Perm(size)
	}
	globalData.min, globalData.max = 0, size
	sync = s
}

package main

import "sort"

// Sorter implements a sorting algo
type sorter interface {
	sort(d *Data)
}

// BubbleSort algorithm
type bubbleSortAlgo struct{}
type insertionSortAlgo struct{}
type shellSortAlgo struct{}
type golibSortAlgo struct{}
type quickSortAlgo struct{} //Hoare Quicksort
type cocktailSortAlgo struct{}
type mergeSortAlgo struct{}
type selectionSortAlgo struct{}

func (algo bubbleSortAlgo) sort(d *Data) {
	swapped := true
	end := 0
	for swapped {
		swapped = false
		for i := 0; i < d.Len()-1-end; i++ {
			if d.Less(i+1, i) {
				d.Swap(i, i+1)
				swapped = true
			}
		}
		end++
	}
}

func (algo insertionSortAlgo) sort(d *Data) {
	for i := 1; i < d.Len(); i++ {
		for j := i; j > 0 && d.Less(j, j-1); j-- {
			d.Swap(j, j-1)
		}
	}
}

func (algo selectionSortAlgo) sort(d *Data) {
	for i := 0; i < d.Len(); i++ {
		min := i
		for j := i + 1; j < d.Len(); j++ {
			if d.Less(j, min) {
				min = j
			}
		}
		d.Swap(i, min)
	}
}

func (algo shellSortAlgo) sort(d *Data) {
	h := 1
	for h < d.Len() {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < d.Len(); i++ {
			for j := i; j >= h && d.Less(j, j-h); j = j - h {
				d.Swap(j, j-h)
			}
		}
		h = h / 3
	}
}

func (algo golibSortAlgo) sort(d *Data) {
	sort.Sort(d)
}

func (algo quickSortAlgo) sort(d *Data) {

	algo.quicksort(d, 0, d.Len()-1)
}

func (algo quickSortAlgo) quicksort(d *Data, from, to int) {
	if from < to {
		mid := algo.partition(d, from, to)
		algo.quicksort(d, from, mid)
		algo.quicksort(d, mid+1, to)
	}
}

func (algo quickSortAlgo) partition(d *Data, from, to int) int {
	i := from - 1
	j := to + 1
	for {
		for i++; d.Less(i, from); i++ {
		}
		for j--; d.Less(from, j); j-- {
		}
		if i >= j {
			return j
		}
		d.Swap(i, j)
	}
}

func (algo mergeSortAlgo) sort(d *Data) {
	if d.Len() <= 1 {
		return
	}

	left, right := d.Split(d.Len() / 2)

	algo.sort(left)
	algo.sort(right)
	d.Merge(left, right)
}

func (algo cocktailSortAlgo) sort(d *Data) {

}

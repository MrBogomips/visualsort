package main

import "sort"

// Sorter implements a sorting algo
type Sorter interface {
	Sort(d *Data)
}

// BubbleSort algorithm
type BubbleSort struct{}
type ShellSort struct{}
type QuickSort struct{}
type CocktailSort struct{}
type MergeSort struct{}

func (algo BubbleSort) Sort(d *Data) {
	for n := d.Len() - 1; n > 0; n-- {
		for j := 0; j < n; j++ {
			if !d.Less(j, n) {
				d.Swap(j, n)
			}
		}
	}
}

// Sort sorts the data
func (algo ShellSort) Sort(d *Data) {
	for i := 0; i < d.Len(); i++ {
		for j := i + 1; j < d.Len(); j++ {
			if !d.Less(i, j) {
				d.Swap(i, j)
			}
		}
	}
}

func (algo QuickSort) Sort(d *Data) {
	sort.Sort(d)
}

func (algo MergeSort) Sort(d *Data) {

}

func (algo CocktailSort) Sort(d *Data) {

}

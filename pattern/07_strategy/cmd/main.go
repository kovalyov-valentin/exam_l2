package main

import (
	"fmt"
	"github.com/kovalyov-valentin/exam_l2/pattern/07_strategy/pkg"
)

func main() {
	nums := []int{2, 4, 6, 8, 0, 1, 3, 5, 7, 9}
	fmt.Println("Массив до сортировки:", nums)
	bubbleSort := &pkg.BubbleSort{}
	insertionSort := &pkg.InsertionSort{}

	ctx := &pkg.Context{}
	ctx.SetSorter(bubbleSort)
	fmt.Println("Bubble Sort:", ctx.ExecuteSort(nums))

	ctx.SetSorter(insertionSort)
	fmt.Println("Insertion Sort:", ctx.ExecuteSort(nums))

}

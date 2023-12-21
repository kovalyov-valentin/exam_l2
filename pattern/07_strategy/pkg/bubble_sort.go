package pkg

// Конкретная стратегия, реализующая интерфейс Sorter
// Представляет свою реализацию алгоритма сортировки
type BubbleSort struct{}

func (bs *BubbleSort) Sort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

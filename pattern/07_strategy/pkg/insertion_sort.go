package pkg

// Конкретная стратегия, реализующая интерфейс Sorter
// Представляет свою реализацию алгоритма сортировки
type InsertionSort struct{}

func (is *InsertionSort) Sort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

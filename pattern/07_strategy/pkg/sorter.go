package pkg

// Sorter представляет интерфейс для всех конкретных стратегий (в данном случае алгоритмов сортировки)
// Каждый алгоритм реализует этот интерфейс
type Sorter interface {
	Sort([]int) []int
}

// Представляет контекст, использующий стратегию
type Context struct {
	sorter Sorter
}

// Данный метод позволяет динамически изменять стратегию
func (c *Context) SetSorter(sorter Sorter) {
	c.sorter = sorter
}

// Данный метод использует текую стратегию для выполнения сортировки
func (c *Context) ExecuteSort(arr []int) []int {
	return c.sorter.Sort(arr)
}

package main

import "github.com/kovalyov-valentin/exam_l2/pattern/02_builder/pkg"

func main() {
	// Наши два вида комплектации
	acerCollector := pkg.GetCollector("acer")
	lenovoCollector := pkg.GetCollector("lenovo")

	// Передаем базовый сборщик для завода
	factory := pkg.NewFactory(acerCollector)
	acerComp := factory.CreateComputer()
	acerComp.Print()

	// Меняем у нашего завода комплектацию. Теперь производим компы lenovo
	factory.SetCollector(lenovoCollector)
	lenovoComp := factory.CreateComputer()
	lenovoComp.Print()

} 
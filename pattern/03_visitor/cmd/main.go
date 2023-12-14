package main

import (
	"fmt"

	"github.com/kovalyov-valentin/exam_l2/pattern/03_visitor/pkg"
)

func main() {
	// Структура объекта представлена срезом []Visitable, содержащим объекты, которые могут быть посещены. 
	// В данном случае, это экземпляры Rice и Pasta.
	products := make([]pkg.Visitable, 2)
	products[0] = &pkg.Rice{
		Product: pkg.Product{
			Price : 32.0, 
			Name: "Some rice",
		},
	}

	products[1] = &pkg.Pasta{
		Product: pkg.Product{
			Price: 40.0,
			Name : "Some pasta",
		},
	}

	// У нас есть два конкретных посетителя: PriceVisitor и NamePrinter
	// PriceVisitor считает сумму цен продуктов
	priceVisitor := &pkg.PriceVisitor{}
	for _, p := range products {
		p.Accept(priceVisitor)
	}

	fmt.Printf("Total: %f\n", priceVisitor.Sum)

	// NamePrinter формирует список имен продуктов
	nameVisitor := &pkg.NamePrinter{}
	for _, p := range products {
		p.Accept(nameVisitor)
	}

	fmt.Printf("\nProduct list:\n--------------\n%s", nameVisitor.ProductList)
}
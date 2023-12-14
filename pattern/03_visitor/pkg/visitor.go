package pkg

import "fmt"

// Интерфейс информации о продукте. Узнаем цену и название продукта
type ProductInfoRetriever interface {
	GetPrice() float32
	GetName() string
}

// Интерфейс посетителя. Есть метод посещения принимающий тип поиска информации о продукте
type Visitor interface {
	Visit(ProductInfoRetriever)
}

// Интерфейс для посетителей. Метод принимает тип посетителя
type Visitable interface {
	Accept(Visitor)
}

//  Продукты
type Product struct {
	Price float32
	Name  string
}

// Метод получения цены
func (p *Product) GetPrice() float32 {
	return p.Price
}

// Метод Accept, нужен чтобы объекты Product и его подтипы принимали Visitor, которые выполняют различные действия в зависимости от своей реализации
func (p *Product) Accept(v Visitor) {
	v.Visit(p)
}

// Метод получения названия продукта
func (p *Product) GetName() string {
	return p.Name
}

// Структуры продуктов рис и макароны
type Rice struct {
	Product
}

type Pasta struct {
	Product
}

// Структура для суммирования наших продуктов
type PriceVisitor struct {
	Sum float32
}

// Метод, который реализует суммирование
func (pv *PriceVisitor) Visit(p ProductInfoRetriever) {
	pv.Sum += p.GetPrice()
}

// Структура для списка продуктов
type NamePrinter struct {
	ProductList string
}

// Метод реализующий вывод списка продуктов
func (n *NamePrinter) Visit(p ProductInfoRetriever) {
	n.ProductList = fmt.Sprintf("%s\n%s", p.GetName(), n.ProductList)
}

package models

import (
	"errors"
	"fmt"
	"time"
)

type Shop struct {
	Name string
	Products []Product 
}

// Функция Sell, которая и является нашим Фасадом. 
// Здесь происходит взаимодействие между банком, картой, магазином и пользователем.
func (s Shop) Sell(user User, product string) error {
	fmt.Println("[Магазин] Запрос к пользователю, для получения остатка по карте")
	time.Sleep(time.Millisecond * 500)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}
	fmt.Printf("[Магазин] Проверка - может ли %s купить товар \n", user.Name)
	time.Sleep(time.Millisecond * 500)
	for _, p := range s.Products {
		if p.Name != product {
			continue
		}
		if p.Price > user.GetBalance() {
			return errors.New("[Магазин] Недостаточно средст для покупки товара")
		}
		fmt.Printf("[Магазин] Товар %s - куплен\n", p.Name)
	}

	return nil 
}

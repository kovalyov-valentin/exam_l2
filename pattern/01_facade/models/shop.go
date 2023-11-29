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

func (s Shop) Sell(user User, product string) error {
	fmt.Println("[Магазин] ЗАпрос к пользователю, для получения остатка по карте")
	time.Sleep(time.Millisecond * 500)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}
	fmt.Printf("[Магазиг] Проверка - может ли %s купить товар \n", user.Name)
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

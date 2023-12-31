package models

import (
	"fmt"
	"time"
)

// Структура карты
type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

// Проверка баланса со стороны карты
func (c Card) CheckBalance() error {
	fmt.Println("[Карта] Запрос в банк для проверки остатка")
	time.Sleep(time.Millisecond * 500)
	return c.Bank.CheckBalance(c.Name)
}

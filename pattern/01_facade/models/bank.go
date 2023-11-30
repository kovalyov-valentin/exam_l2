package models

import (
	"errors"
	"fmt"
	"time"
)

// Струкртура банка
type Bank struct {
	Name string 
	Cards []Card
}

// Функция проверки баланса на стороне банка
func (b Bank) CheckBalance(cardNumber string) error {
	fmt.Printf("[Банк] Получение остатка по карте: %s", cardNumber)
	time.Sleep(time.Millisecond * 500)

	for _, card := range b.Cards {
		if card.Name != cardNumber {
			continue
		} 
		if card.Balance <= 0 {
			return errors.New("[Банк] Недостаточно средств")
		}
	}
	fmt.Println("[Банк] Остаток положительный")
	return nil
}
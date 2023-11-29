package main

import (
	"fmt"

	"github.com/kovalyov-valentin/exam_l2/pattern/01_facade/models"
)

var (
	bank = models.Bank{
		Name:  "ВечноЗеленый Банс",
		Cards: []models.Card{},
	}
	card1 = models.Card{
		Name:    "Card №1",
		Balance: 200,
		Bank:    &bank,
	}
	card2 = models.Card{
		Name:    "Card №2",
		Balance: 5,
		Bank:    &bank,
	}
	user1 = models.User{
		Name: "Consumer1",
		Card: &card1,
	}
	user2 = models.User{
		Name: "Consumer2",
		Card: &card2,
	}

	product = models.Product{
		Name:  "Cheese",
		Price: 150,
	}
	shop = models.Shop{
		Name:     "Fixprice",
		Products: []models.Product{
			product,
		},
	}
)

func main() {
	fmt.Println("[Банк] Выпуск карт")
	bank.Cards = append(bank.Cards, card1, card2)
	fmt.Printf("[%s]\n", user1.Name)
	err := shop.Sell(user1, product.Name) // shop.Sell это и есть facade 
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	fmt.Printf("[%s]\n", user2.Name)
	err = shop.Sell(user2, product.Name)
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
}
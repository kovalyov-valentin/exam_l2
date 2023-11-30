package models

// Наша структура User будет иметь два поля имя и карта
type User struct {
	Name string 
	Card *Card
}

// Функция возвращающая количество денег на карте
func(u User) GetBalance() float64 {
	return u.Card.Balance
}
package pkg

import "fmt"

type ItemRequestState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestState) RequestItem() error {
	return fmt.Errorf("Item already requested")
}

func (i *ItemRequestState) AddItem(count int) error {
	return fmt.Errorf("Item Dispense in progress")
}

func (i *ItemRequestState) InsertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		fmt.Errorf("Inserted money is less.Plese insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

func (i *ItemRequestState) DispenseItem() error {
	return fmt.Errorf("Please insert money first")
}

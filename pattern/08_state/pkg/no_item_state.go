package pkg

import "fmt"

type NoItemState struct {
	vendingMachine *VendingMachine
}

func (i *NoItemState) RequestItem() error {
	return fmt.Errorf("Item out of stack")
}

func (i *NoItemState) AddItem(count int) error {
	fmt.Printf("Adding %d items\n", count)
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

func (i *NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stack")
}

func (i *NoItemState) DispenseItem() error {
	return fmt.Errorf("Item out of stack")
}

package pkg

type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State
	currentState  State
	itemCount     int
	itemPrice     int
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	// Первоначальная инициализация состояний
	hasItemState := &HasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &ItemRequestState{
		vendingMachine: v,
	}
	hasMoneyState := &HasMoneyState{
		vendingMachine: v,
	}
	noItemState := &NoItemState{
		vendingMachine: v,
	}
	// Первоначальная инициализация машины
	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

// Запрос о каком-то товаре
func (v *VendingMachine) RequestItem() error {
	return v.currentState.RequestItem()
}

// Добавление товара
func (v *VendingMachine) AddItem(count int) error {
	return v.currentState.AddItem(count)
}

// Внесение денег
func (v *VendingMachine) InsertMoney(money int) error {
	return v.currentState.InsertMoney(money)
}

// Выдача товара
func (v *VendingMachine) DispenseItem() error {
	return v.currentState.DispenseItem()
}

// Задает состояние
func (v *VendingMachine) setState(s State) {
	v.currentState = s
}

// Когда товар уже выдан, мы фиксируем изменение количества. Добавление новой единицы товара в вендинг
func (v *VendingMachine) incrementItemCount(count int) {
	v.itemCount = v.itemCount + count
}

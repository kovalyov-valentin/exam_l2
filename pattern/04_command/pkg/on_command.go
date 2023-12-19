package pkg

// Это реализация одной из наших объектов команд (Concrete Command), отвечающая за включение.
// Она содержит получателя
type onCommand struct {
	device Device
}

func NewOnCommand(device Device) *onCommand {
	return &onCommand{
		device: device,
	}
}

func (c *onCommand) Execute() {
	c.device.On()
}

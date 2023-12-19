package pkg

// Это реализация одной из наших объектов команд (Concrete Command), отвечающая за выключение.
// Она содержит получателя
type offCommand struct {
	device Device
}

func NewOffCommand(device Device) *offCommand {
	return &offCommand{
		device: device,
	}
}

func (c *offCommand) Execute() {
	c.device.Off()
}

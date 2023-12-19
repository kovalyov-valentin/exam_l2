package pkg

// Button в данной реализации является нашим invoker(отправитель).
// Запрос посылается в invoker, а он передает его встроенный объект-команду.
type button struct {
	command Command
}

func NewButton(command Command) *button {
	return &button{
		command: command,
	}
}

func (b *button) Press() {
	b.command.Execute()
}

package pkg

import "fmt"

type PersonaComputer struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewPersonalComputer() Computer {
	return PersonaComputer{
		Type:    PersonalComputerType,
		Core:    8,
		Memory:  32,
		Monitor: true,
	}
}

func (pc PersonaComputer) GetType() string {
	return pc.Type
}

func (pc PersonaComputer) PrintDetails() {
	fmt.Printf("%s Core: %d Memory: %d Monitor: %t\n", pc.Type, pc.Core, pc.Memory, pc.Monitor)
}

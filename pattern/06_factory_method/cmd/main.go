package main

import "github.com/kovalyov-valentin/exam_l2/pattern/06_factory_method/pkg"

var types = []string{pkg.PersonalComputerType, pkg.ServerType, pkg.NotebookType, "mobile"}

func main() {
	for _, typeName := range types {
		computer := pkg.New(typeName)
		if computer == nil {
			continue
		}
		computer.PrintDetails()
	}
}

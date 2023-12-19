package main

import "github.com/kovalyov-valentin/exam_l2/pattern/04_command/pkg"

func main() {
	tv1 := pkg.NewTv()
	onCommand := pkg.NewOnCommand(tv1)
	offCommand := pkg.NewOffCommand(tv1)

	onButton := pkg.NewButton(onCommand)
	onButton.Press()

	offButton := pkg.NewButton(offCommand)
	offButton.Press()
}

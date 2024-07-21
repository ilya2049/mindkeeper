package main

import (
	"fmt"

	"mindkeeper/internal/app/commands"
	"mindkeeper/internal/infra"
)

func main() {
	getNextStateCommandHandler := commands.NewGetNextStateCommandHandler()
	keyboard := infra.NewKeyboard(getNextStateCommandHandler)

	if err := keyboard.Listen(); err != nil {
		fmt.Println(err)
	}
}

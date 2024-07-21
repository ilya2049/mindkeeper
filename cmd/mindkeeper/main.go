package main

import (
	"fmt"

	"mindkeeper/internal/infra"
)

func main() {
	keyboard := infra.NewKeyboard()

	if err := keyboard.Listen(); err != nil {
		fmt.Println(err)
	}
}

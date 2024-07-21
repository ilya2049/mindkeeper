package infra

import (
	"errors"
	"fmt"
	"mindkeeper/internal/app/commands"

	"github.com/eiannone/keyboard"
)

type Keyboard struct {
	getNextStateCommandHandler commands.GetNextStateCommandHandler
}

func NewKeyboard(
	getNextStateCommandHandler commands.GetNextStateCommandHandler,
) Keyboard {
	return Keyboard{
		getNextStateCommandHandler: getNextStateCommandHandler,
	}
}

func (k Keyboard) Listen() (err error) {
	if err := keyboard.Open(); err != nil {
		return err
	}

	defer func() {
		if closingError := keyboard.Close(); closingError != nil {
			err = errors.Join(err, closingError)
		}
	}()

	fmt.Println("Press ESC to quit. Press ENTER to continue.")

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			return fmt.Errorf("get key: %w", err)
		}

		if key == keyboard.KeyEsc {
			break
		}

		if key == keyboard.KeyEnter {
			fmt.Println(k.getNextStateCommandHandler.Handle())
		}
	}

	return nil
}

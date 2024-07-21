package infra

import (
	"errors"
	"fmt"
	"mindkeeper/internal/app/commands"

	"github.com/eiannone/keyboard"
)

type Keyboard struct {
	getQuestionAnswerCommandHandler commands.GetQuestionAnswerCommandHandler
}

func NewKeyboard(
	getQuestionAnswerCommandHandler commands.GetQuestionAnswerCommandHandler,
) Keyboard {
	return Keyboard{
		getQuestionAnswerCommandHandler: getQuestionAnswerCommandHandler,
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
			fmt.Println(k.getQuestionAnswerCommandHandler.Handle())
		}
	}

	return nil
}

package main

import (
	"fmt"

	"mindkeeper/internal/app/commands"
	"mindkeeper/internal/infra"
)

func main() {
	questionaryUploader := infra.NewQuestionnaireUploader(
		infra.NewQuestionnaireParser(),
		infra.NewQuestionnaireFileNameProvider(),
	)

	aQuestionnaire, err := questionaryUploader.Upload()
	if err != nil {
		fmt.Println("upload a questionnaire", err)

		return
	}

	getNextStateCommandHandler := commands.NewGetQuestionAnswerCommandHandler(aQuestionnaire)
	keyboard := infra.NewKeyboard(getNextStateCommandHandler)

	if err := keyboard.Listen(); err != nil {
		fmt.Println("listen keyboard", err)
	}
}

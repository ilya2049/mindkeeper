package main

import (
	"fmt"

	"mindkeeper/internal/app/commands"
	"mindkeeper/internal/domain/questionnaire"
	"mindkeeper/internal/infra"
)

func main() {
	aQuestionnaire := questionnaire.New([]questionnaire.QuestionnaireEntry{
		{
			Question: "1",
			Answer:   "2",
		},
		{
			Question: "3",
			Answer:   "4",
		},
	})

	getNextStateCommandHandler := commands.NewGetQuestionAnswerCommandHandler(aQuestionnaire)
	keyboard := infra.NewKeyboard(getNextStateCommandHandler)

	if err := keyboard.Listen(); err != nil {
		fmt.Println(err)
	}
}

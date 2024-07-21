package commands

import "mindkeeper/internal/domain/questionnaire"

type GetQuestionAnswerCommandHandler struct {
	questionnaire questionnaire.Questionnaire
}

func NewGetQuestionAnswerCommandHandler(
	questionnaire questionnaire.Questionnaire,
) GetQuestionAnswerCommandHandler {
	return GetQuestionAnswerCommandHandler{
		questionnaire: questionnaire,
	}
}

func (h *GetQuestionAnswerCommandHandler) Handle() string {
	return h.questionnaire.Read()
}

package commands

import "github.com/ilya2049/mindkeeper/internal/domain/questionnaire"

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

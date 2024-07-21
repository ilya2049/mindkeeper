package infra

import "os"

type QuestionnaireFileNameProvider struct {
}

func NewQuestionnaireFileNameProvider() QuestionnaireFileNameProvider {
	return QuestionnaireFileNameProvider{}
}

func (p QuestionnaireFileNameProvider) ProvideQuestionnaireFileName() string {
	if len(os.Args) == 1 {
		return ""
	}

	return os.Args[1]
}

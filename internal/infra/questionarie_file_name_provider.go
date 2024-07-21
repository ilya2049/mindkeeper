package infra

type QuestionnaireFileNameProvider struct {
}

func NewQuestionnaireFileNameProvider() QuestionnaireFileNameProvider {
	return QuestionnaireFileNameProvider{}
}

func (p QuestionnaireFileNameProvider) ProvideQuestionnaireFileName() string {
	return "questionnaire.ini"
}

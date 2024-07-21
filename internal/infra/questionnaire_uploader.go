package infra

import (
	"errors"
	"fmt"
	"os"

	"mindkeeper/internal/domain/questionnaire"
)

type QuestionnaireUploader struct {
	fileNameProvider QuestionnaireFileNameProvider
	parser           QuestionnaireParser
}

func NewQuestionnaireUploader(
	parser QuestionnaireParser,
	fileNameProvider QuestionnaireFileNameProvider,
) QuestionnaireUploader {
	return QuestionnaireUploader{
		parser:           parser,
		fileNameProvider: fileNameProvider,
	}
}

func (u QuestionnaireUploader) Upload() (q questionnaire.Questionnaire, err error) {
	questionnaireFileName := u.fileNameProvider.ProvideQuestionnaireFileName()

	questionnaireFile, err := os.Open(questionnaireFileName)
	if err != nil {
		return questionnaire.Questionnaire{}, fmt.Errorf("read questionnaire file: %w", err)
	}

	defer func() {
		if closingError := questionnaireFile.Close(); closingError != nil {
			err = errors.Join(err, closingError)
		}
	}()

	return u.parser.Parse(questionnaireFile)
}

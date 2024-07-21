package infra

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"

	"mindkeeper/internal/domain/questionnaire"
)

var (
	ErrQuestionAnswerLength = errors.New("failed to split question and answer")
)

type QuestionnaireParser struct {
}

func NewQuestionnaireParser() QuestionnaireParser {
	return QuestionnaireParser{}
}

func (p QuestionnaireParser) Parse(questionnaireFile io.Reader) (questionnaire.Questionnaire, error) {
	scanner := bufio.NewScanner(questionnaireFile)

	entries := []questionnaire.QuestionnaireEntry{}

	const (
		question = 0
		answer   = 1
	)

	for line := 1; scanner.Scan(); line++ {
		questionAndAnswer := strings.Split(scanner.Text(), "=")
		if len(questionAndAnswer) != 2 {
			return questionnaire.Questionnaire{}, fmt.Errorf("%w: line %d", ErrQuestionAnswerLength, line)
		}

		entries = append(entries, questionnaire.QuestionnaireEntry{
			Question: strings.TrimSpace(questionAndAnswer[question]),
			Answer:   strings.TrimSpace(questionAndAnswer[answer]),
		})
	}

	if err := scanner.Err(); err != nil {
		return questionnaire.Questionnaire{}, fmt.Errorf("scan questionnaire file: %w", err)
	}

	return questionnaire.New(entries), nil
}

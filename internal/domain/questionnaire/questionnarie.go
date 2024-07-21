package questionnaire

type QuestionnaireEntry struct {
	Question string
	Answer   string
}

type Questionnaire struct {
	entries            []QuestionnaireEntry
	answersAlreadyRead int
	isQuestionRead     bool
}

func New(entries []QuestionnaireEntry) Questionnaire {
	return Questionnaire{
		entries: entries,
	}
}

func (q *Questionnaire) Read() string {
	if q.answersAlreadyRead == len(q.entries) {
		q.answersAlreadyRead = 0
	}

	if q.isQuestionRead {
		defer func() {
			q.answersAlreadyRead++
			q.isQuestionRead = false
		}()

		return q.entries[q.answersAlreadyRead].Answer
	}

	q.isQuestionRead = true

	return q.entries[q.answersAlreadyRead].Question
}

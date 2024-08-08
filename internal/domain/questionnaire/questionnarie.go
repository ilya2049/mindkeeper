package questionnaire

import "math/rand/v2"

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
	entriesCopy := make([]QuestionnaireEntry, len(entries))
	copy(entriesCopy, entries)

	q := Questionnaire{
		entries: entriesCopy,
	}

	q.shuffleEntires()

	return q
}

func (q Questionnaire) shuffleEntires() {
	rand.Shuffle(len(q.entries), func(i, j int) {
		q.entries[i], q.entries[j] = q.entries[j], q.entries[i]
	})
}

func (q *Questionnaire) Read() string {
	if q.answersAlreadyRead == len(q.entries) {
		q.answersAlreadyRead = 0
		q.shuffleEntires()
	}

	if q.isQuestionRead {
		defer func() {
			q.answersAlreadyRead++
			q.isQuestionRead = false
		}()

		return q.readAnswer()
	}

	q.isQuestionRead = true

	return q.readQuestion()
}

func (q Questionnaire) readQuestion() string {
	return "Q: " + q.entries[q.answersAlreadyRead].Question
}

func (q Questionnaire) readAnswer() string {
	return "A: " + q.entries[q.answersAlreadyRead].Answer + "\n"
}

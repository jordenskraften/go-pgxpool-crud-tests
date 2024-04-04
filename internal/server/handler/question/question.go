package question

// QuestionResponse структура для отправки вопроса клиенту.
type QuestionResponse struct {
	Id            int    `json:"id"`
	TopicId       int    `json:"topic_id"`
	QuestionTitle string `json:"question_title"`
	QuestionText  string `json:"question_text"`
	AnswerText    string `json:"answer_text"`
}

// ErrorResponse структура для отправки ошибки клиенту.
type ErrorResponse struct {
	Error string `json:"error"`
}

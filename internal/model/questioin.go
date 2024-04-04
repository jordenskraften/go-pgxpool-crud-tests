package model

type Question struct {
	Id             int
	Topic_id       int
	Question_title string
	Question_text  string
	Answer_text    string
}
type QuestionDTO struct {
	TopicID       int    `json:"topic_id"`
	QuestionTitle string `json:"question_title"`
	QuestionText  string `json:"question_text"`
	AnswerText    string `json:"answer_text"`
}

package domain

import (
	"time"
)

// Class-related models -----------------------------------------------------------
type CreateClassParams struct {
	Name string `json:"name"` // Class name
}

type Class struct {
	ID   string `json:"id"`   // Class ID
	Name string `json:"name"` // Class name
}

// Topic-related models ----------------------------------------------------------
type CreateTopicParams struct {
	ClassID string `json:"class_id" binding:"required"`
	Name    string `json:"name" binding:"required"`
}

type Topic struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Lesson-related models --------------------------------------------------------
type CreateLessonParams struct {
	ClassID  string   `json:"class_id"` 
	TopicID  string   `json:"topic_id"`
	VideoURL []string `json:"video_url"`
}


type Lesson struct {
	ID        string     `json:"id"`       
	ClassID   string     `json:"class_id"`  
	TopicID   string     `json:"topic_id"`
	VideoURL  []string   `json:"video_url"`
	CreatedAt time.Time  `json:"created_at"` 
	DeletedAt *time.Time `json:"deleted_at"` 
}

// Question-related models ------------------------------------------------------------------
type CreateQuestionParams struct {
	TopicID          string   `json:"topic_id" validate:"required"`
	QuestionType     string   `json:"question_type" validate:"required,oneof=multiple_choice true_false short_answer"`
	QuestionLevel    string   `json:"question_level" validate:"required,oneof=easy medium hard"`
	QuestionText     string   `json:"question_text" validate:"required"`
	QuestionImageURL *string  `json:"question_image_url,omitempty" validate:"omitempty,url"`
	Options          []string `json:"options" validate:"required,min=2,dive,required"`
	OptionsURL       []string `json:"options_url" validate:"omitempty,dive,url"`
	Answer           string   `json:"answer" validate:"required"`
	Solution         *string  `json:"solution,omitempty" validate:"omitempty"`
	SolutionImageURL *string  `json:"solution_image_url,omitempty" validate:"omitempty,url"`
}

type Question struct {
	ID               string   `json:"id" validate:"required"`
	TopicID          string   `json:"topic_id" validate:"required"`
	QuestionType     string   `json:"question_type" validate:"required,oneof=multiple_choice true_false short_answer"`
	QuestionLevel    string   `json:"question_level" validate:"required,oneof=easy medium hard"`
	QuestionText     string   `json:"question_text" validate:"required"`
	QuestionImageURL *string  `json:"question_image_url,omitempty" validate:"omitempty,url"`
	Options          []string `json:"options" validate:"required,min=2,dive,required"`
	OptionsURL       []string `json:"options_url" validate:"omitempty,dive,url"`
	Answer           string   `json:"answer" validate:"required"`
	Solution         *string  `json:"solution,omitempty" validate:"omitempty"`
	SolutionImageURL *string  `json:"solution_image_url,omitempty" validate:"omitempty,url"`
}

//Json data --------------------------------------------------
type Data struct {
	Class CreateClassParams `json:"class"`
	Topic CreateTopicParams `json:"topic"`
	Lesson CreateLessonParams `json:"lesson"`
	Question CreateQuestionParams `json:"question"`
}
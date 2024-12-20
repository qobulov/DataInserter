package domain

import (
	"time"
)

// Class-related models -----------------------------------------------------------
type CreateClassParams struct {
	Name string `json:"name" example:"Class 1"` // Class name
}

type Class struct {
	ID   string `json:"id" example:"d737b072-3c6a-49fb-9b76-e367b2328e97"` // Class ID
	Name string `json:"name" example:"Class 1"`                            // Class name
}

// Chapter-related models ----------------------------------------------------------
type Chapter struct {
	ID      string `json:"id" example:"ce01f2cb-b084-4e73-b433-e3e6fd43f56d"`
	Name    string `json:"name" example:"Chapter 1: Natural Sonlar ustida amallar"`
	ClassID string `json:"class_id" example:"d737b072-3c6a-49fb-9b76-e367b2328e97"`
}

type CreateChapterParams struct {
	ClassID string `json:"class_id" binding:"required" example:"d737b072-3c6a-49fb-9b76-e367b2328e97"`
	Name    string `json:"name" binding:"required" example:"Chapter 1: Natural Sonlar ustida amallar"`
}

// Topic-related models ----------------------------------------------------------
type Topic struct {
	ID        string `json:"id" example:"aecd7715-d7d6-4145-85e0-e6ef606f567f"`
	Title     string `json:"title" binding:"required" example:"Natural Sonlar"`
	Name      string `json:"name" binding:"required" example:"Natural Sonlarni qo'shish"`
	ClassID   string `json:"class_id" example:"d737b072-3c6a-49fb-9b76-e367b2328e97"`
	ChapterID string `json:"chapter_id" example:"ce01f2cb-b084-4e73-b433-e3e6fd43f56d"`
}

type CreateTopicParams struct {
	ClassID   string `json:"class_id" binding:"required" example:"d737b072-3c6a-49fb-9b76-e367b2328e97"`
	ChapterID string `json:"chapter_id" binding:"required" example:"ce01f2cb-b084-4e73-b433-e3e6fd43f56d"`
	Title     string `json:"title" binding:"required" example:"Natural Sonlar"`
	Name      string `json:"name" binding:"required" example:"Natural Sonlarni qo'shish"`
}

// Lesson-related models --------------------------------------------------------
type Lesson struct {
	ID        string     `json:"id" example:"95590b59-2346-4402-96ba-4628ee1d9574"`
	ClassID   string     `json:"class_id" example:"d737b072-3c6a-49fb-9b76-e367b2328e97"`
	TopicID   string     `json:"topic_id" example:"aecd7715-d7d6-4145-85e0-e6ef606f567f"`
	ChapterID string     `json:"chapter_id" example:"ce01f2cb-b084-4e73-b433-e3e6fd43f56d"`
	VideoURL  []string   `json:"video_url" example:"https://example.com/video.mp4"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type CreateLessonParams struct {
	ClassID   string   `json:"class_id" example:"d737b072-3c6a-49fb-9b76-e367b2328e97"`
	TopicID   string   `json:"topic_id" example:"aecd7715-d7d6-4145-85e0-e6ef606f567f"`
	ChapterID string   `json:"chapter_id" example:"ce01f2cb-b084-4e73-b433-e3e6fd43f56d"`
	VideoURL  []string `json:"video_url" example:"https://example.com/video.mp4"`
}

// Question-related models ------------------------------------------------------------------
type Question struct {
	ID               string   `json:"id"`
	TopicID          string   `json:"topic_id"`
	QuestionType     string   `json:"question_type"`
	QuestionLevel    string   `json:"question_level"`
	QuestionText     string   `json:"question_text"`
	QuestionImageURL *string  `json:"question_image_url,omitempty"`
	QuestionVideoURL *string  `json:"question_video_url,omitempty" validate:"omitempty, url"`
	Options          []string `json:"options"`
	OptionsURL       []string `json:"options_url"`
	Answer           string   `json:"answer"`
	Solution         *string  `json:"solution,omitempty"`
	SolutionImageURL *string  `json:"solution_image_url,omitempty"`
	QuestionLanguage string   `json:"question_language"`
}

type CreateQuestionParams struct {
	TopicID          string   `json:"topic_id" validate:"required"`
	QuestionType     string   `json:"question_type" validate:"required,oneof=multiple_choice true_false short_answer"`
	QuestionLevel    string   `json:"question_level" validate:"required,oneof=easy medium hard"`
	QuestionText     string   `json:"question_text" validate:"required"`
	QuestionImageURL *string  `json:"question_image_url,omitempty" validate:"omitempty,url"`
	QuestionVideoURL *string  `json:"question_video_url,omitempty" validate:"omitempty, url"`
	Options          []string `json:"options" validate:"required,min=2,dive,required"`
	OptionsURL       []string `json:"options_url" validate:"omitempty,dive,url"`
	Answer           string   `json:"answer" validate:"required"`
	Solution         *string  `json:"solution,omitempty" validate:"omitempty"`
	SolutionImageURL *string  `json:"solution_image_url,omitempty" validate:"omitempty,url"`
	QuestionLanguage string   `json:"question_language" validate:"required,oneof=english hindi"`
}

// Json data --------------------------------------------------
type Data struct {
	Class    CreateClassParams    `json:"class"`
	Topic    CreateTopicParams    `json:"topic"`
	Lesson   CreateLessonParams   `json:"lesson"`
	Question CreateQuestionParams `json:"question"`
}

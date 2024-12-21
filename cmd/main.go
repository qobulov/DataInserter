package main

import (
	"aljabr/domain"
	logger "aljabr/logs"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
)

type Response struct {
    Data   interface{} `json:"data"`
    Status int         `json:"status"`
}

var Logger = logger.NewLogger()

func main() {
	// Load data from JSON
	data, err := readDataFromJSON("data.json")
	if err != nil {
		log.Fatalf("Error reading data: %v", err)
	}

	// Insert classes and update topics, lessons and questions
	classURL := "http://206.189.142.88:8086/api/classes"
	chapterURL := "http://206.189.142.88:8086/api/chapters"
	topicURL := "http://206.189.142.88:8086/api/topics"
	questionURL := "http://206.189.142.88:8086/api/questions"
	lessonURL := "http://206.189.142.88:8086/api/lessons"

	insertClassDataToAPI(classURL, chapterURL, topicURL, questionURL, lessonURL, data)
}

func insertClassDataToAPI(classURL, chapterURL, topicURL, questionURL, lessonURL string, data []domain.Data) {
	Logger.Info("Inserting classes...")

	for i, item := range data {
		classID, err := sendPostRequest(classURL, item.Class)
		if err != nil {
			Logger.Error("Error inserting class: %v", slog.Any("err", err))
			log.Printf("Error inserting class: %v\n", err)
			continue
		}
		// Update chapter,topic and lesson with new class_id
		data[i].Topic.ClassID = classID
		data[i].Chapter.ClassID = classID
		// data[i].Lesson.ClassID = classID
		log.Printf("Class inserted successfully! ID: %s\n", classID)
	}

	Logger.Info("Classes inserted successfully!")

	insertChapterDataToAPI(chapterURL, topicURL, questionURL, lessonURL, data)
}

func insertChapterDataToAPI(chapterURL, topicURL, questionURL, lessonURL string, data []domain.Data) {
	Logger.Info("Inserting chapters...")

	for i, item := range data {
		chapterID, err := sendPostRequest(chapterURL, item.Chapter)
		if item.Chapter.ClassID == "" {
			log.Printf("Skipping chapter insertion, missing class_id.")
			continue
		}
		
		if err != nil {
			Logger.Error("Error inserting class: %v", slog.Any("err", err))
			log.Printf("Error inserting class: %v\n", err)
			continue
		}
		// Update topic and lesson with new chapter_id
		data[i].Topic.ChapterID = chapterID
		// data[i].Lesson.ChapterID = chapterID
		log.Printf("Chapter inserted successfully! ID: %s\n", chapterID)
	}

	Logger.Info("Chapters inserted successfully!")

	insertTopicDataToAPI(topicURL, questionURL, lessonURL, data)
}

func insertTopicDataToAPI(topicURL, questionURL, lessonURL string, data []domain.Data) {
	Logger.Info("Inserting topics...")

	for i, item := range data {
		topicID, err := sendPostRequest(topicURL, item.Topic)
		if err != nil {
			Logger.Error("Error inserting topic: %v", slog.Any("err", err))
			log.Printf("Error inserting topic: %v\n", err)
			continue
		}
		// Update question and lesson with new topic_id
		data[i].Question.TopicID = topicID
		// data[i].Lesson.TopicID = topicID
		log.Printf("Topic inserted successfully! ID: %s\n", topicID)
	}

	Logger.Info("Topics inserted successfully!")
	lessonURL = "http://206.189.142.88:8086/api/lessons"
	// insertLessonDataToAPI(lessonURL, data)
	insertQuestionDataToAPI(questionURL, data)
}

func insertQuestionDataToAPI(questionURL string, data []domain.Data) {
	Logger.Info("Inserting questions...")

	for _, item := range data {
		id, err := sendPostRequest(questionURL, item.Question)
		if err != nil {
			Logger.Error("Error inserting question: %v", slog.Any("err", err))
			log.Printf("Error inserting question: %v\n", err)
		} else {
			log.Printf("Question inserted successfully! ID: %s\n", id)
		}
	}

	Logger.Info("Questions inserted successfully!")
}

func insertLessonDataToAPI(lessonURL string, data []domain.Data) {
	Logger.Info("Inserting lessons...")

	for _, item := range data {
		id, err := sendPostRequest(lessonURL, item) // item.Lesson
		if err != nil {
			Logger.Error("Error inserting lesson: %v", slog.Any("err", err))
			log.Printf("Error inserting lesson: %v\n", err)
		} else {
			log.Printf("Lesson inserted successfully! ID: %s\n", id)
		}
	}

	Logger.Info("Lessons inserted successfully!")
}

func sendPostRequest(url string, body interface{}) (string, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("error marshaling data: %w", err)
	}

	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Parse the response
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	var res Response
	if err := json.Unmarshal(responseBody, &res); err != nil {
		return "", fmt.Errorf("error unmarshaling response: %w", err)
	}

	// Ensure status code is as expected
	if resp.StatusCode != http.StatusCreated {
		log.Printf("Unexpected status code: %d\n", resp.StatusCode)
		log.Printf("Response Body: %s\n", responseBody)
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return res.Data.(string), nil
}

func readDataFromJSON(filePath string) ([]domain.Data, error) {
	var data []domain.Data
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("error unmarshaling file: %w", err)
	}
	return data, nil
}

package main

import (
	"aljabr/domain"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Data   string `json:"data"`
	Status int    `json:"status"`
}

func main() {
	// Load data from JSON
	data, err := readDataFromJSON("data.json")
	if err != nil {
		log.Fatalf("Error reading data: %v", err)
	}

	// Insert classes and update topics with new class IDs
	classURL := "http://54.169.168.148:8086/api/classes"
	topicURL := "http://54.169.168.148:8086/api/topics"
	questionURL := "http://54.169.168.148:8086/api/questions"
	lessonURL := "http://54.169.168.148:8086/api/lessons"
	insertClassDataToAPI(classURL, topicURL, questionURL, lessonURL, data)
}

func insertClassDataToAPI(classURL, topicURL, questionURL, lessonURL string, data []domain.Data) {
	for i, item := range data {
		classID, err := sendPostRequest(classURL, item.Class)
		if err != nil {
			log.Printf("Error inserting class: %v\n", err)
			continue
		}
		// Update topic with new class_id
		data[i].Topic.ClassID = classID
		data[i].Lesson.ClassID = classID
		log.Printf("Class inserted successfully! ID: %s\n", classID)
	}

	insertTopicDataToAPI(topicURL, questionURL, lessonURL, data)
	// log.Println("Data:", data)
}

func insertTopicDataToAPI(topicURL, questionURL, lessonURL string, data []domain.Data) {
	for i, item := range data {
		topicID, err := sendPostRequest(topicURL, item.Topic)
		if err != nil {
			log.Printf("Error inserting topic: %v\n", err)
			continue
		}
		// Update question with new topic_id
		data[i].Question.TopicID = topicID
		data[i].Lesson.TopicID = topicID
		log.Printf("Topic inserted successfully! ID: %s\n", topicID)
	}

	insertLessonDataToAPI(lessonURL, data)
	insertQuestionDataToAPI(questionURL, data)
}

func insertQuestionDataToAPI(questionURL string, data []domain.Data) {
	for _, item := range data {
		id, err := sendPostRequest(questionURL, item.Question)
		if err != nil {
			log.Printf("Error inserting question: %v\n", err)
		} else {
			log.Printf("Question inserted successfully! ID: %s\n", id)
		}
	}
}

func insertLessonDataToAPI(lessonURL string, data []domain.Data) {
	for _, item := range data {
		id, err := sendPostRequest(lessonURL, item.Lesson)
		if err != nil {
			log.Printf("Error inserting lesson: %v\n", err)
		} else {
			log.Printf("Lesson inserted successfully! ID: %s\n", id)
		}
	}
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
	

	return res.Data, nil
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

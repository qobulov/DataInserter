# JSON Data Inserter

This project is a utility for inserting educational content from JSON files into a server, ensuring the completeness of metadata for classes, topics, lessons, and questions.

## Features

- Inserts lessons, topics, and questions into a server.
- Processes JSON data in a structured format.
- Supports multimedia links (images, videos) for lessons and solutions.

## JSON Structure

The JSON input adheres to the following structure:

```json
[
  {
    "class": { "name": "Class Name" },
    "topic": { "id": "Unique Topic ID", "name": "Topic Name" },
    "lesson": {
      "class_id": "Class ID",
      "topic_id": "Topic ID",
      "video_url": ["Video URL 1", "Video URL 2"]
    },
    "question": {
      "topic_id": "Topic ID",
      "question_type": "Question Type",
      "question_level": "Difficulty Level",
      "question_text": "Question Text",
      "question_image_url": "Image URL",
      "options": ["Option 1", "Option 2", "Option 3"],
      "options_url": ["Option Image URL 1", "Option Image URL 2"],
      "answer": "Correct Answer",
      "solution": "Explanation",
      "solution_image_url": "Solution Image URL"
    }
  }
]


## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/qobulov/DataInserter.git
   ```
2. Navigate to the project directory:
   ```bash
   cd DataInserter
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```

## Usage

1. Run the tool:
   ```bash
   go run cmd/main.go
   ```
2. Provide the PDF file for processing.
3. The JSON output will be generated in the specified directory.

## Contribution

Contributions are welcome! Feel free to fork the repository, make changes, and submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
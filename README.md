# JSON Latex Data Inserter

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
    "class": {
      "name": "<Class Between 1-8>"
    },
    "chapter":{
      "class_id":"<name>",
      "name": "<Chapter Name>"
    },
    "topic": {
      "class_id": "<Class ID>",
      "chapter_id":"<Chapter ID>",
      "title": "<Topic Collection Title>",
      "name": "<Topic Name>"
    },
    "question": {
      "topic_id": "<Topic ID>",
      "question_type": "<Question Type>",
      "question_level": "<Difficulty Level>",
      "question_text": "<Question Text> in latex",
      "question_image_url": "<Image URL>",
      "question_language": "<Question Language>",
      "options": ["<Option 1>", "<Option 2>", "<Option 3> ++ ... if needed (in latex)"],
      "options_url": ["<Option Image URL 1>", "<Option Image URL 2> ... in latex if needed"],
      "answer": ["<Correct Answer 1>","<Correct Answer 2> ... if needed (in latex)"],
      "solution": "<Explanation> in latex if needed",
      "solution_image_url": "<Solution Image URL>"
    }
  }
]
```

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
1. If your `data.json` is ready for the insert then go ahead! 
2. Run the tool:
   ```bash
   go run cmd/main.go
   ```

## Contribution

Contributions are welcome! Feel free to fork the repository, make changes, and submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
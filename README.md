# PDF-to-Text and JSON Conversion Tool

This project is a utility for extracting educational content from PDFs, structuring it into JSON format, and ensuring completeness of metadata for classes, topics, lessons, and questions.

## Features

- Extracts lessons, topics, and questions from PDFs.
- Outputs data in detailed JSON format.
- Includes multimedia links (images, videos) for lessons and solutions.

## Installation

1. Clone the repository:
   \`\`\`bash
   git clone https://github.com/qobulov/DataInserter.git
   \`\`\`
2. Navigate to the project directory:
   \`\`\`bash
   cd DataInserter
   \`\`\`
3. Install dependencies:
   \`\`\`bash
   go mod tidy
   \`\`\`

## Usage

1. Run the tool:
   \`\`\`bash
   go run main.go
   \`\`\`
2. Provide the PDF file for processing.
3. The JSON output will be generated in the specified directory.

## Contribution

Contributions are welcome! Feel free to fork the repository, make changes, and submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
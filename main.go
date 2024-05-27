package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/jpoz/groq"
)

const (
	llama3_8b     = "llama3-8b-8192"
	llama3_70b    = "llama3-70b-8192"
	mixtral       = "mixtral-8x7b-32768"
	gemma_7b      = "gemma-7b-it"
	whisper_large = "whisper-large-v3"
)

func main() {
	q := "Hai groq"

	model := flag.String("model", mixtral, "The model to use. Available models: llama3_8b, llama3_70b, mixtral, gemma_7b, whisper_large")
	flag.Parse()

	// Combine remaining arguments to form the query
	args := flag.Args()
	if len(args) > 0 {
		q = strings.Join(args, " ")
	}

	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		log.Fatal("GROQ_API_KEY is not set")
	}

	client := groq.NewClient(groq.WithAPIKey(apiKey))
	chatCompletion, err := client.CreateChatCompletion(groq.CompletionCreateParams{
		Model: SwitchModel(*model),
		Messages: []groq.Message{
			{
				Role:    "user",
				Content: q,
			},
		},
		Stream: true,
	})
	if err != nil {
		panic(err)
	}

	// Determine the path to the home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error determining home directory: %v", err)
	}

	// Create the directory if it does not exist
	outputDir := filepath.Join(homeDir, "groq")
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}

	// Determine the absolute path for the markdown file
	outputFilePath := filepath.Join(outputDir, "response.md")

	// Open or create the markdown file
	file, err := os.Create(outputFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the question to the file
	_, err = file.WriteString(fmt.Sprintf("# Question: %s\n\n", q))
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString("## Response:\n\n")
	if err != nil {
		panic(err)
	}

	// Stream the response to the file and console
	for delta := range chatCompletion.Stream {
		content := delta.Choices[0].Delta.Content

		_, err := file.WriteString(content)
		if err != nil {
			panic(err)
		}
	}

	fmt.Print("\n")

	// Run the glow command on the created markdown file
	runGlow(outputFilePath)
}

func SwitchModel(model string) string {
	switch model {
	case "llama3_8b":
		return llama3_8b
	case "llama3_70b":
		return llama3_70b
	case "gemma_7b":
		return gemma_7b
	case "whisper_large":
		return whisper_large
	default:
		return mixtral
	}
}

func runGlow(filePath string) {
	cmd := exec.Command("glow", filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running glow:", err)
	}
}

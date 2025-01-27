package translators

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"strings"

	// "os"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func geminiTranslate(targetLanguage string, text string, sourceLanguage string) (string, string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return "", "", err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	// Ask the model to respond with JSON.
	model.ResponseMIMEType = "application/json"
	// Specify the schema.
	model.ResponseSchema = &genai.Schema{
		Type:  genai.TypeArray,
		Items: &genai.Schema{Type: genai.TypeString},
	}

	command := "Translate this sentence"
	if sourceLanguage != "" {
		command = command + " from " + sourceLanguage
	} else {
		command = command + " detect the source language code with format language_code|||translated_text."
	}
	command = command + " to " + targetLanguage
	message := "Respond concisely without any introductory or closing remarks, additional comments, or greetings." + command + text
	res, err := model.GenerateContent(ctx, genai.Text(message))
	if err != nil {
		return "", "", err
	}
	result := getResponse(res)

	stringSlice := strings.Split(result, "|||")
	if len(stringSlice) <= 1 {
		return result, "", fmt.Errorf("Invalid response format")
	}
	return stringSlice[1], stringSlice[0], nil
}

func getResponse(resp *genai.GenerateContentResponse) string {
	for _, part := range resp.Candidates[0].Content.Parts {
		if txt, ok := part.(genai.Text); ok {
			var result []string
			if err := json.Unmarshal([]byte(txt), &result); err != nil {
				slog.Error("Error unmarshalling response: " + err.Error())
				return ""
			}
			return result[0]
		}
	}
	return ""
}

package translators

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	straicoAPIURL = "https://api.straico.com/v0/prompt/completion"
	straicoModel  = "openai/gpt-4o-mini"
	timeout       = 10 * time.Second
)

type straicoRequest struct {
	Model   string `json:"model"`
	Message string `json:"message"`
}

type straicoResponse struct {
	Data struct {
		Completion struct {
			Choices []struct {
				Message struct {
					Content string `json:"content"`
				} `json:"message"`
			} `json:"choices"`
		} `json:"completion"`
	} `json:"data"`
}

func buildTranslationCommand(targetLanguage, sourceLanguage string) string {
	command := "Translate this sentence"
	if sourceLanguage != "" && sourceLanguage != "auto" {
		command += " from " + sourceLanguage
	} else {
		command += " detect the source language code with format language_code|||translated_text"
	}
	return command + " to " + targetLanguage
}

func getStraicoAPIKey() (string, error) {
	apiKey, exists := os.LookupEnv("STRAICO_API_KEY")
	if !exists {
		return "", fmt.Errorf("STRAICO_API_KEY environment variable not found")
	}
	return apiKey, nil
}

func makeStraicoRequest(payload *straicoRequest) (*straicoResponse, error) {
	apiKey, err := getStraicoAPIKey()
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	client := &http.Client{Timeout: timeout}
	req, err := http.NewRequest(http.MethodPost, straicoAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d req: %s", resp.StatusCode, jsonData)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	var response straicoResponse
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("Error parsing JSON: err:%v\n res:%s\n req:%s\n", err, body, jsonData)
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &response, nil
}

func parseTranslationResponse(content, sourceLanguage string) (string, string, error) {
	if sourceLanguage != "" {
		return content, sourceLanguage, nil
	}

	parts := strings.Split(content, "|||")
	if len(parts) <= 1 {
		return "", "", fmt.Errorf("invalid response format: expected language_code|||translated_text")
	}

	return strings.TrimSpace(parts[1]), strings.TrimSpace(parts[0]), nil
}

func straicoTranslate(targetLanguage, text, sourceLanguage string) (string, string, error) {
	command := buildTranslationCommand(targetLanguage, sourceLanguage)

	payload := &straicoRequest{
		Model:   straicoModel,
		Message: fmt.Sprintf("Respond concisely without any introductory or closing remarks, additional comments, or greetings.%s : '%s'", command, text),
	}

	response, err := makeStraicoRequest(payload)
	if err != nil {
		return "", "", err
	}

	if len(response.Data.Completion.Choices) == 0 {
		return "", "", fmt.Errorf("no translation choices returned")
	}

	content := response.Data.Completion.Choices[0].Message.Content
	if content == "" {
		return "", "", fmt.Errorf("empty translation content")
	}

	return parseTranslationResponse(content, sourceLanguage)
}

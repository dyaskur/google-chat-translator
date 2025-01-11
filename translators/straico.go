package translators

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func straicoTranslate(targetLanguage string, text string, sourceLanguage string) (string, string, error) {
	url := "https://api.straico.com/v0/prompt/completion"
	method := "POST"
	command := "Translate this sentence"
	if sourceLanguage != "" {
		command = command + " from " + sourceLanguage
	} else {
		command = command + " detect the source language code with format language_code|||translated_text."
	}
	command = command + " to " + targetLanguage

	// JSON payload as a raw string
	payload := `{
		"model": "openai/gpt-4o-mini",
		"message": "Respond concisely without any introductory or closing remarks, additional comments, or greetings.` + command + ` : '` + text + `'"
	}`

	// Convert the payload to a byte buffer
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return "", "", err
	}
	apiKey, ok := os.LookupEnv("STRAICO_API_KEY")
	if !ok {
		return "", "", fmt.Errorf("STRAICO_API_KEY not found")
	}
	// Set headers
	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "application/json")

	// Execute the request
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return "", "", err
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return "", "", err
	}

	// fmt.Println("Response:")
	// fmt.Println(string(body))
	//convert to struct
	// Parse the JSON response
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return "", "", err
	}

	// Navigate to the "content" field
	content := ""
	if data, ok := response["data"].(map[string]interface{}); ok {
		if completion, ok := data["completion"].(map[string]interface{}); ok {
			if choices, ok := completion["choices"].([]interface{}); ok && len(choices) > 0 {
				if choice, ok := choices[0].(map[string]interface{}); ok {
					if message, ok := choice["message"].(map[string]interface{}); ok {
						if contentVal, ok := message["content"].(string); ok {
							content = contentVal
						}
					}
				}
			}
		}
	}

	if content != "" {
		if sourceLanguage != "" {
			return sourceLanguage, content, nil
		}
		stringSlice := strings.Split(content, "|||")
		if len(stringSlice) <= 1 {
			return content, "", fmt.Errorf("Invalid response format")
		}
		return stringSlice[1], stringSlice[0], nil
	} else {
		return "something went wrong", "", fmt.Errorf("Invalid response format")
	}
}

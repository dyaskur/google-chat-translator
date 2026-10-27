package translators

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// MyMemory is free (5k chars/day anonymous, 50k/day with an email), so it takes
// over from Google Translate once the monthly Google budget is used up.
// Spec: https://mymemory.translated.net/doc/spec.php
var myMemoryAPIURL = "https://api.mymemory.translated.net/get"

const (
	myMemoryMaxBytes = 500 // q is limited to 500 UTF-8 bytes per request
	myMemoryTimeout  = 5 * time.Second
)

type myMemoryResponse struct {
	ResponseData struct {
		TranslatedText   string `json:"translatedText"`
		DetectedLanguage string `json:"detectedLanguage"`
	} `json:"responseData"`
	// responseStatus is the number 200 on success but a string like "403" on errors
	ResponseStatus  json.RawMessage `json:"responseStatus"`
	ResponseDetails string          `json:"responseDetails"`
	QuotaFinished   *bool           `json:"quotaFinished"`
}

func myMemoryTranslate(targetLanguage string, text string, sourceLanguage string) (string, string, error) {
	if len(text) > myMemoryMaxBytes {
		return "", "", fmt.Errorf("mymemory: text longer than %d bytes", myMemoryMaxBytes)
	}

	source := myMemoryLanguageCode(sourceLanguage)
	if source == "" || source == "auto" {
		source = "autodetect"
	}

	params := url.Values{}
	params.Set("q", text)
	params.Set("langpair", source+"|"+myMemoryLanguageCode(targetLanguage))
	if email := os.Getenv("MYMEMORY_EMAIL"); email != "" {
		params.Set("de", email)
	}

	client := &http.Client{Timeout: myMemoryTimeout}
	resp, err := client.Get(myMemoryAPIURL + "?" + params.Encode())
	if err != nil {
		return "", "", fmt.Errorf("mymemory: request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("mymemory: unexpected status code: %d", resp.StatusCode)
	}

	var response myMemoryResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", "", fmt.Errorf("mymemory: parsing response: %w", err)
	}

	status := strings.Trim(string(response.ResponseStatus), `"`)
	if status != "200" {
		return "", "", fmt.Errorf("mymemory: status %s: %s", status, response.ResponseDetails)
	}
	if response.QuotaFinished != nil && *response.QuotaFinished {
		return "", "", fmt.Errorf("mymemory: daily quota finished")
	}

	result := response.ResponseData.TranslatedText
	if result == "" {
		return "", "", fmt.Errorf("mymemory: empty translation")
	}
	if strings.HasPrefix(result, "MYMEMORY WARNING") {
		return "", "", fmt.Errorf("mymemory: quota warning")
	}

	detected := sourceLanguage
	if source == "autodetect" {
		detected = normalizeDetectedLanguage(response.ResponseData.DetectedLanguage)
		if detected == "" {
			return "", "", fmt.Errorf("mymemory: no detected language")
		}
	}

	return html.UnescapeString(result), detected, nil
}

// myMemoryLanguageCode maps the app's (Google) language codes to ones MyMemory accepts.
func myMemoryLanguageCode(code string) string {
	switch code {
	case "zh":
		return "zh-CN"
	case "iw":
		return "he"
	}
	return code
}

// normalizeDetectedLanguage turns region codes like "en-GB" into the app's "en".
func normalizeDetectedLanguage(code string) string {
	switch code {
	case "zh-CN":
		return "zh"
	case "zh-TW":
		return code
	}
	if base, _, found := strings.Cut(code, "-"); found {
		return base
	}
	return code
}

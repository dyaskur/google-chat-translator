package translators

import (
	"encoding/json"
	"errors"
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

// errMyMemoryQuota means the daily free quota is used up, so there is no point
// calling MyMemory again until it resets.
var errMyMemoryQuota = errors.New("mymemory: daily quota exhausted")

// myMemoryLanguages are the languages MyMemory translates well. For the rest
// of the app's languages it either rejects the code or returns poor
// translation-memory matches, so those go straight to Google.
var myMemoryLanguages = map[string]bool{
	"ar": true, "bg": true, "bn": true, "ca": true, "cs": true, "da": true,
	"de": true, "el": true, "en": true, "es": true, "et": true, "fa": true,
	"fi": true, "fr": true, "he": true, "hi": true, "hr": true, "hu": true,
	"id": true, "it": true, "ja": true, "ko": true, "lt": true, "lv": true,
	"ms": true, "nl": true, "no": true, "pl": true, "pt": true, "ro": true,
	"ru": true, "sk": true, "sl": true, "sr": true, "sv": true, "th": true,
	"tr": true, "uk": true, "ur": true, "vi": true, "zh": true, "zh-TW": true,
}

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
	if !myMemoryLanguages[targetLanguage] {
		return "", "", fmt.Errorf("mymemory: unsupported target language %q", targetLanguage)
	}
	if sourceLanguage != "" && sourceLanguage != "auto" && !myMemoryLanguages[sourceLanguage] {
		return "", "", fmt.Errorf("mymemory: unsupported source language %q", sourceLanguage)
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

	if resp.StatusCode == http.StatusTooManyRequests {
		return "", "", errMyMemoryQuota
	}
	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("mymemory: unexpected status code: %d", resp.StatusCode)
	}

	var response myMemoryResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", "", fmt.Errorf("mymemory: parsing response: %w", err)
	}

	status := strings.Trim(string(response.ResponseStatus), `"`)
	if status == "429" {
		return "", "", errMyMemoryQuota
	}
	if status != "200" {
		return "", "", fmt.Errorf("mymemory: status %s: %s", status, response.ResponseDetails)
	}
	if response.QuotaFinished != nil && *response.QuotaFinished {
		return "", "", errMyMemoryQuota
	}

	result := response.ResponseData.TranslatedText
	if result == "" {
		return "", "", fmt.Errorf("mymemory: empty translation")
	}
	if strings.HasPrefix(result, "MYMEMORY WARNING") {
		return "", "", errMyMemoryQuota
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

package translators

import (
	"errors"
	"log/slog"
)

type translateFunc func(targetLanguage string, text string, sourceLanguage string) (string, string, error)

// translators are tried in order until one succeeds. To add an AI provider,
// put it before googleTranslate so Google stays the final fallback.
var translators = []translateFunc{
	googleTranslate,
}

func TranslateText(targetLanguage string, text string, sourceLanguage string) (string, string, error) {
	var errs []error
	for _, translate := range translators {
		result, detectedLanguage, err := translate(targetLanguage, text, sourceLanguage)
		if err == nil {
			return result, detectedLanguage, nil
		}
		slog.Warn("translator failed, trying next: " + err.Error())
		errs = append(errs, err)
	}
	return "", "", errors.Join(errs...)
}

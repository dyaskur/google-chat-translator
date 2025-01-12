package translators

import (
	"testing"
)

func TestTranslate(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping integration test")
	}
	result, sourceLanguage, err := geminiTranslate("en", "kowe ki sopo to jane, aku lo wong bagus", "")

	if err != nil {
		t.Errorf("error: %s", err.Error())
	}

	if result == "" {
		t.Errorf("result is empty")
	}

	if sourceLanguage != "jw" && sourceLanguage != "jv" {
		t.Errorf("sourceLanguage: %s", sourceLanguage)
	}
}

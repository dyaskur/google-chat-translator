package translators

import (
	"testing"
)

func TestStraicoTranslate(t *testing.T) {
	result, sourceLanguage, err := straicoTranslate("en", "kowe ki sopo to jane, aku lo wong bagus", "")

	if err != nil {
		t.Errorf("error: %s", err.Error())
	}

	if result == "" {
		t.Errorf("result is empty", result)
	}

	if sourceLanguage != "jv" {
		t.Errorf("sourceLanguage: %s", sourceLanguage)
	}
}

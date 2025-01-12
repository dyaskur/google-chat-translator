package translators

import (
	"testing"
)

func TestStraicoTranslate(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping integration test")
	}
	result, sourceLanguage, err := straicoTranslate("en", "kowe ki sopo to jane, aku lo wong bagus", "")

	if err != nil {
		t.Errorf("error: %s", err.Error())
	}

	if result == "" {
		t.Errorf("result is empty")
	}

	if sourceLanguage != "jw" {
		t.Errorf("sourceLanguage: %s", sourceLanguage)
	}
}

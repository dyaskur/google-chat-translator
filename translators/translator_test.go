package translators

import (
	"errors"
	"testing"
)

func TestTranslateTextFallsBackInOrder(t *testing.T) {
	original := translators
	defer func() { translators = original }()

	var calls []string
	translators = []translateFunc{
		func(string, string, string) (string, string, error) {
			calls = append(calls, "first")
			return "", "", errors.New("first failed")
		},
		func(string, string, string) (string, string, error) {
			calls = append(calls, "second")
			return "hello", "id", nil
		},
	}

	result, sourceLanguage, err := TranslateText("en", "halo", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "hello" || sourceLanguage != "id" {
		t.Errorf("got %q, %q", result, sourceLanguage)
	}
	if len(calls) != 2 {
		t.Errorf("calls: %v", calls)
	}
}

func TestTranslateTextReturnsErrorWhenAllFail(t *testing.T) {
	original := translators
	defer func() { translators = original }()

	translators = []translateFunc{
		func(string, string, string) (string, string, error) { return "", "", errors.New("a") },
		func(string, string, string) (string, string, error) { return "", "", errors.New("b") },
	}

	if _, _, err := TranslateText("en", "halo", ""); err == nil {
		t.Error("expected error")
	}
}

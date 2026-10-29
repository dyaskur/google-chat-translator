package translators

import (
	"errors"
	"strings"
	"testing"
	"time"
)

func stubTranslators(t *testing.T, usage int64) *[]string {
	t.Helper()
	originalGoogle, originalMyMemory, originalUsage, originalNow := google, myMemory, googleUsage, now
	t.Cleanup(func() {
		google, myMemory, googleUsage, now = originalGoogle, originalMyMemory, originalUsage, originalNow
		myMemoryPausedUntil = time.Time{}
	})
	myMemoryPausedUntil = time.Time{}

	var calls []string
	google = func(string, string, string) (string, string, error) {
		calls = append(calls, "google")
		return "from google", "id", nil
	}
	myMemory = func(string, string, string) (string, string, error) {
		calls = append(calls, "mymemory")
		return "from mymemory", "id", nil
	}
	googleUsage = func() int64 { return usage }
	return &calls
}

func TestTranslateTextUsesGoogleUnderMonthlyLimit(t *testing.T) {
	calls := stubTranslators(t, 1000)

	result, _, err := TranslateText("en", "halo", "")
	if err != nil || result != "from google" {
		t.Fatalf("got %q, %v", result, err)
	}
	if strings.Join(*calls, ",") != "google" {
		t.Errorf("calls: %v", *calls)
	}
}

func TestTranslateTextUsesMyMemoryOverMonthlyLimit(t *testing.T) {
	calls := stubTranslators(t, defaultGoogleMonthlyCharLimit)

	result, _, err := TranslateText("en", "halo", "")
	if err != nil || result != "from mymemory" {
		t.Fatalf("got %q, %v", result, err)
	}
	if strings.Join(*calls, ",") != "mymemory" {
		t.Errorf("calls: %v", *calls)
	}
}

func TestTranslateTextFallsBackToGoogleWhenMyMemoryFails(t *testing.T) {
	calls := stubTranslators(t, defaultGoogleMonthlyCharLimit)
	myMemory = func(string, string, string) (string, string, error) {
		*calls = append(*calls, "mymemory")
		return "", "", errors.New("quota finished")
	}

	result, _, err := TranslateText("en", "halo", "")
	if err != nil || result != "from google" {
		t.Fatalf("got %q, %v", result, err)
	}
	if strings.Join(*calls, ",") != "mymemory,google" {
		t.Errorf("calls: %v", *calls)
	}
}

func TestTranslateTextSendsLongTextToGoogle(t *testing.T) {
	calls := stubTranslators(t, defaultGoogleMonthlyCharLimit)

	result, _, err := TranslateText("en", strings.Repeat("a", myMemoryMaxBytes+1), "")
	if err != nil || result != "from google" {
		t.Fatalf("got %q, %v", result, err)
	}
	if strings.Join(*calls, ",") != "google" {
		t.Errorf("calls: %v", *calls)
	}
}

func TestTranslateTextReturnsErrorWhenAllFail(t *testing.T) {
	stubTranslators(t, 0)
	google = func(string, string, string) (string, string, error) { return "", "", errors.New("a") }
	myMemory = func(string, string, string) (string, string, error) { return "", "", errors.New("b") }

	if _, _, err := TranslateText("en", "halo", ""); err == nil {
		t.Error("expected error")
	}
}

func TestGoogleMonthlyCharLimitFromEnv(t *testing.T) {
	t.Setenv("GOOGLE_MONTHLY_CHAR_LIMIT", "1234")
	if limit := googleMonthlyCharLimit(); limit != 1234 {
		t.Errorf("limit: %d", limit)
	}
	t.Setenv("GOOGLE_MONTHLY_CHAR_LIMIT", "")
	if limit := googleMonthlyCharLimit(); limit != defaultGoogleMonthlyCharLimit {
		t.Errorf("limit: %d", limit)
	}
}

func TestTranslateTextSkipsMyMemoryForTheRestOfTheDayAfterQuota(t *testing.T) {
	calls := stubTranslators(t, defaultGoogleMonthlyCharLimit)
	current := time.Date(2026, 10, 28, 3, 38, 0, 0, time.UTC)
	now = func() time.Time { return current }
	myMemory = func(string, string, string) (string, string, error) {
		*calls = append(*calls, "mymemory")
		return "", "", errMyMemoryQuota
	}

	for i := 0; i < 2; i++ {
		if result, _, err := TranslateText("en", "halo", ""); err != nil || result != "from google" {
			t.Fatalf("got %q, %v", result, err)
		}
	}
	if strings.Join(*calls, ",") != "mymemory,google,google" {
		t.Errorf("calls: %v", *calls)
	}

	*calls = nil
	current = time.Date(2026, 10, 29, 0, 0, 1, 0, time.UTC)
	TranslateText("en", "halo", "")
	if strings.Join(*calls, ",") != "mymemory,google" {
		t.Errorf("calls after reset: %v", *calls)
	}
}

func TestTranslateTextDoesNotPauseMyMemoryOnOtherErrors(t *testing.T) {
	calls := stubTranslators(t, defaultGoogleMonthlyCharLimit)
	myMemory = func(string, string, string) (string, string, error) {
		*calls = append(*calls, "mymemory")
		return "", "", errors.New("invalid language")
	}

	TranslateText("en", "halo", "")
	TranslateText("en", "halo", "")
	if strings.Join(*calls, ",") != "mymemory,google,mymemory,google" {
		t.Errorf("calls: %v", *calls)
	}
}

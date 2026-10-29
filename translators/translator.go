package translators

import (
	"errors"
	"log/slog"
	"os"
	"strconv"
	"sync"
	"time"
	"unicode/utf8"

	"yaskur.com/chat-translator/utils"
)

type translateFunc func(targetLanguage string, text string, sourceLanguage string) (string, string, error)

// Google Translate is used until this many characters are sent in a month
// (the first 500k are free), then MyMemory takes over to keep costs down.
const defaultGoogleMonthlyCharLimit = 400000

var (
	google   translateFunc = meteredGoogleTranslate
	myMemory translateFunc = myMemoryTranslate

	googleUsage       = monthlyGoogleUsage
	recordGoogleUsage = addMonthlyGoogleUsage

	now = time.Now

	// MyMemory is skipped until this time after its daily quota runs out.
	myMemoryPausedUntil   time.Time
	myMemoryPausedUntilMu sync.Mutex
)

func TranslateText(targetLanguage string, text string, sourceLanguage string) (string, string, error) {
	var errs []error
	for _, translate := range translatorsFor(text) {
		result, detectedLanguage, err := translate(targetLanguage, text, sourceLanguage)
		if err == nil {
			return result, detectedLanguage, nil
		}
		slog.Warn("translator failed, trying next", "error", err)
		if errors.Is(err, errMyMemoryQuota) {
			pauseMyMemory()
		}
		errs = append(errs, err)
	}
	return "", "", errors.Join(errs...)
}

// translatorsFor returns the translators to try in order for text.
func translatorsFor(text string) []translateFunc {
	if len(text) > myMemoryMaxBytes || myMemoryPaused() {
		return []translateFunc{google}
	}
	if googleUsage() >= googleMonthlyCharLimit() {
		return []translateFunc{myMemory, google}
	}
	return []translateFunc{google, myMemory}
}

func meteredGoogleTranslate(targetLanguage string, text string, sourceLanguage string) (string, string, error) {
	result, detectedLanguage, err := googleTranslate(targetLanguage, text, sourceLanguage)
	if err == nil {
		recordGoogleUsage(utf8.RuneCountInString(text))
	}
	return result, detectedLanguage, err
}

func googleMonthlyCharLimit() int64 {
	if limit, err := strconv.ParseInt(os.Getenv("GOOGLE_MONTHLY_CHAR_LIMIT"), 10, 64); err == nil {
		return limit
	}
	return defaultGoogleMonthlyCharLimit
}

func googleUsageKey() string {
	return "google_translate_chars:" + time.Now().UTC().Format("2006-01")
}

func monthlyGoogleUsage() int64 {
	value, err := utils.GetCache(googleUsageKey())
	if err != nil {
		// no usage recorded yet this month, or Redis is unavailable
		return 0
	}
	usage, _ := strconv.ParseInt(value, 10, 64)
	return usage
}

func addMonthlyGoogleUsage(chars int) {
	if _, err := utils.IncrementCounter(googleUsageKey(), int64(chars), 40*24*time.Hour); err != nil {
		slog.Warn("failed to record google translate usage", "error", err)
	}
}

// pauseMyMemory skips MyMemory until the next UTC midnight, when its daily
// quota resets.
func pauseMyMemory() {
	myMemoryPausedUntilMu.Lock()
	defer myMemoryPausedUntilMu.Unlock()
	myMemoryPausedUntil = now().UTC().Truncate(24 * time.Hour).Add(24 * time.Hour)
}

func myMemoryPaused() bool {
	myMemoryPausedUntilMu.Lock()
	defer myMemoryPausedUntilMu.Unlock()
	return now().Before(myMemoryPausedUntil)
}

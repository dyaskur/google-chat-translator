package translators

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func serveMyMemory(t *testing.T, body string, check func(r *http.Request)) {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if check != nil {
			check(r)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(server.Close)

	original := myMemoryAPIURL
	myMemoryAPIURL = server.URL
	t.Cleanup(func() { myMemoryAPIURL = original })
}

func TestMyMemoryTranslateAutodetect(t *testing.T) {
	t.Setenv("MYMEMORY_EMAIL", "dev@example.com")
	serveMyMemory(t, `{"responseData":{"translatedText":"I&#39;m hungry","match":0.85,"detectedLanguage":"id"},"quotaFinished":false,"responseDetails":"","responseStatus":200}`,
		func(r *http.Request) {
			if got := r.URL.Query().Get("langpair"); got != "autodetect|zh-CN" {
				t.Errorf("langpair: %s", got)
			}
			if got := r.URL.Query().Get("de"); got != "dev@example.com" {
				t.Errorf("de: %s", got)
			}
		})

	result, source, err := myMemoryTranslate("zh", "saya lapar", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "I'm hungry" || source != "id" {
		t.Errorf("got %q, %q", result, source)
	}
}

func TestMyMemoryTranslateKeepsGivenSource(t *testing.T) {
	serveMyMemory(t, `{"responseData":{"translatedText":"hello"},"quotaFinished":false,"responseDetails":"","responseStatus":200}`,
		func(r *http.Request) {
			if got := r.URL.Query().Get("langpair"); got != "id|en" {
				t.Errorf("langpair: %s", got)
			}
		})

	_, source, err := myMemoryTranslate("en", "halo", "id")
	if err != nil || source != "id" {
		t.Errorf("got %q, %v", source, err)
	}
}

func TestMyMemoryTranslateErrors(t *testing.T) {
	cases := map[string]string{
		"invalid language": `{"responseData":{"translatedText":"'JW' IS AN INVALID TARGET LANGUAGE"},"quotaFinished":null,"responseDetails":"'JW' IS AN INVALID TARGET LANGUAGE","responseStatus":"403","matches":""}`,
		"quota finished":   `{"responseData":{"translatedText":"hello"},"quotaFinished":true,"responseDetails":"","responseStatus":200}`,
		"quota warning":    `{"responseData":{"translatedText":"MYMEMORY WARNING: YOU USED ALL AVAILABLE FREE TRANSLATIONS FOR TODAY"},"responseDetails":"","responseStatus":200}`,
		"empty":            `{"responseData":{"translatedText":""},"responseDetails":"","responseStatus":200}`,
	}
	for name, body := range cases {
		t.Run(name, func(t *testing.T) {
			serveMyMemory(t, body, nil)
			if _, _, err := myMemoryTranslate("en", "halo", "id"); err == nil {
				t.Error("expected error")
			}
		})
	}
}

func TestMyMemoryTranslateReportsQuotaExhausted(t *testing.T) {
	cases := map[string]string{
		"quota finished": `{"responseData":{"translatedText":"hello"},"quotaFinished":true,"responseDetails":"","responseStatus":200}`,
		"quota warning":  `{"responseData":{"translatedText":"MYMEMORY WARNING: YOU USED ALL AVAILABLE FREE TRANSLATIONS FOR TODAY"},"responseDetails":"","responseStatus":200}`,
		"status 429":     `{"responseData":{"translatedText":""},"responseDetails":"TOO MANY REQUESTS","responseStatus":"429"}`,
	}
	for name, body := range cases {
		t.Run(name, func(t *testing.T) {
			serveMyMemory(t, body, nil)
			if _, _, err := myMemoryTranslate("en", "halo", "id"); !errors.Is(err, errMyMemoryQuota) {
				t.Errorf("expected quota error, got %v", err)
			}
		})
	}
}

func TestMyMemoryTranslateReportsHTTP429AsQuota(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	t.Cleanup(server.Close)
	original := myMemoryAPIURL
	myMemoryAPIURL = server.URL
	t.Cleanup(func() { myMemoryAPIURL = original })

	if _, _, err := myMemoryTranslate("en", "halo", "id"); !errors.Is(err, errMyMemoryQuota) {
		t.Errorf("expected quota error, got %v", err)
	}
}

func TestMyMemoryTranslateSkipsUnsupportedLanguagesWithoutCallingTheAPI(t *testing.T) {
	serveMyMemory(t, `{}`, func(r *http.Request) {
		t.Errorf("unexpected request: %s", r.URL)
	})

	if _, _, err := myMemoryTranslate("jv", "halo", ""); err == nil {
		t.Error("expected error for unsupported target")
	}
	if _, _, err := myMemoryTranslate("en", "maayong buntag", "ceb"); err == nil {
		t.Error("expected error for unsupported source")
	}
}

func TestMyMemoryTranslateAcceptsHebrewAlias(t *testing.T) {
	serveMyMemory(t, `{"responseData":{"translatedText":"shalom"},"responseStatus":200}`,
		func(r *http.Request) {
			if got := r.URL.Query().Get("langpair"); got != "he|en" {
				t.Errorf("langpair: %s", got)
			}
		})

	if _, _, err := myMemoryTranslate("en", "שלום", "iw"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestMyMemoryTranslateRejectsLongText(t *testing.T) {
	long := make([]byte, myMemoryMaxBytes+1)
	if _, _, err := myMemoryTranslate("en", string(long), "id"); err == nil {
		t.Error("expected error")
	}
}

func TestNormalizeDetectedLanguage(t *testing.T) {
	for in, want := range map[string]string{"id": "id", "en-GB": "en", "zh-CN": "zh", "zh-TW": "zh-TW"} {
		if got := normalizeDetectedLanguage(in); got != want {
			t.Errorf("%s: got %s, want %s", in, got, want)
		}
	}
}

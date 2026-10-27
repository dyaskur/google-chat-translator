package utils

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"testing"
)

func TestCloudLoggingHandlerWritesEachAttrOnce(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(&CloudLoggingHandler{Handler: slog.NewJSONHandler(&buf, nil)})

	logger.Info("hello", "space", "DM")

	// A duplicated key decodes to a single value, so count raw occurrences instead.
	if n := bytes.Count(buf.Bytes(), []byte(`"space"`)); n != 1 {
		t.Fatalf("expected space once, got %d times: %s", n, buf.String())
	}
	var entry map[string]any
	if err := json.Unmarshal(buf.Bytes(), &entry); err != nil {
		t.Fatal(err)
	}
	if entry["severity"] != "Info" {
		t.Fatalf("unexpected severity %v", entry["severity"])
	}
}

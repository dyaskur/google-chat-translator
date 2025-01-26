package handlers

import (
	"encoding/json"
	"google.golang.org/api/chat/v1"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"yaskur.com/chat-translator/types"
)

func TestGetCommandName(t *testing.T) {
	event := types.ChatEvent{
		Message: &chat.Message{
			Annotations: []*chat.Annotation{
				{SlashCommand: &chat.SlashCommandMetadata{CommandName: "translate"}},
			},
			SlashCommand: &chat.SlashCommand{CommandId: 1},
		},
	}
	result := getCommandName(event)
	if result != "translate" {
		t.Errorf("expected 'translate', got '%s'", result)
	}
}

func TestChatHandler(t *testing.T) {
	reqBody := `{"type":"MESSAGE","message":{"text":"Hello"}}`
	req := httptest.NewRequest("POST", "/", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	ChatHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var reply chat.Message
	json.NewDecoder(w.Body).Decode(&reply)
	if reply.Text == "" {
		t.Errorf("expected a reply text, got an empty string")
	}
}

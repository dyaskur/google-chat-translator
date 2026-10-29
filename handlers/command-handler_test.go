package handlers

import (
	"testing"

	"google.golang.org/api/chat/v1"
	"yaskur.com/chat-translator/types"
)

func TestResolveTargetLanguagePrefersCommandName(t *testing.T) {
	event := types.ChatEvent{
		Message: &chat.Message{
			Annotations: []*chat.Annotation{
				{SlashCommand: &chat.SlashCommandMetadata{CommandName: "/korean"}},
			},
			SlashCommand: &chat.SlashCommand{CommandId: 72}, // Konkani's ID
		},
	}
	if got := resolveTargetLanguage(event, 72).Code; got != "ko" {
		t.Errorf("expected ko, got %s", got)
	}
}

func TestResolveTargetLanguageFallsBackToID(t *testing.T) {
	event := types.ChatEvent{Message: &chat.Message{SlashCommand: &chat.SlashCommand{CommandId: 71}}}
	if got := resolveTargetLanguage(event, 71).Code; got != "ko" {
		t.Errorf("expected ko, got %s", got)
	}
}

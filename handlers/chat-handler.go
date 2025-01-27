package handlers

import (
	"encoding/json"
	"google.golang.org/api/chat/v1"
	"log/slog"
	"net/http"
	"os"
	"yaskur.com/chat-translator/types"
	"yaskur.com/chat-translator/utils"
)

func getCommandName(event types.ChatEvent) string {
	if event.Message == nil || event.Message.SlashCommand == nil || event.Message.Annotations == nil {
		return ""
	}
	for _, element := range event.Message.Annotations {
		if element.SlashCommand != nil {
			return element.SlashCommand.CommandName
		}
	}
	return ""
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	var event types.ChatEvent
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		return
	}
	if os.Getenv("DEBUG") == "true" {
		logDebugInfo(event)
	}

	var reply chat.Message
	switch event.Type {
	case "MESSAGE":
		if event.Message != nil {
			reply = handleMessageEvent(event)
		}
	case "CARD_CLICKED", "SUBMIT_FORM":
		reply = ActionHandler(event)
	case "ADDED_TO_SPACE":
		reply = handleAddedToSpaceEvent()
	default:
		http.Error(w, "Unsupported event type", http.StatusNotImplemented)
		return
	}

	// Respond with the constructed chat message.
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(reply); err != nil {
		slog.Error("Failed to send response: " + err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// logDebugInfo logs event details for debugging purposes.
func logDebugInfo(event types.ChatEvent) {
	command := getCommandName(event)
	locale := "?"
	if event.Common != nil {
		locale = event.Common.UserLocale
	}
	slog.Debug("Received event",
		"type", event.Type,
		"time", event.EventTime,
		"user", event.User.DisplayName,
		"email", event.User.Email,
		"space", event.Space.Type,
		"command", command,
		"locale", locale,
	)
}

// handleMessageEvent handles MESSAGE events.
func handleMessageEvent(event types.ChatEvent) chat.Message {
	message := event.Message

	if message.SlashCommand != nil {
		return CommandHandler(event)
	}

	if message.Text != "" {
		locale := "en"
		if event.Common != nil {
			locale = event.Common.UserLocale
		}

		return chat.Message{
			ActionResponse: &chat.ActionResponse{
				Type: "NEW_MESSAGE",
			},
			Text: buildDefaultMessage(locale),
		}
	}

	return chat.Message{}
}

// handleAddedToSpaceEvent handles events when the bot is added to a space.
func handleAddedToSpaceEvent() chat.Message {

	return chat.Message{
		ActionResponse: &chat.ActionResponse{
			Type: "NEW_MESSAGE",
		},
		Text: "Welcome to Abang Translator! I can translate your messages to any language. " +
			"\nPlease use a command to perform translations. Examples:\n" +
			"`/spanish Hello everyone`\n" +
			"`/arabic Semangat menjalani hari, semoga produktif!`\n" +
			"`/japanese ¡Vamos a empezar!`\n" +
			"`/russian Buenos dias`\n" +
			"`/french Wie geht's?`\n" +
			"\nBy default, the original message will be shown. Use `/config` to change this." +
			"\nTo see all available languages, use the `/translate` command.",
	}
}

// buildDefaultMessage constructs a default message for users.
func buildDefaultMessage(locale string) string {
	greeting := utils.GetRandomGreeting(locale)
	instruction := utils.GetRandomInstruction(locale)
	exampleCommand := utils.GetRandomExampleCommand(locale)

	return greeting + "\n" + instruction +
		"\ne.g:\n" +
		"`" + exampleCommand + "`\n" +
		"`/arabic Semangat menjalani hari, semoga produktif!`\n" +
		"`/japanese ¡Vamos a empezar!`\n" +
		"`/russian Buenos dias`\n" +
		"`/french Wie geht's?`\n" +
		"\nBy default, the original message will be shown. Use `/config` to change this."
}

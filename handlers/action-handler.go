package handlers

import (
	"encoding/json"
	"fmt"
	"google.golang.org/api/chat/v1"
	"log/slog"
	"strconv"
	"yaskur.com/chat-translator/cards"
	"yaskur.com/chat-translator/translators"
	"yaskur.com/chat-translator/types"
	"yaskur.com/chat-translator/utils"
)

const (
	ActionSetShowOriginalText = "setShowOriginalText"
	ActionTranslate           = "translate"
)

func ActionHandler(event types.ChatEvent) chat.Message {
	action := event.Common.InvokedFunction
	switch action {
	case ActionSetShowOriginalText:
		return handleSetShowOriginalText(event)
	case ActionTranslate:
		return handleTranslate(event)
	default:
		slog.Warn("Unknown action: %s", action)
		return chat.Message{Text: "Unknown Action"}
	}
}

func handleSetShowOriginalText(event types.ChatEvent) chat.Message {

	var input = event.Common.Parameters["show_original_text"]
	fmt.Println(input)
	showOriginalText, _ := strconv.ParseBool(input)
	config := types.Config{
		ShowOriginalText: showOriginalText,
	}
	configJson, _ := json.Marshal(config)
	utils.SetCache(event.Space.Name, string(configJson))

	return chat.Message{
		ActionResponse: &chat.ActionResponse{
			Type: "DIALOG",
			DialogAction: &chat.DialogAction{
				Dialog: &chat.Dialog{
					Body: cards.ConfigForm(config),
				},
			},
		},
	}
}

func handleTranslate(event types.ChatEvent) chat.Message {
	configKey := "home_" + event.User.Name
	var formInput types.FormInput
	var errorMessage string
	var translatedText string
	var source string
	formInput = extractFormInput(*event.Common)
	err := validateFormInput(formInput)
	if err != nil {
		slog.Warn("invalid validation: %v", err)
		errorMessage = fmt.Sprint(err)
	} else {
		translatedText, source, err = translators.TranslateText(formInput.Target, formInput.Text, formInput.Source)
		if err != nil {
			slog.Error("error translate: %v", err)
			errorMessage = fmt.Sprint(err)
		} else {
			configJson, _ := json.Marshal(formInput)
			utils.SetCache(configKey, string(configJson))
		}
	}
	formInput.Result = translatedText
	return chat.Message{
		ActionResponse: &chat.ActionResponse{
			Type: "DIALOG",
			DialogAction: &chat.DialogAction{
				Dialog: &chat.Dialog{
					Body: cards.TranslateForm(formInput, source, errorMessage).Card,
				},
			},
		},
	}
}

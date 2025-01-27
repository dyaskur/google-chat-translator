package handlers

import (
	"encoding/json"
	"log/slog"

	"google.golang.org/api/chat/v1"
	"yaskur.com/chat-translator/cards"
	"yaskur.com/chat-translator/translators"
	"yaskur.com/chat-translator/types"
	"yaskur.com/chat-translator/utils"
)

func CommandHandler(event types.ChatEvent) chat.Message {
	commandID := int16(event.Message.SlashCommand.CommandId)
	configKey := event.Space.Name

	config := getConfig(configKey)
	switch commandID {
	case 1: // /config
		return handleConfigCommand(config)
	case 2: // /help
		return handleHelpCommand()
	case 3: // /translate
		return handleTranslateCommand(event, configKey)
	default:
		return handleLanguageTranslation(event, config, commandID)
	}
}

func getConfig(configKey string) types.Config {
	configJSON, err := utils.GetCache(configKey)
	if err != nil || configJSON == "" {
		return types.Config{ShowOriginalText: true}
	}

	var config types.Config
	if err := json.Unmarshal([]byte(configJSON), &config); err != nil {
		slog.Error("Failed to unmarshal config: " + err.Error())
		return types.Config{ShowOriginalText: true}
	}

	return config
}

func handleConfigCommand(config types.Config) chat.Message {
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

func handleHelpCommand() chat.Message {
	replyText := "I can translate your messages to any language. " +
		"Please use command to do translation, e.g.:\n" +
		"`/spanish Hello everyone`\n" +
		"`/arabic Semangat menjalani hari, semoga produktif!`\n" +
		"`/japanese ¡Vamos a empezar!`\n" +
		"`/russian Buenos dias`\n" +
		"`/french Wie geht's?`\n" +
		"By default, original message will be shown, use `/config` to change that.\n" +
		"Use `/translate` for a translate form to view all available languages.\n"

	return chat.Message{
		ActionResponse: &chat.ActionResponse{
			Type: "NEW_MESSAGE",
		},
		Text: replyText,
	}
}

func handleTranslateCommand(event types.ChatEvent, configKey string) chat.Message {
	formInput := getFormInput(configKey)

	return chat.Message{
		ActionResponse: &chat.ActionResponse{
			Type: "DIALOG",
			DialogAction: &chat.DialogAction{
				Dialog: &chat.Dialog{
					Body: cards.TranslateForm(formInput, "", "").Card,
				},
			},
		},
	}
}

func getFormInput(configKey string) types.FormInput {
	lastInputJSON, _ := utils.GetCache(configKey)
	var formInput types.FormInput
	if lastInputJSON != "" {
		_ = json.Unmarshal([]byte(lastInputJSON), &formInput)
	}
	return formInput
}

func handleLanguageTranslation(event types.ChatEvent, config types.Config, commandID int16) chat.Message {
	targetLanguage := utils.GetById(commandID)
	translatedText, source, err := translators.TranslateText(targetLanguage.Code, event.Message.ArgumentText, "")
	if err != nil {
		log.Printf("Translation error: %v", err)
		return chat.Message{}
	}

	sourceLanguage := utils.GetByCode(source)
	response := formatResponse(event.User.DisplayName, targetLanguage.Language, translatedText, sourceLanguage.Language, event.Message.ArgumentText, config.ShowOriginalText)

	return chat.Message{
		ActionResponse: &chat.ActionResponse{
			Type: "NEW_MESSAGE",
		},
		Text: response,
	}
}

func formatResponse(userDisplayName, targetLanguage, translatedText, sourceLanguage, originalText string, showOriginal bool) string {
	response := "_" + userDisplayName + " said: (translated to " + targetLanguage + ")_\n" + translatedText
	if showOriginal {
		response += "\nTranslated from " + sourceLanguage + ", original message:\n" + originalText
	}
	return response
}

package handlers

import (
	"encoding/json"
	"google.golang.org/api/chat/v1"
	"log"
	"strconv"
	"yaskur.com/chat-translator/cards"
	"yaskur.com/chat-translator/types"
	"yaskur.com/chat-translator/utils"
)

func CommandHandler(event chat.DeprecatedEvent) chat.Message {
	message := event.Message
	commandId := int16(message.SlashCommand.CommandId)
	log.Printf("commandID: %s", strconv.FormatInt(message.SlashCommand.CommandId, 10))
	configKey := event.Space.Name
	configJson, _ := utils.GetCache(configKey)

	var config types.Config
	if configJson != "" {
		err := json.Unmarshal([]byte(configJson), &config)
		if err != nil {
			panic(err)
		}
	} else {
		config = types.Config{
			ShowOriginalText: true,
		}
	}
	if commandId == 1 {
		// commandId 1 = /config
		reply := chat.Message{
			ActionResponse: &chat.ActionResponse{
				Type: "DIALOG",
				DialogAction: &chat.DialogAction{
					Dialog: &chat.Dialog{
						Body: cards.ConfigForm(config),
					},
				},
			},
		}
		return reply
	} else if commandId == 3 {
		// commandId 1 = /translate

		var formInput types.FormInput
		lastInputJson, _ := utils.GetCache(configKey)
		if lastInputJson != "" {
			err := json.Unmarshal([]byte(lastInputJson), &formInput)
			if err != nil {
				panic(err)
			}
		}
		reply := chat.Message{
			ActionResponse: &chat.ActionResponse{
				Type: "DIALOG",
				DialogAction: &chat.DialogAction{
					Dialog: &chat.Dialog{
						Body: cards.TranslateForm(formInput, "", "").Card,
					},
				},
			},
		}
		return reply
	}

	targetLanguage := utils.GetById(commandId)
	log.Printf("targetLanguage: %s", targetLanguage.Code)
	translatedText, source, err := utils.TranslateText(targetLanguage.Code, message.ArgumentText, "")
	if err != nil {
		log.Fatal(err)
	}
	sourceLanguage := utils.GetByCode(source)

	user := event.User
	response := "_" + user.DisplayName + " said: (translated to " + targetLanguage.Language + ")_\n" + translatedText
	if config.ShowOriginalText {
		response = response + "\nTranslated from " + sourceLanguage.Language + ", original message:\n" + message.ArgumentText
	}
	reply := chat.Message{
		ActionResponse: &chat.ActionResponse{
			Type: "NEW_MESSAGE",
		},
		Text: response,
	}
	return reply
}

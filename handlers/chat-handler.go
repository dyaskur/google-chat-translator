package handlers

import (
	"encoding/json"
	"google.golang.org/api/chat/v1"
	"log"
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
		command := getCommandName(event)
		locale := "err"
		if event.Common != nil {
			locale = event.Common.UserLocale
		}
		log.Printf("type %s; time %s; user %s; email %s; space %s; command %s; locale %s;", event.Type, event.EventTime, event.User.DisplayName, event.User.Email, event.Space.Type, command, locale)

		if event.Message != nil {
			messageJson, _ := json.Marshal(event.Message)
			log.Printf("messageJson %s", messageJson)
		}
	}

	var reply chat.Message
	if event.Type == "MESSAGE" || event.Message != nil {
		message := event.Message
		if message.SlashCommand != nil {
			reply = CommandHandler(event)
		} else if message.Text != "" {
			log.Printf(message.Text)
			locale := "en"
			if event.Common != nil {
				locale = event.Common.UserLocale
			}
			greeting := utils.GetRandomGreeting(locale)
			instruction := utils.GetRandomInstruction(locale)
			reply = chat.Message{
				ActionResponse: &chat.ActionResponse{
					Type: "NEW_MESSAGE",
				},
				Text: greeting + "\n" + instruction +
					"\ne.g: \n" +
					"`/spanish Hello everyone`\n" +
					"`/arabic Semangat menjalani hari, semoga produktif!`\n" +
					"`/japanese ¡Vamos a empezar!`\n" +
					"`/russian Buenos dias`\n" +
					"`/french Wie geht's?`\n" +
					"`\n By default original message will be shown, use `/config` to change that`\n",
			}
		}

	} else if event.Type == "CARD_CLICKED" || event.Type == "SUBMIT_FORM" {
		reply = ActionHandler(event)
	} else if event.Type == "ADDED_TO_SPACE" {

		reply = chat.Message{
			ActionResponse: &chat.ActionResponse{
				Type: "NEW_MESSAGE",
			},
			Text: "Welcome to Abang translator! I can translate your messages to any language. " +
				"Please use command to do translation" + "\ne.g: \n" +
				"`/spanish Hello everyone`\n" +
				"`/arabic Semangat menjalani hari, semoga produktif!`\n" +
				"`/japanese ¡Vamos a empezar!`\n" +
				"`/russian Buenos dias`\n" +
				"`/french Wie geht's?`\n" +
				"`\n By default original message will be shown, use `/config` to change that`\n" +
				"If you want to use translate form and see all available languages use `/translate` command",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(reply)
	if err != nil {
		log.Fatal(err)
	}
}

package handlers

import (
	"encoding/json"
	"google.golang.org/api/chat/v1"
	"log"
	"net/http"
)

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	var event chat.DeprecatedEvent
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		return
	}
	log.Printf(event.EventTime, event.Message.Text)
	log.Printf("%#v", event)
	commonJson, _ := json.Marshal(event.Common.FormInputs)
	actionJson, _ := json.Marshal(event.Action)
	log.Printf("Common %#v", commonJson)
	log.Printf("Action %#v", actionJson)
	var reply chat.Message
	if event.Type == "MESSAGE" {
		message := event.Message
		if message.SlashCommand != nil {
			reply = CommandHandler(event)
		} else if message.Text != "" {
			log.Printf(message.Text)
			reply = chat.Message{
				ActionResponse: &chat.ActionResponse{
					Type: "NEW_MESSAGE",
				},
				Text: "Hello, please use command to do translation" +
					"\ne.g: \n" +
					"`/spanish Hello everyone`\n" +
					"`/arabic Semangat menjalani hari, semoga produktif!`\n" +
					"`/japanese ¡Vamos a empezar!`\n" +
					"`/russian Buenos dias`\n" +
					"`/french Wie geht's?`\n" +
					"`\n By default original message will be shown, use `/config` to change that`\n",
			}
		}

	} else if event.Type == "CARD_CLICKED" {
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
				"`\n By default original message will be shown, use `/config` to change that`\n",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(reply)
	if err != nil {
		log.Fatal(err)
	}
}

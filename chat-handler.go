package translator

import (
	"encoding/json"
	"google.golang.org/api/chat/v1"
	"log"
	"net/http"
)

func chatHandler(w http.ResponseWriter, r *http.Request) {
	var event chat.DeprecatedEvent
	json.NewDecoder(r.Body).Decode(&event)
	log.Printf(event.EventTime, event.Message.Text)
	log.Printf("%#v", event)
	var reply chat.Message
	var err error
	if event.Type == "MESSAGE" {
		message := event.Message
		if message.SlashCommand != nil {
			reply = commandHandler(event)
		} else if message.Text != "" {
			log.Printf(message.Text)
			reply = chat.Message{

				ActionResponse: &chat.ActionResponse{
					Type: "NEW_MESSAGE",
				},
				Text: "Hello, please use command to do translation" +
					"\ne.g: \n" +
					"`/spanish Hello everyone`",
			}
		}

	} else if event.Type == "CARD_CLICKED" {
		reply = chat.Message{
			ActionResponse: &chat.ActionResponse{
				Type: "UPDATE_MESSAGE",
			},
			Text: "You clicked a card",
		}
	} else if event.Type == "ADDED_TO_SPACE" {

		reply = chat.Message{
			ActionResponse: &chat.ActionResponse{
				Type: "NEW_MESSAGE",
			},
			Text: "Welcome to Abang translator! I can translate your messages to any language. " +
				"Please use `/translate` command to do it.",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(reply)
	if err != nil {
		log.Fatal(err)
	}
}

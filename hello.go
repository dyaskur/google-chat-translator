package helloworld

import (
	"encoding/json"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"google.golang.org/api/chat/v1"
	"log"
	"net/http"
)

func init() {
	functions.HTTP("HelloHTTP", HelloHTTP)
}

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func HelloHTTP(w http.ResponseWriter, r *http.Request) {

	var event chat.DeprecatedEvent
	json.NewDecoder(r.Body).Decode(&event)
	log.Printf(event.EventTime, event.Message.Text)
	log.Printf("%#v", event)

	if event.Type == "MESSAGE" {
		commandHandler(event)
	}
	translatedText, err := translateText("en", event.Message.Text)
	if err != nil {
		log.Fatal(err)
	}

	resul := chat.Message{
		ActionResponse: &chat.ActionResponse{
			Type: "NEW_MESSAGE",
		},
		Text: translatedText,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resul)
	if err != nil {
		log.Fatal(err)
	}
}

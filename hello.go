package translator

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"net/http"
	"yaskur.com/chat-translator/handlers"
)

func init() {
	functions.HTTP("HelloHTTP", HelloHTTP)
}

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func HelloHTTP(w http.ResponseWriter, r *http.Request) {

	requestType := r.URL.Query().Get("r")

	if requestType == "chat" {
		handlers.ChatHandler(w, r)
	} else {
		handlers.HomeHandler(w, r)
	}

}

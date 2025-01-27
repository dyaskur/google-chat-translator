package translator

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"log/slog"
	"net/http"
	"os"
	"yaskur.com/chat-translator/handlers"
	"yaskur.com/chat-translator/utils"
)

func init() {

	opts := &slog.HandlerOptions{
		Level:     slog.LevelDebug, // Set minimum level
		AddSource: true,
	}

	jsonHandler := slog.NewJSONHandler(os.Stdout, opts)
	cloudHandler := &utils.CloudLoggingHandler{Handler: jsonHandler}
	logger := slog.New(cloudHandler)

	// Set the default logger so that all slog calls will use this
	slog.SetDefault(logger)

	functions.HTTP("TranslatorHTTP", TranslatorHTTP)

}

// TranslatorHTTP is an HTTP Cloud Function with a request parameter.
func TranslatorHTTP(w http.ResponseWriter, r *http.Request) {
	requestType := r.URL.Query().Get("r")

	if requestType == "chat" {
		handlers.ChatHandler(w, r)
	} else {
		handlers.HomeHandler(w, r)
	}

}

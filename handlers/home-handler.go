package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"google.golang.org/api/chat/v1"
	"yaskur.com/chat-translator/cards"
	"yaskur.com/chat-translator/translators"
	"yaskur.com/chat-translator/types"
	"yaskur.com/chat-translator/utils"
)

type Navigation struct {
	PushCard   *chat.GoogleAppsCardV1Card `json:"push_card,omitempty"`
	UpdateCard *chat.GoogleAppsCardV1Card `json:"update_card,omitempty"`
}

type Action struct {
	Navigation []Navigation `json:"navigations"`
}

type RenderAction struct {
	Action Action `json:"action"`
}

type ActionResponse struct {
	RenderAction RenderAction `json:"render_actions"`
}

type AuthorizationEventObject struct {
	UserOAuthToken string `json:"userOAuthToken,omitempty"`
}

type Chat struct {
	Type  string     `json:"type,omitempty"`
	User  chat.User  `json:"user,omitempty"`
	Space chat.Space `json:"space,omitempty"`
}

type ChatRequest struct {
	CommonEventObject        chat.CommonEventObject   `json:"commonEventObject,omitempty"`
	AuthorizationEventObject AuthorizationEventObject `json:"authorizationEventObject,omitempty"`
	Chat                     Chat                     `json:"chat,omitempty"`
}

func extractFormInput(event chat.CommonEventObject) types.FormInput {
	getValue := func(key string) string {
		if inputs, exists := event.FormInputs[key]; exists && len(inputs.StringInputs.Value) > 0 {
			return inputs.StringInputs.Value[0]
		}
		return ""
	}

	return types.FormInput{
		Source: getValue("source"),
		Target: getValue("target"),
		Text:   getValue("text"),
		Result: getValue("result"),
	}
}

func validateFormInput(input types.FormInput) error {
	var err error
	if input.Target == "" || input.Text == "" {
		err = fmt.Errorf("please fill target language and text")
	}
	return err
}

func handleSubmitForm(event ChatRequest, configKey string) RenderAction {
	formInput := extractFormInput(event.CommonEventObject)
	var errorMessage string
	var translatedText, source string
	err := validateFormInput(formInput)
	if err == nil {
		var err error
		translatedText, source, err = translators.TranslateText(formInput.Target, formInput.Text, formInput.Source)
		if err != nil {
			slog.Error("Translation error" + err.Error())
			errorMessage = err.Error()
		} else {
			if configJson, err := json.Marshal(formInput); err == nil {
				utils.SetCache(configKey, string(configJson))
			}
		}
	} else {
		errorMessage = err.Error()
	}
	slog.Info("Translation result", formInput.Target, translatedText, formInput.Source, formInput.Text)
	formInput.Result = translatedText
	return RenderAction{
		Action: Action{
			Navigation: []Navigation{{
				UpdateCard: cards.TranslateForm(formInput, source, errorMessage).Card,
			}},
		},
	}
}

func handleInitialLoad(configKey string) RenderAction {
	var formInput types.FormInput
	if lastInputJson, err := utils.GetCache(configKey); err == nil && lastInputJson != "" {
		if err := json.Unmarshal([]byte(lastInputJson), &formInput); err != nil {
			slog.Error("Error unmarshaling cached input: " + err.Error())
		}
	}

	return RenderAction{
		Action: Action{
			Navigation: []Navigation{{
				PushCard: cards.TranslateForm(formInput, "", "").Card,
			}},
		},
	}
}

// HomeHandler handles the HTTP request and returns a JSON response.
// It decodes the request body into a ChatRequest struct, validates the form input,
// translates the text, and returns the appropriate response
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var event ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		slog.Error("Error decoding request: " + err.Error())
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	configKey := fmt.Sprintf("home_%s", event.Chat.User.Name)
	var result RenderAction

	if event.Chat.Type == "SUBMIT_FORM" {
		result = handleSubmitForm(event, configKey)
	} else {
		result = handleInitialLoad(configKey)
	}

	w.Header().Set("Content-Type", "application/json")
	var response interface{} = result
	if event.Chat.Type == "SUBMIT_FORM" {
		response = ActionResponse{RenderAction: result}
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("Error encoding response: " + err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

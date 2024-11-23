package handlers

import (
	"encoding/json"
	"fmt"
	"google.golang.org/api/chat/v1"
	"log"
	"net/http"
	"yaskur.com/chat-translator/cards"
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

func getFormInput(event chat.CommonEventObject) types.FormInput {
	getValue := func(key string) string {
		if inputs, exists := event.FormInputs[key]; exists && len(inputs.StringInputs.Value) > 0 {
			return inputs.StringInputs.Value[0]
		}
		return "" // Default value
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

// HomeHandler handles the HTTP request and returns a JSON response.
// It decodes the request body into a ChatRequest struct, validates the form input,
// translates the text, and returns the appropriate response
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a ChatRequest struct
	var event ChatRequest
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf(event.Chat.Type)
	log.Printf("%#v", event)
	var formInput types.FormInput
	var result RenderAction
	var errorMessage string
	var translatedText string
	var source string
	configKey := "home_" + event.Chat.User.Name
	if event.Chat.Type == "SUBMIT_FORM" {
		formInput = getFormInput(event.CommonEventObject)
		err := validateFormInput(formInput)
		if err != nil {
			log.Printf("invalid validation: %v", err)
			errorMessage = fmt.Sprint(err)
		} else {
			translatedText, source, err = utils.TranslateText(formInput.Target, formInput.Text, formInput.Source)
			if err != nil {
				log.Printf("error translate: %v", err)
				errorMessage = fmt.Sprint(err)
			} else {
				configJson, _ := json.Marshal(formInput)
				utils.SetCache(configKey, string(configJson))
			}
		}
		formInput.Result = translatedText
		result = RenderAction{Action: Action{
			Navigation: []Navigation{{
				UpdateCard: cards.TranslateForm(formInput, source, errorMessage).Card,
			}},
		}}
	} else {
		lastInputJson, _ := utils.GetCache(configKey)
		if lastInputJson != "" {
			err := json.Unmarshal([]byte(lastInputJson), &formInput)
			if err != nil {
				panic(err)
			}
		}
		result = RenderAction{Action: Action{
			Navigation: []Navigation{{
				PushCard: cards.TranslateForm(formInput, "", errorMessage).Card,
			}},
		}}
	}
	w.Header().Set("Content-Type", "application/json")
	if event.Chat.Type == "SUBMIT_FORM" {
		err = json.NewEncoder(w).Encode(ActionResponse{RenderAction: result})
	} else {
		err = json.NewEncoder(w).Encode(result)
	}

	if err != nil {
		log.Fatal(err)
	}
	return
}

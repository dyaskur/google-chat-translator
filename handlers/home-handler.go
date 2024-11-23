package handlers

import (
	"encoding/json"
	"fmt"
	"google.golang.org/api/chat/v1"
	"log"
	"net/http"
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

type FormInput struct {
	Source string `json:"source,omitempty"`
	Target string `json:"target,omitempty"`
	Text   string `json:"text,omitempty"`
	Result string `json:"result,omitempty"`
}

func TranslateForm(formInput FormInput, detectedLanguage string, error string) chat.CardWithId {

	card := chat.CardWithId{
		Card: &chat.GoogleAppsCardV1Card{
			Header: &chat.GoogleAppsCardV1CardHeader{
				Title:    "Translate",
				Subtitle: error,
			},
			Sections: SelectionWidgets(formInput, detectedLanguage),
		},
	}
	return card
}

func SelectionWidgets(formInput FormInput, detectedSource string) []*chat.GoogleAppsCardV1Section {
	log.Printf(formInput.Source + " source")
	var detectedLanguage string
	if detectedSource != "" {
		detectedLanguage = " - " + utils.GroupByCode()[detectedSource].Language
	}
	sourceItems := []*chat.GoogleAppsCardV1SelectionItem{
		{
			Text:     "Detected Language" + detectedLanguage,
			Value:    "auto",
			Selected: formInput.Source == "auto" || formInput.Source == "",
		},
	}
	var targetItems []*chat.GoogleAppsCardV1SelectionItem
	for _, v := range utils.LanguagesData {
		sourceItems = append(sourceItems, &chat.GoogleAppsCardV1SelectionItem{
			Text:     v.Language,
			Value:    v.Code,
			Selected: formInput.Source == v.Code,
		})
		targetItems = append(targetItems, &chat.GoogleAppsCardV1SelectionItem{
			Text:     v.Language,
			Value:    v.Code,
			Selected: formInput.Target == v.Code,
		})
	}

	return []*chat.GoogleAppsCardV1Section{
		{Widgets: []*chat.GoogleAppsCardV1Widget{
			{
				Columns: &chat.GoogleAppsCardV1Columns{
					ColumnItems: []*chat.GoogleAppsCardV1Column{
						{
							Widgets: []*chat.GoogleAppsCardV1Widgets{{
								SelectionInput: &chat.GoogleAppsCardV1SelectionInput{
									Type:  "DROPDOWN",
									Label: "Source",
									Name:  "source",
									Items: sourceItems,
								},
							},
							},
						},
						{
							Widgets: []*chat.GoogleAppsCardV1Widgets{{
								SelectionInput: &chat.GoogleAppsCardV1SelectionInput{
									Type:  "DROPDOWN",
									Label: "Target",
									Name:  "target",
									Items: targetItems,
									OnChangeAction: &chat.GoogleAppsCardV1Action{
										Function: "translate",
									},
								},
							},
							},
						},
					},
				},
			},
			{TextInput: &chat.GoogleAppsCardV1TextInput{
				Label:    "Text input",
				Type:     "MULTIPLE_LINE",
				Name:     "text",
				HintText: "Enter the text to translate",
				Value:    formInput.Text,
				OnChangeAction: &chat.GoogleAppsCardV1Action{
					Function: "translate",
				},
			}},
			{
				DecoratedText: &chat.GoogleAppsCardV1DecoratedText{
					Text: "",
					Button: &chat.GoogleAppsCardV1Button{
						Text: "Translate",
						OnClick: &chat.GoogleAppsCardV1OnClick{
							Action: &chat.GoogleAppsCardV1Action{
								Function: "translate",
							},
						},
					},
				},
			},
			{TextInput: &chat.GoogleAppsCardV1TextInput{
				Label: "Result",
				Type:  "MULTIPLE_LINE",
				Name:  "result",
				Value: formInput.Result,
			}},
		}},
	}
}

func getFormInput(event chat.CommonEventObject) FormInput {
	getValue := func(key string) string {
		if inputs, exists := event.FormInputs[key]; exists && len(inputs.StringInputs.Value) > 0 {
			return inputs.StringInputs.Value[0]
		}
		return "" // Default value
	}

	return FormInput{
		Source: getValue("source"),
		Target: getValue("target"),
		Text:   getValue("text"),
		Result: getValue("result"),
	}
}

func validateFormInput(input FormInput) error {
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
	var formInput FormInput
	var result RenderAction
	var errorMessage string
	var translatedText string
	var source string
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
			}
		}
		formInput.Result = translatedText
		result = RenderAction{Action: Action{
			Navigation: []Navigation{{
				UpdateCard: TranslateForm(formInput, source, errorMessage).Card,
			}},
		}}
	} else {
		result = RenderAction{Action: Action{
			Navigation: []Navigation{{
				PushCard: TranslateForm(formInput, "", errorMessage).Card,
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

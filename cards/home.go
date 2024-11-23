package cards

import (
	"google.golang.org/api/chat/v1"
	"yaskur.com/chat-translator/types"
	"yaskur.com/chat-translator/utils"
)

func TranslateForm(formInput types.FormInput, detectedLanguage string, error string) chat.CardWithId {

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

func SelectionWidgets(formInput types.FormInput, detectedSource string) []*chat.GoogleAppsCardV1Section {
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

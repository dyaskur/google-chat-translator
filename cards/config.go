package cards

import (
	"google.golang.org/api/chat/v1"
	"strconv"
	"yaskur.com/chat-translator/types"
)

// ConfigForm generates a card based on the provided configuration.
func ConfigForm(config types.Config) *chat.GoogleAppsCardV1Card {
	return &chat.GoogleAppsCardV1Card{
		Header: &chat.GoogleAppsCardV1CardHeader{
			Title: "Abang Translator Config",
		},
		Sections: []*chat.GoogleAppsCardV1Section{
			{
				Widgets: []*chat.GoogleAppsCardV1Widget{
					{
						DecoratedText: &chat.GoogleAppsCardV1DecoratedText{
							SwitchControl: &chat.GoogleAppsCardV1SwitchControl{
								Selected:    config.ShowOriginalText,
								Name:        "show_original_text",
								ControlType: "SWITCH",
								OnChangeAction: &chat.GoogleAppsCardV1Action{
									Function: "setShowOriginalText",
									Parameters: []*chat.GoogleAppsCardV1ActionParameter{
										{
											Key:   "show_original_text",
											Value: strconv.FormatBool(!config.ShowOriginalText),
										},
									},
								},
							},
							StartIcon: &chat.GoogleAppsCardV1Icon{
								MaterialIcon: &chat.GoogleAppsCardV1MaterialIcon{
									Name: "source_notes",
								},
							},
							Text:        "Show Original Text: ",
							BottomLabel: "If enabled, the original text will be shown in the response. If disabled, only the translated text will be shown.",
						},
					},
				},
			},
		},
	}
}

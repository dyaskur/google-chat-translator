package cards

import "google.golang.org/api/chat/v1"

// ConfigForm generates a card based on the provided configuration.
func ConfigForm(showOriginalText bool) *chat.GoogleAppsCardV1Card {
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
								Selected:    showOriginalText,
								Name:        "show_original_text",
								ControlType: "SWITCH",
							},
							StartIcon: &chat.GoogleAppsCardV1Icon{
								KnownIcon: "CHECK",
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

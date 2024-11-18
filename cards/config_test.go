package cards

import (
	"testing"
	"yaskur.com/chat-translator/types"
)

func TestConfigFormCreatesCardWithCorrectTitleAndSwitchState(t *testing.T) {
	config := types.Config{ShowOriginalText: true}
	card := ConfigForm(config)

	if card.Header.Title != "Abang Translator Config" {
		t.Errorf("Expected card title to be 'Abang Translator Config', got '%s'", card.Header.Title)
	}

	if len(card.Sections) == 0 || len(card.Sections[0].Widgets) == 0 {
		t.Fatal("Expected card to have at least one section and one widget")
	}

	widget := card.Sections[0].Widgets[0]
	if widget.DecoratedText.SwitchControl.Selected != config.ShowOriginalText {
		t.Errorf("Expected switch control selected state to be '%v', got '%v'", config.ShowOriginalText, widget.DecoratedText.SwitchControl.Selected)
	}
}

func TestConfigFormHandlesFalseShowOriginalText(t *testing.T) {
	config := types.Config{ShowOriginalText: false}
	card := ConfigForm(config)

	if len(card.Sections) == 0 || len(card.Sections[0].Widgets) == 0 {
		t.Fatal("Expected card to have at least one section and one widget")
	}

	widget := card.Sections[0].Widgets[0]
	if widget.DecoratedText.SwitchControl.Selected != config.ShowOriginalText {
		t.Errorf("Expected switch control selected state to be '%v', got '%v'", config.ShowOriginalText, widget.DecoratedText.SwitchControl.Selected)
	}
}

func TestConfigFormHandlesTrueShowOriginalText(t *testing.T) {
	config := types.Config{ShowOriginalText: true}
	card := ConfigForm(config)

	if len(card.Sections) == 0 || len(card.Sections[0].Widgets) == 0 {
		t.Fatal("Expected card to have at least one section and one widget")
	}

	widget := card.Sections[0].Widgets[0]
	if widget.DecoratedText.SwitchControl.Selected != config.ShowOriginalText {
		t.Errorf("Expected switch control selected state to be '%v', got '%v'", config.ShowOriginalText, widget.DecoratedText.SwitchControl.Selected)
	}
}

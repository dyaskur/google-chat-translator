package types

// Config represents the configuration for the Abang Translator.
type Config struct {
	ShowOriginalText bool `json:"show_original_text,omitempty"`
}

type FormInput struct {
	Source string `json:"source,omitempty"`
	Target string `json:"target,omitempty"`
	Text   string `json:"text,omitempty"`
	Result string `json:"result,omitempty"`
}

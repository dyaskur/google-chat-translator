package utils

import (
	"cloud.google.com/go/translate"
	"context"
	"fmt"
	"golang.org/x/text/language"
)

func TranslateText(targetLanguage string, text string, sourceLanguage string) (string, string, error) {
	// text := "The Go Gopher is cute"
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", "", fmt.Errorf("language target.Parse: %w", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", "", err
	}
	defer func(client *translate.Client) {
		err := client.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(client)
	opts := &translate.Options{}
	opts.Format = translate.Text
	if sourceLanguage != "" && sourceLanguage != "auto" {
		source, err := language.Parse(sourceLanguage)
		if err != nil {
			return "", "", fmt.Errorf("language source.Parse: %w", err)
		}
		opts.Source = source
	}
	resp, err := client.Translate(ctx, []string{text}, lang, opts)
	if err != nil {
		return "", "", fmt.Errorf("translate: %w", err)
	}
	if len(resp) == 0 {
		return "", "", fmt.Errorf("translate returned empty response to text: %s", text)
	}
	return resp[0].Text, resp[0].Source.String(), nil
}

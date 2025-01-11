package translators

import "log"

func TranslateText(targetLanguage string, text string, sourceLanguage string) (string, string, error) {
	result, sourceLanguage, err := straicoTranslate(targetLanguage, text, sourceLanguage)
	if err == nil {
		return result, sourceLanguage, nil
	}
	log.Printf("error %s; trying google translator", err.Error())
	// default to use google translator
	return googleTranslate(targetLanguage, text, sourceLanguage)
}

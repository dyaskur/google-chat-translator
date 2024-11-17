package translator

import (
	"google.golang.org/api/chat/v1"
	"log"
	"strconv"
	"yaskur.com/chat-translator/utils"
)

func commandHandler(event chat.DeprecatedEvent) chat.Message {
	message := event.Message
	log.Printf("commandID: %s", strconv.FormatInt(message.SlashCommand.CommandId, 10))
	targetLanguage := utils.GetById(int16(message.SlashCommand.CommandId))
	log.Printf("targetLanguage: %s", targetLanguage.Code)
	translatedText, source, err := translateText(targetLanguage.Code, message.ArgumentText, "")
	if err != nil {
		log.Fatal(err)
	}
	sourceLanguage := utils.GetByCode(source)

	user := event.User
	response := "_" + user.DisplayName + " said: (translated to " + targetLanguage.Language + ")_\n" + translatedText
	response = response + "\nTranslated from " + sourceLanguage.Language + ", original message:\n" + message.ArgumentText
	reply := chat.Message{
		ActionResponse: &chat.ActionResponse{
			Type: "NEW_MESSAGE",
		},
		Text: response,
	}
	return reply
}

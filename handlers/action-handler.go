package handlers

import (
	"encoding/json"
	"fmt"
	"google.golang.org/api/chat/v1"
	"log"
	"strconv"
	"yaskur.com/chat-translator/cards"
	"yaskur.com/chat-translator/types"
	"yaskur.com/chat-translator/utils"
)

func ActionHandler(event types.ChatEvent) chat.Message {
	action := event.Common.InvokedFunction
	reply := chat.Message{}
	if action == "setShowOriginalText" {
		var input = event.Common.Parameters["show_original_text"]
		fmt.Println(input)
		showOriginalText, _ := strconv.ParseBool(input)
		config := types.Config{
			ShowOriginalText: showOriginalText,
		}
		configJson, _ := json.Marshal(config)
		utils.SetCache(event.Space.Name, string(configJson))
		reply = chat.Message{
			ActionResponse: &chat.ActionResponse{
				Type: "DIALOG",
				DialogAction: &chat.DialogAction{
					Dialog: &chat.Dialog{
						Body: cards.ConfigForm(config),
					},
				},
			},
		}
	} else if event.Action.ActionMethodName == "translate" {
		configKey := "home_" + event.User.Name
		var formInput types.FormInput
		var errorMessage string
		var translatedText string
		var source string
		formInput = getFormInput(*event.Common)
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
		reply = chat.Message{
			ActionResponse: &chat.ActionResponse{
				Type: "DIALOG",
				DialogAction: &chat.DialogAction{
					Dialog: &chat.Dialog{
						Body: cards.TranslateForm(formInput, source, errorMessage).Card,
					},
				},
			},
		}

	}
	return reply
}

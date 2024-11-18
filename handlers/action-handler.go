package handlers

import (
	"encoding/json"
	"fmt"
	"google.golang.org/api/chat/v1"
	"strconv"
	"yaskur.com/chat-translator/cards"
	"yaskur.com/chat-translator/types"
	"yaskur.com/chat-translator/utils"
)

func ActionHandler(event chat.DeprecatedEvent) chat.Message {
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
	}
	return reply
}

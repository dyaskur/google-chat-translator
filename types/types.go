package types

import (
	"google.golang.org/api/chat/v1"
)

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

type User struct {
	// DisplayName: Output only. The user's display name.
	DisplayName string `json:"displayName,omitempty"`
	// DomainId: Unique identifier of the user's Google Workspace domain.
	DomainId string `json:"domainId,omitempty"`
	// IsAnonymous: Output only. When `true`, the user is deleted or their profile
	// is not visible.
	IsAnonymous bool `json:"isAnonymous,omitempty"`
	// Name: Resource name for a Google Chat user. Format: `users/{user}`.
	// `users/app` can be used as an alias for the calling app bot user. For human
	// users, `{user}` is the same user identifier as: - the `id` for the Person
	// (https://developers.google.com/people/api/rest/v1/people) in the People API.
	// For example, `users/123456789` in Chat API represents the same person as the
	// `123456789` Person profile ID in People API. - the `id` for a user
	// (https://developers.google.com/admin-sdk/directory/reference/rest/v1/users)
	// in the Admin SDK Directory API. - the user's email address can be used as an
	// alias for `{user}` in API requests. For example, if the People API Person
	// profile ID for `user@example.com` is `123456789`, you can use
	// `users/user@example.com` as an alias to reference `users/123456789`. Only
	// the canonical resource name (for example `users/123456789`) will be returned
	// from the API.
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	// Type: User type.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - Default value for the enum. DO NOT USE.
	//   "HUMAN" - Human user.
	//   "BOT" - Chat app user.
	Type string `json:"type,omitempty"`
	// ForceSendFields is a list of field names (e.g. "DisplayName") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "DisplayName") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

type ChatEvent struct {
	// Action: For `CARD_CLICKED` interaction events, the form action data
	// associated when a user clicks a card or dialog. To learn more, see Read form
	// data input by users on cards
	// (https://developers.google.com/workspace/chat/read-form-data).
	Action *chat.FormAction `json:"action,omitempty"`
	// Common: Represents information about the user's client, such as locale, host
	// app, and platform. For Chat apps, `CommonEventObject` includes information
	// submitted by users interacting with dialogs
	// (https://developers.google.com/workspace/chat/dialogs), like data entered on
	// a card.
	Common *chat.CommonEventObject `json:"common,omitempty"`
	// ConfigCompleteRedirectUrl: For `MESSAGE` interaction events, the URL that
	// users must be redirected to after they complete an authorization or
	// configuration flow outside of Google Chat. For more information, see Connect
	// a Chat app with other services and tools
	// (https://developers.google.com/workspace/chat/connect-web-services-tools).
	ConfigCompleteRedirectUrl string `json:"configCompleteRedirectUrl,omitempty"`
	// DialogEventType: The type of dialog
	// (https://developers.google.com/workspace/chat/dialogs) interaction event
	// received.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - Default value. Unspecified.
	//   "REQUEST_DIALOG" - A user opens a dialog.
	//   "SUBMIT_DIALOG" - A user clicks an interactive element of a dialog. For
	// example, a user fills out information in a dialog and clicks a button to
	// submit the information.
	//   "CANCEL_DIALOG" - A user closes a dialog without submitting information.
	// The Chat app only receives this interaction event when users click the close
	// icon in the top right corner of the dialog. When the user closes the dialog
	// by other means (such as refreshing the browser, clicking outside the dialog
	// box, or pressing the escape key), no event is sent. .
	DialogEventType string `json:"dialogEventType,omitempty"`
	// EventTime: The timestamp indicating when the interaction event occurred.
	EventTime string `json:"eventTime,omitempty"`
	// IsDialogEvent: For `CARD_CLICKED` and `MESSAGE` interaction events, whether
	// the user is interacting with or about to interact with a dialog
	// (https://developers.google.com/workspace/chat/dialogs).
	IsDialogEvent bool `json:"isDialogEvent,omitempty"`
	// Message: For `ADDED_TO_SPACE`, `CARD_CLICKED`, and `MESSAGE` interaction
	// events, the message that triggered the interaction event, if applicable.
	Message *chat.Message `json:"message,omitempty"`
	// Space: The space in which the user interacted with the Chat app.
	Space *chat.Space `json:"space,omitempty"`
	// ThreadKey: The Chat app-defined key for the thread related to the
	// interaction event. See `spaces.messages.thread.threadKey`
	// (/chat/api/reference/rest/v1/spaces.messages#Thread.FIELDS.thread_key) for
	// more information.
	ThreadKey string `json:"threadKey,omitempty"`
	// Token: A secret value that legacy Chat apps can use to verify if a request
	// is from Google. Google randomly generates the token, and its value remains
	// static. You can obtain, revoke, or regenerate the token from the Chat API
	// configuration page
	// (https://console.cloud.google.com/apis/api/chat.googleapis.com/hangouts-chat)
	// in the Google Cloud Console. Modern Chat apps don't use this field. It is
	// absent from API responses and the Chat API configuration page
	// (https://console.cloud.google.com/apis/api/chat.googleapis.com/hangouts-chat).
	Token string `json:"token,omitempty"`
	// Type: The type (/workspace/chat/api/reference/rest/v1/EventType) of user
	// interaction with the Chat app, such as `MESSAGE` or `ADDED_TO_SPACE`.
	//
	// Possible values:
	//   "UNSPECIFIED" - Default value for the enum. DO NOT USE.
	//   "MESSAGE" - A user sends the Chat app a message, or invokes the Chat app
	// in a space, such as any of the following examples: * Any message in a direct
	// message (DM) space with the Chat app. * A message in a multi-person space
	// where a person @mentions the Chat app, or uses one of its slash commands. *
	// If you've configured link previews for your Chat app, a user posts a message
	// that contains a link that matches the configured URL pattern.
	//   "ADDED_TO_SPACE" - A user adds the Chat app to a space, or a Google
	// Workspace administrator installs the Chat app in direct message spaces for
	// users in their organization. Chat apps typically respond to this interaction
	// event by posting a welcome message in the space. When administrators install
	// Chat apps, the `space.adminInstalled` field is set to `true` and users can't
	// uninstall them. To learn about Chat apps installed by administrators, see
	// Google Workspace Admin Help's documentation, [Install Marketplace apps in
	// your domain](https://support.google.com/a/answer/172482).
	//   "REMOVED_FROM_SPACE" - A user removes the Chat app from a space, or a
	// Google Workspace administrator uninstalls the Chat app for a user in their
	// organization. Chat apps can't respond with messages to this event, because
	// they have already been removed. When administrators uninstall Chat apps, the
	// `space.adminInstalled` field is set to `false`. If a user installed the Chat
	// app before the administrator, the Chat app remains installed for the user
	// and the Chat app doesn't receive a `REMOVED_FROM_SPACE` interaction event.
	//   "CARD_CLICKED" - A user clicks an interactive element of a card or dialog
	// from a Chat app, such as a button. To receive an interaction event, the
	// button must trigger another interaction with the Chat app. For example, a
	// Chat app doesn't receive a `CARD_CLICKED` interaction event if a user clicks
	// a button that opens a link to a website, but receives interaction events in
	// the following examples: * The user clicks a `Send feedback` button on a
	// card, which opens a dialog for the user to input information. * The user
	// clicks a `Submit` button after inputting information into a card or dialog.
	// If a user clicks a button to open, submit, or cancel a dialog, the
	// `CARD_CLICKED` interaction event's `isDialogEvent` field is set to `true`
	// and includes a
	// [`DialogEventType`](https://developers.google.com/workspace/chat/api/referenc
	// e/rest/v1/DialogEventType).
	//   "WIDGET_UPDATED" - A user updates a widget in a card message or dialog.
	Type string `json:"type,omitempty"`
	// User: The user that interacted with the Chat app.
	User *User `json:"user,omitempty"`
	// ForceSendFields is a list of field names (e.g. "Action") to unconditionally
	// include in API requests. By default, fields with empty or default values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "Action") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

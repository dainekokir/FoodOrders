package api

type sendMessageBtnReqBody struct {
	ChatID      int64               `json:"chat_id"`
	Text        string              `json:"text"`
	ReplyMarkup ReplyKeyboardMarkup `json:"reply_markup,omitempty"`
	Entities    []MessageEntity     `json:"entities,omitempty"`
}

type sendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}
type MessageEntity struct {
	// Type - Type of the entity. Currently, can be “mention” (@username), “hashtag” (#hashtag),
	// “cashtag” ($USD), “bot_command” (/start@jobs_bot), “url” (https://telegram.org), “email”
	// (do-not-reply@telegram.org), “phone_number” (+1-212-555-0123), “bold” (bold text), “italic”
	// (italic text), “underline” (underlined text), “strikethrough” (strikethrough text), “spoiler”
	// (spoiler message), “code” (monowidth string), “pre” (monowidth block), “text_link” (for clickable
	// text URLs), “text_mention” (for users without usernames (https://telegram.org/blog/edit#new-mentions))
	Type string `json:"type"`

	// Offset - Offset in UTF-16 code units to the start of the entity
	Offset int `json:"offset"`

	// Length - Length of the entity in UTF-16 code units
	Length int `json:"length"`

	// URL - Optional. For “text_link” only, URL that will be opened after user taps on the text
	URL string `json:"url,omitempty"`

	// User - Optional. For “text_mention” only, the mentioned user
	User *User `json:"user,omitempty"`

	// Language - Optional. For “pre” only, the programming language of the entity text
	Language string `json:"language,omitempty"`
}
type User struct {
	// ID - Unique identifier for this user or bot. This number may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting it. But it has at most 52
	// significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	ID int64 `json:"id"`

	// IsBot - True, if this user is a bot
	IsBot bool `json:"is_bot"`

	// FirstName - User's or bot's first name
	FirstName string `json:"first_name"`

	// LastName - Optional. User's or bot's last name
	LastName string `json:"last_name,omitempty"`

	// Username - Optional. User's or bot's username
	Username string `json:"username,omitempty"`

	// LanguageCode - Optional. IETF language tag (https://en.wikipedia.org/wiki/IETF_language_tag) of the
	// user's language
	LanguageCode string `json:"language_code,omitempty"`

	// IsPremium - Optional. True, if this user is a Telegram Premium user
	IsPremium bool `json:"is_premium,omitempty"`

	// AddedToAttachmentMenu - Optional. True, if this user added the bot to the attachment menu
	AddedToAttachmentMenu bool `json:"added_to_attachment_menu,omitempty"`

	// CanJoinGroups - Optional. True, if the bot can be invited to groups. Returned only in getMe
	// (https://core.telegram.org/bots/api#getme).
	CanJoinGroups bool `json:"can_join_groups,omitempty"`

	// CanReadAllGroupMessages - Optional. True, if privacy mode (https://core.telegram.org/bots#privacy-mode)
	// is disabled for the bot. Returned only in getMe (https://core.telegram.org/bots/api#getme).
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages,omitempty"`

	// SupportsInlineQueries - Optional. True, if the bot supports inline queries. Returned only in getMe
	// (https://core.telegram.org/bots/api#getme).
	SupportsInlineQueries bool `json:"supports_inline_queries,omitempty"`
}
type ReplyKeyboardMarkup struct {
	// Keyboard - Array of button rows, each represented by an Array of KeyboardButton
	// (https://core.telegram.org/bots/api#keyboardbutton) objects
	Keyboard [][]KeyboardButton `json:"keyboard"`

	// ResizeKeyboard - Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make
	// the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom
	// keyboard is always of the same height as the app's standard keyboard.
	ResizeKeyboard bool `json:"resize_keyboard,omitempty"`

	// OneTimeKeyboard - Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard
	// will still be available, but clients will automatically display the usual letter-keyboard in the chat - the
	// user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`

	// InputFieldPlaceholder - Optional. The placeholder to be shown in the input field when the keyboard is
	// active; 1-64 characters
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`

	// Selective - Optional. Use this parameter if you want to show the keyboard to specific users only.
	// Targets: 1) users that are @mentioned in the text of the Message (https://core.telegram.org/bots/api#message)
	// object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	// Example: A user requests to change the bot's language, bot replies to the request with a keyboard to select
	// the new language. Other users in the group don't see the keyboard.
	Selective bool `json:"selective,omitempty"`
}
type KeyboardButton struct {
	// Text - Text of the button. If none of the optional fields are used, it will be sent as a message when the
	// button is pressed
	Text string `json:"text"`

	// RequestContact - Optional. If True, the user's phone number will be sent as a contact when the button is
	// pressed. Available in private chats only.
	RequestContact bool `json:"request_contact,omitempty"`

	// RequestLocation - Optional. If True, the user's current location will be sent when the button is pressed.
	// Available in private chats only.
	RequestLocation bool `json:"request_location,omitempty"`

	// RequestPoll - Optional. If specified, the user will be asked to create a poll and send it to the bot when
	// the button is pressed. Available in private chats only.
	RequestPoll *KeyboardButtonPollType `json:"request_poll,omitempty"`

	// WebApp - Optional. If specified, the described Web App (https://core.telegram.org/bots/webapps) will be
	// launched when the button is pressed. The Web App will be able to send a “web_app_data” service message.
	// Available in private chats only.
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}
type KeyboardButtonPollType struct {
	// Type - Optional. If quiz is passed, the user will be allowed to create only polls in the quiz mode. If
	// regular is passed, only regular polls will be allowed. Otherwise, the user will be allowed to create a poll
	// of any type.
	Type string `json:"type,omitempty"`
}
type WebAppInfo struct {
	// URL - An HTTPS URL of a Web App to be opened with additional data as specified in Initializing Web Apps
	// (https://core.telegram.org/bots/webapps#initializing-web-apps)
	URL string `json:"url"`
}
type InlineKeyboardMarkup struct {
	// InlineKeyboard - Array of button rows, each represented by an Array of InlineKeyboardButton
	// (https://core.telegram.org/bots/api#inlinekeyboardbutton) objects
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}
type InlineKeyboardButton struct {
	// Text - Label text on the button
	Text string `json:"text"`

	// URL - Optional. HTTP or tg:// URL to be opened when the button is pressed. Links tg://user?id=<user_id>
	// can be used to mention a user by their ID without using a username, if this is allowed by their privacy
	// settings.
	URL string `json:"url,omitempty"`

	// CallbackData - Optional. Data to be sent in a callback query
	// (https://core.telegram.org/bots/api#callbackquery) to the bot when button is pressed, 1-64 bytes
	CallbackData string `json:"callback_data,omitempty"`

	// WebApp - Optional. Description of the Web App (https://core.telegram.org/bots/webapps) that will be
	// launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of
	// the user using the method answerWebAppQuery (https://core.telegram.org/bots/api#answerwebappquery). Available
	// only in private chats between a user and the bot.
	WebApp *WebAppInfo `json:"web_app,omitempty"`

	// LoginURL - Optional. An HTTPS URL used to automatically authorize the user. Can be used as a replacement
	// for the Telegram Login Widget (https://core.telegram.org/widgets/login).
	LoginURL *LoginURL `json:"login_url,omitempty"`

	// SwitchInlineQuery - Optional. If set, pressing the button will prompt the user to select one of their
	// chats, open that chat and insert the bot's username and the specified inline query in the input field. May be
	// empty, in which case just the bot's username will be inserted.
	// Note: This offers an easy way for users to start using your bot in inline mode
	// (https://core.telegram.org/bots/inline) when they are currently in a private chat with it. Especially useful
	// when combined with switch_pm… (https://core.telegram.org/bots/api#answerinlinequery) actions - in this case
	// the user will be automatically returned to the chat they switched from, skipping the chat selection screen.
	SwitchInlineQuery string `json:"switch_inline_query,omitempty"`

	// SwitchInlineQueryCurrentChat - Optional. If set, pressing the button will insert the bot's username and
	// the specified inline query in the current chat's input field. May be empty, in which case only the bot's
	// username will be inserted.
	// This offers a quick way for the user to open your bot in inline mode in the same chat - good for selecting
	// something from multiple options.
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`

	// CallbackGame - Optional. Description of the game that will be launched when the user presses the button.
	// NOTE: This type of button must always be the first button in the first row.
	CallbackGame *CallbackGame `json:"callback_game,omitempty"`

	// Pay - Optional. Specify True, to send a Pay button (https://core.telegram.org/bots/api#payments).
	// NOTE: This type of button must always be the first button in the first row and can only be used in invoice
	// messages.
	Pay bool `json:"pay,omitempty"`
}
type CallbackGame struct{}
type LoginURL struct {
	// URL - An HTTP URL to be opened with user authorization data added to the query string when the button is
	// pressed. If the user refuses to provide authorization data, the original URL without information about the
	// user will be opened. The data added is the same as described in Receiving authorization data
	// (https://core.telegram.org/widgets/login#receiving-authorization-data).
	// NOTE: You must always check the hash of the received data to verify the authentication and the integrity of
	// the data as described in Checking authorization
	// (https://core.telegram.org/widgets/login#checking-authorization).
	URL string `json:"url"`

	// ForwardText - Optional. New text of the button in forwarded messages.
	ForwardText string `json:"forward_text,omitempty"`

	// BotUsername - Optional. Username of a bot, which will be used for user authorization. See Setting up a
	// bot (https://core.telegram.org/widgets/login#setting-up-a-bot) for more details. If not specified, the
	// current bot's username will be assumed. The URL's domain must be the same as the domain linked with the bot.
	// See Linking your domain to the bot (https://core.telegram.org/widgets/login#linking-your-domain-to-the-bot)
	// for more details.
	BotUsername string `json:"bot_username,omitempty"`

	// RequestWriteAccess - Optional. Pass True to request the permission for your bot to send messages to the
	// user.
	RequestWriteAccess bool `json:"request_write_access,omitempty"`
}

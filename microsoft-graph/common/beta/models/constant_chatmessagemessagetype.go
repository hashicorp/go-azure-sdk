package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageMessageType string

const (
	ChatMessageMessageTypechatEvent          ChatMessageMessageType = "ChatEvent"
	ChatMessageMessageTypemessage            ChatMessageMessageType = "Message"
	ChatMessageMessageTypesystemEventMessage ChatMessageMessageType = "SystemEventMessage"
	ChatMessageMessageTypetyping             ChatMessageMessageType = "Typing"
)

func PossibleValuesForChatMessageMessageType() []string {
	return []string{
		string(ChatMessageMessageTypechatEvent),
		string(ChatMessageMessageTypemessage),
		string(ChatMessageMessageTypesystemEventMessage),
		string(ChatMessageMessageTypetyping),
	}
}

func parseChatMessageMessageType(input string) (*ChatMessageMessageType, error) {
	vals := map[string]ChatMessageMessageType{
		"chatevent":          ChatMessageMessageTypechatEvent,
		"message":            ChatMessageMessageTypemessage,
		"systemeventmessage": ChatMessageMessageTypesystemEventMessage,
		"typing":             ChatMessageMessageTypetyping,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessageMessageType(input)
	return &out, nil
}

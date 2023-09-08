package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageInfoMessageType string

const (
	ChatMessageInfoMessageTypechatEvent          ChatMessageInfoMessageType = "ChatEvent"
	ChatMessageInfoMessageTypemessage            ChatMessageInfoMessageType = "Message"
	ChatMessageInfoMessageTypesystemEventMessage ChatMessageInfoMessageType = "SystemEventMessage"
	ChatMessageInfoMessageTypetyping             ChatMessageInfoMessageType = "Typing"
)

func PossibleValuesForChatMessageInfoMessageType() []string {
	return []string{
		string(ChatMessageInfoMessageTypechatEvent),
		string(ChatMessageInfoMessageTypemessage),
		string(ChatMessageInfoMessageTypesystemEventMessage),
		string(ChatMessageInfoMessageTypetyping),
	}
}

func parseChatMessageInfoMessageType(input string) (*ChatMessageInfoMessageType, error) {
	vals := map[string]ChatMessageInfoMessageType{
		"chatevent":          ChatMessageInfoMessageTypechatEvent,
		"message":            ChatMessageInfoMessageTypemessage,
		"systemeventmessage": ChatMessageInfoMessageTypesystemEventMessage,
		"typing":             ChatMessageInfoMessageTypetyping,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessageInfoMessageType(input)
	return &out, nil
}

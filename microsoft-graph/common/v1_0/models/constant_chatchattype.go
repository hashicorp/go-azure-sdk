package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatChatType string

const (
	ChatChatTypegroup    ChatChatType = "Group"
	ChatChatTypemeeting  ChatChatType = "Meeting"
	ChatChatTypeoneOnOne ChatChatType = "OneOnOne"
)

func PossibleValuesForChatChatType() []string {
	return []string{
		string(ChatChatTypegroup),
		string(ChatChatTypemeeting),
		string(ChatChatTypeoneOnOne),
	}
}

func parseChatChatType(input string) (*ChatChatType, error) {
	vals := map[string]ChatChatType{
		"group":    ChatChatTypegroup,
		"meeting":  ChatChatTypemeeting,
		"oneonone": ChatChatTypeoneOnOne,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatChatType(input)
	return &out, nil
}

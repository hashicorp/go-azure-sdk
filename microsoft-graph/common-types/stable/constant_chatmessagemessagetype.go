package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageMessageType string

const (
	ChatMessageMessageType_ChatEvent          ChatMessageMessageType = "chatEvent"
	ChatMessageMessageType_Message            ChatMessageMessageType = "message"
	ChatMessageMessageType_SystemEventMessage ChatMessageMessageType = "systemEventMessage"
	ChatMessageMessageType_Typing             ChatMessageMessageType = "typing"
)

func PossibleValuesForChatMessageMessageType() []string {
	return []string{
		string(ChatMessageMessageType_ChatEvent),
		string(ChatMessageMessageType_Message),
		string(ChatMessageMessageType_SystemEventMessage),
		string(ChatMessageMessageType_Typing),
	}
}

func (s *ChatMessageMessageType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatMessageMessageType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatMessageMessageType(input string) (*ChatMessageMessageType, error) {
	vals := map[string]ChatMessageMessageType{
		"chatevent":          ChatMessageMessageType_ChatEvent,
		"message":            ChatMessageMessageType_Message,
		"systemeventmessage": ChatMessageMessageType_SystemEventMessage,
		"typing":             ChatMessageMessageType_Typing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessageMessageType(input)
	return &out, nil
}

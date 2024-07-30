package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageInfoMessageType string

const (
	ChatMessageInfoMessageType_ChatEvent          ChatMessageInfoMessageType = "chatEvent"
	ChatMessageInfoMessageType_Message            ChatMessageInfoMessageType = "message"
	ChatMessageInfoMessageType_SystemEventMessage ChatMessageInfoMessageType = "systemEventMessage"
	ChatMessageInfoMessageType_Typing             ChatMessageInfoMessageType = "typing"
)

func PossibleValuesForChatMessageInfoMessageType() []string {
	return []string{
		string(ChatMessageInfoMessageType_ChatEvent),
		string(ChatMessageInfoMessageType_Message),
		string(ChatMessageInfoMessageType_SystemEventMessage),
		string(ChatMessageInfoMessageType_Typing),
	}
}

func (s *ChatMessageInfoMessageType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatMessageInfoMessageType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatMessageInfoMessageType(input string) (*ChatMessageInfoMessageType, error) {
	vals := map[string]ChatMessageInfoMessageType{
		"chatevent":          ChatMessageInfoMessageType_ChatEvent,
		"message":            ChatMessageInfoMessageType_Message,
		"systemeventmessage": ChatMessageInfoMessageType_SystemEventMessage,
		"typing":             ChatMessageInfoMessageType_Typing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessageInfoMessageType(input)
	return &out, nil
}

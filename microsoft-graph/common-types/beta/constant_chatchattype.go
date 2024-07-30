package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatChatType string

const (
	ChatChatType_Group    ChatChatType = "group"
	ChatChatType_Meeting  ChatChatType = "meeting"
	ChatChatType_OneOnOne ChatChatType = "oneOnOne"
)

func PossibleValuesForChatChatType() []string {
	return []string{
		string(ChatChatType_Group),
		string(ChatChatType_Meeting),
		string(ChatChatType_OneOnOne),
	}
}

func (s *ChatChatType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatChatType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatChatType(input string) (*ChatChatType, error) {
	vals := map[string]ChatChatType{
		"group":    ChatChatType_Group,
		"meeting":  ChatChatType_Meeting,
		"oneonone": ChatChatType_OneOnOne,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatChatType(input)
	return &out, nil
}

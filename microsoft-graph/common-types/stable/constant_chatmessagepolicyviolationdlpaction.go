package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessagePolicyViolationDlpAction string

const (
	ChatMessagePolicyViolationDlpAction_BlockAccess         ChatMessagePolicyViolationDlpAction = "blockAccess"
	ChatMessagePolicyViolationDlpAction_BlockAccessExternal ChatMessagePolicyViolationDlpAction = "blockAccessExternal"
	ChatMessagePolicyViolationDlpAction_None                ChatMessagePolicyViolationDlpAction = "none"
	ChatMessagePolicyViolationDlpAction_NotifySender        ChatMessagePolicyViolationDlpAction = "notifySender"
)

func PossibleValuesForChatMessagePolicyViolationDlpAction() []string {
	return []string{
		string(ChatMessagePolicyViolationDlpAction_BlockAccess),
		string(ChatMessagePolicyViolationDlpAction_BlockAccessExternal),
		string(ChatMessagePolicyViolationDlpAction_None),
		string(ChatMessagePolicyViolationDlpAction_NotifySender),
	}
}

func (s *ChatMessagePolicyViolationDlpAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatMessagePolicyViolationDlpAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatMessagePolicyViolationDlpAction(input string) (*ChatMessagePolicyViolationDlpAction, error) {
	vals := map[string]ChatMessagePolicyViolationDlpAction{
		"blockaccess":         ChatMessagePolicyViolationDlpAction_BlockAccess,
		"blockaccessexternal": ChatMessagePolicyViolationDlpAction_BlockAccessExternal,
		"none":                ChatMessagePolicyViolationDlpAction_None,
		"notifysender":        ChatMessagePolicyViolationDlpAction_NotifySender,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessagePolicyViolationDlpAction(input)
	return &out, nil
}

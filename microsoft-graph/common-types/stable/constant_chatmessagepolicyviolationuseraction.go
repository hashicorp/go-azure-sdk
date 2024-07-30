package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessagePolicyViolationUserAction string

const (
	ChatMessagePolicyViolationUserAction_None                ChatMessagePolicyViolationUserAction = "none"
	ChatMessagePolicyViolationUserAction_Override            ChatMessagePolicyViolationUserAction = "override"
	ChatMessagePolicyViolationUserAction_ReportFalsePositive ChatMessagePolicyViolationUserAction = "reportFalsePositive"
)

func PossibleValuesForChatMessagePolicyViolationUserAction() []string {
	return []string{
		string(ChatMessagePolicyViolationUserAction_None),
		string(ChatMessagePolicyViolationUserAction_Override),
		string(ChatMessagePolicyViolationUserAction_ReportFalsePositive),
	}
}

func (s *ChatMessagePolicyViolationUserAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatMessagePolicyViolationUserAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatMessagePolicyViolationUserAction(input string) (*ChatMessagePolicyViolationUserAction, error) {
	vals := map[string]ChatMessagePolicyViolationUserAction{
		"none":                ChatMessagePolicyViolationUserAction_None,
		"override":            ChatMessagePolicyViolationUserAction_Override,
		"reportfalsepositive": ChatMessagePolicyViolationUserAction_ReportFalsePositive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessagePolicyViolationUserAction(input)
	return &out, nil
}

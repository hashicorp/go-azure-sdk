package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessagePolicyViolationVerdictDetails string

const (
	ChatMessagePolicyViolationVerdictDetails_AllowFalsePositiveOverride        ChatMessagePolicyViolationVerdictDetails = "allowFalsePositiveOverride"
	ChatMessagePolicyViolationVerdictDetails_AllowOverrideWithJustification    ChatMessagePolicyViolationVerdictDetails = "allowOverrideWithJustification"
	ChatMessagePolicyViolationVerdictDetails_AllowOverrideWithoutJustification ChatMessagePolicyViolationVerdictDetails = "allowOverrideWithoutJustification"
	ChatMessagePolicyViolationVerdictDetails_None                              ChatMessagePolicyViolationVerdictDetails = "none"
)

func PossibleValuesForChatMessagePolicyViolationVerdictDetails() []string {
	return []string{
		string(ChatMessagePolicyViolationVerdictDetails_AllowFalsePositiveOverride),
		string(ChatMessagePolicyViolationVerdictDetails_AllowOverrideWithJustification),
		string(ChatMessagePolicyViolationVerdictDetails_AllowOverrideWithoutJustification),
		string(ChatMessagePolicyViolationVerdictDetails_None),
	}
}

func (s *ChatMessagePolicyViolationVerdictDetails) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatMessagePolicyViolationVerdictDetails(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatMessagePolicyViolationVerdictDetails(input string) (*ChatMessagePolicyViolationVerdictDetails, error) {
	vals := map[string]ChatMessagePolicyViolationVerdictDetails{
		"allowfalsepositiveoverride":        ChatMessagePolicyViolationVerdictDetails_AllowFalsePositiveOverride,
		"allowoverridewithjustification":    ChatMessagePolicyViolationVerdictDetails_AllowOverrideWithJustification,
		"allowoverridewithoutjustification": ChatMessagePolicyViolationVerdictDetails_AllowOverrideWithoutJustification,
		"none":                              ChatMessagePolicyViolationVerdictDetails_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessagePolicyViolationVerdictDetails(input)
	return &out, nil
}

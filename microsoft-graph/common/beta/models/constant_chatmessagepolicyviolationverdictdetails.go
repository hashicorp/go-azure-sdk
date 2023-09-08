package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessagePolicyViolationVerdictDetails string

const (
	ChatMessagePolicyViolationVerdictDetailsallowFalsePositiveOverride        ChatMessagePolicyViolationVerdictDetails = "AllowFalsePositiveOverride"
	ChatMessagePolicyViolationVerdictDetailsallowOverrideWithJustification    ChatMessagePolicyViolationVerdictDetails = "AllowOverrideWithJustification"
	ChatMessagePolicyViolationVerdictDetailsallowOverrideWithoutJustification ChatMessagePolicyViolationVerdictDetails = "AllowOverrideWithoutJustification"
	ChatMessagePolicyViolationVerdictDetailsnone                              ChatMessagePolicyViolationVerdictDetails = "None"
)

func PossibleValuesForChatMessagePolicyViolationVerdictDetails() []string {
	return []string{
		string(ChatMessagePolicyViolationVerdictDetailsallowFalsePositiveOverride),
		string(ChatMessagePolicyViolationVerdictDetailsallowOverrideWithJustification),
		string(ChatMessagePolicyViolationVerdictDetailsallowOverrideWithoutJustification),
		string(ChatMessagePolicyViolationVerdictDetailsnone),
	}
}

func parseChatMessagePolicyViolationVerdictDetails(input string) (*ChatMessagePolicyViolationVerdictDetails, error) {
	vals := map[string]ChatMessagePolicyViolationVerdictDetails{
		"allowfalsepositiveoverride":        ChatMessagePolicyViolationVerdictDetailsallowFalsePositiveOverride,
		"allowoverridewithjustification":    ChatMessagePolicyViolationVerdictDetailsallowOverrideWithJustification,
		"allowoverridewithoutjustification": ChatMessagePolicyViolationVerdictDetailsallowOverrideWithoutJustification,
		"none":                              ChatMessagePolicyViolationVerdictDetailsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessagePolicyViolationVerdictDetails(input)
	return &out, nil
}

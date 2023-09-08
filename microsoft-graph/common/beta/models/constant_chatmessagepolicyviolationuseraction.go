package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessagePolicyViolationUserAction string

const (
	ChatMessagePolicyViolationUserActionnone                ChatMessagePolicyViolationUserAction = "None"
	ChatMessagePolicyViolationUserActionoverride            ChatMessagePolicyViolationUserAction = "Override"
	ChatMessagePolicyViolationUserActionreportFalsePositive ChatMessagePolicyViolationUserAction = "ReportFalsePositive"
)

func PossibleValuesForChatMessagePolicyViolationUserAction() []string {
	return []string{
		string(ChatMessagePolicyViolationUserActionnone),
		string(ChatMessagePolicyViolationUserActionoverride),
		string(ChatMessagePolicyViolationUserActionreportFalsePositive),
	}
}

func parseChatMessagePolicyViolationUserAction(input string) (*ChatMessagePolicyViolationUserAction, error) {
	vals := map[string]ChatMessagePolicyViolationUserAction{
		"none":                ChatMessagePolicyViolationUserActionnone,
		"override":            ChatMessagePolicyViolationUserActionoverride,
		"reportfalsepositive": ChatMessagePolicyViolationUserActionreportFalsePositive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessagePolicyViolationUserAction(input)
	return &out, nil
}

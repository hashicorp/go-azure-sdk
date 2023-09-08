package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessagePolicyViolationDlpAction string

const (
	ChatMessagePolicyViolationDlpActionblockAccess         ChatMessagePolicyViolationDlpAction = "BlockAccess"
	ChatMessagePolicyViolationDlpActionblockAccessExternal ChatMessagePolicyViolationDlpAction = "BlockAccessExternal"
	ChatMessagePolicyViolationDlpActionnone                ChatMessagePolicyViolationDlpAction = "None"
	ChatMessagePolicyViolationDlpActionnotifySender        ChatMessagePolicyViolationDlpAction = "NotifySender"
)

func PossibleValuesForChatMessagePolicyViolationDlpAction() []string {
	return []string{
		string(ChatMessagePolicyViolationDlpActionblockAccess),
		string(ChatMessagePolicyViolationDlpActionblockAccessExternal),
		string(ChatMessagePolicyViolationDlpActionnone),
		string(ChatMessagePolicyViolationDlpActionnotifySender),
	}
}

func parseChatMessagePolicyViolationDlpAction(input string) (*ChatMessagePolicyViolationDlpAction, error) {
	vals := map[string]ChatMessagePolicyViolationDlpAction{
		"blockaccess":         ChatMessagePolicyViolationDlpActionblockAccess,
		"blockaccessexternal": ChatMessagePolicyViolationDlpActionblockAccessExternal,
		"none":                ChatMessagePolicyViolationDlpActionnone,
		"notifysender":        ChatMessagePolicyViolationDlpActionnotifySender,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessagePolicyViolationDlpAction(input)
	return &out, nil
}

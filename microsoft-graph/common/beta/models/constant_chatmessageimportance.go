package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageImportance string

const (
	ChatMessageImportancehigh   ChatMessageImportance = "High"
	ChatMessageImportancenormal ChatMessageImportance = "Normal"
	ChatMessageImportanceurgent ChatMessageImportance = "Urgent"
)

func PossibleValuesForChatMessageImportance() []string {
	return []string{
		string(ChatMessageImportancehigh),
		string(ChatMessageImportancenormal),
		string(ChatMessageImportanceurgent),
	}
}

func parseChatMessageImportance(input string) (*ChatMessageImportance, error) {
	vals := map[string]ChatMessageImportance{
		"high":   ChatMessageImportancehigh,
		"normal": ChatMessageImportancenormal,
		"urgent": ChatMessageImportanceurgent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessageImportance(input)
	return &out, nil
}

package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageRuleActionsMarkImportance string

const (
	MessageRuleActionsMarkImportancehigh   MessageRuleActionsMarkImportance = "High"
	MessageRuleActionsMarkImportancelow    MessageRuleActionsMarkImportance = "Low"
	MessageRuleActionsMarkImportancenormal MessageRuleActionsMarkImportance = "Normal"
)

func PossibleValuesForMessageRuleActionsMarkImportance() []string {
	return []string{
		string(MessageRuleActionsMarkImportancehigh),
		string(MessageRuleActionsMarkImportancelow),
		string(MessageRuleActionsMarkImportancenormal),
	}
}

func parseMessageRuleActionsMarkImportance(input string) (*MessageRuleActionsMarkImportance, error) {
	vals := map[string]MessageRuleActionsMarkImportance{
		"high":   MessageRuleActionsMarkImportancehigh,
		"low":    MessageRuleActionsMarkImportancelow,
		"normal": MessageRuleActionsMarkImportancenormal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageRuleActionsMarkImportance(input)
	return &out, nil
}

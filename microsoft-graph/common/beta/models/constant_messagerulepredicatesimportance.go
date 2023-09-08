package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageRulePredicatesImportance string

const (
	MessageRulePredicatesImportancehigh   MessageRulePredicatesImportance = "High"
	MessageRulePredicatesImportancelow    MessageRulePredicatesImportance = "Low"
	MessageRulePredicatesImportancenormal MessageRulePredicatesImportance = "Normal"
)

func PossibleValuesForMessageRulePredicatesImportance() []string {
	return []string{
		string(MessageRulePredicatesImportancehigh),
		string(MessageRulePredicatesImportancelow),
		string(MessageRulePredicatesImportancenormal),
	}
}

func parseMessageRulePredicatesImportance(input string) (*MessageRulePredicatesImportance, error) {
	vals := map[string]MessageRulePredicatesImportance{
		"high":   MessageRulePredicatesImportancehigh,
		"low":    MessageRulePredicatesImportancelow,
		"normal": MessageRulePredicatesImportancenormal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageRulePredicatesImportance(input)
	return &out, nil
}

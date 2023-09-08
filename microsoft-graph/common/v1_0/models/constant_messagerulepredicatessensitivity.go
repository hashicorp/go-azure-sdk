package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageRulePredicatesSensitivity string

const (
	MessageRulePredicatesSensitivityconfidential MessageRulePredicatesSensitivity = "Confidential"
	MessageRulePredicatesSensitivitynormal       MessageRulePredicatesSensitivity = "Normal"
	MessageRulePredicatesSensitivitypersonal     MessageRulePredicatesSensitivity = "Personal"
	MessageRulePredicatesSensitivityprivate      MessageRulePredicatesSensitivity = "Private"
)

func PossibleValuesForMessageRulePredicatesSensitivity() []string {
	return []string{
		string(MessageRulePredicatesSensitivityconfidential),
		string(MessageRulePredicatesSensitivitynormal),
		string(MessageRulePredicatesSensitivitypersonal),
		string(MessageRulePredicatesSensitivityprivate),
	}
}

func parseMessageRulePredicatesSensitivity(input string) (*MessageRulePredicatesSensitivity, error) {
	vals := map[string]MessageRulePredicatesSensitivity{
		"confidential": MessageRulePredicatesSensitivityconfidential,
		"normal":       MessageRulePredicatesSensitivitynormal,
		"personal":     MessageRulePredicatesSensitivitypersonal,
		"private":      MessageRulePredicatesSensitivityprivate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageRulePredicatesSensitivity(input)
	return &out, nil
}

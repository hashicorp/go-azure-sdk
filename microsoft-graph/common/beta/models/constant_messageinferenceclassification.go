package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageInferenceClassification string

const (
	MessageInferenceClassificationfocused MessageInferenceClassification = "Focused"
	MessageInferenceClassificationother   MessageInferenceClassification = "Other"
)

func PossibleValuesForMessageInferenceClassification() []string {
	return []string{
		string(MessageInferenceClassificationfocused),
		string(MessageInferenceClassificationother),
	}
}

func parseMessageInferenceClassification(input string) (*MessageInferenceClassification, error) {
	vals := map[string]MessageInferenceClassification{
		"focused": MessageInferenceClassificationfocused,
		"other":   MessageInferenceClassificationother,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageInferenceClassification(input)
	return &out, nil
}

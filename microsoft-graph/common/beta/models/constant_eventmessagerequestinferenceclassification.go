package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageRequestInferenceClassification string

const (
	EventMessageRequestInferenceClassificationfocused EventMessageRequestInferenceClassification = "Focused"
	EventMessageRequestInferenceClassificationother   EventMessageRequestInferenceClassification = "Other"
)

func PossibleValuesForEventMessageRequestInferenceClassification() []string {
	return []string{
		string(EventMessageRequestInferenceClassificationfocused),
		string(EventMessageRequestInferenceClassificationother),
	}
}

func parseEventMessageRequestInferenceClassification(input string) (*EventMessageRequestInferenceClassification, error) {
	vals := map[string]EventMessageRequestInferenceClassification{
		"focused": EventMessageRequestInferenceClassificationfocused,
		"other":   EventMessageRequestInferenceClassificationother,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageRequestInferenceClassification(input)
	return &out, nil
}

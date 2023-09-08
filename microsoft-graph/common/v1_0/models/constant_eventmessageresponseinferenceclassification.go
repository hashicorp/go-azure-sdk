package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageResponseInferenceClassification string

const (
	EventMessageResponseInferenceClassificationfocused EventMessageResponseInferenceClassification = "Focused"
	EventMessageResponseInferenceClassificationother   EventMessageResponseInferenceClassification = "Other"
)

func PossibleValuesForEventMessageResponseInferenceClassification() []string {
	return []string{
		string(EventMessageResponseInferenceClassificationfocused),
		string(EventMessageResponseInferenceClassificationother),
	}
}

func parseEventMessageResponseInferenceClassification(input string) (*EventMessageResponseInferenceClassification, error) {
	vals := map[string]EventMessageResponseInferenceClassification{
		"focused": EventMessageResponseInferenceClassificationfocused,
		"other":   EventMessageResponseInferenceClassificationother,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageResponseInferenceClassification(input)
	return &out, nil
}

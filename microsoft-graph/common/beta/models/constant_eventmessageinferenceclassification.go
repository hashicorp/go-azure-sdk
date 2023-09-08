package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageInferenceClassification string

const (
	EventMessageInferenceClassificationfocused EventMessageInferenceClassification = "Focused"
	EventMessageInferenceClassificationother   EventMessageInferenceClassification = "Other"
)

func PossibleValuesForEventMessageInferenceClassification() []string {
	return []string{
		string(EventMessageInferenceClassificationfocused),
		string(EventMessageInferenceClassificationother),
	}
}

func parseEventMessageInferenceClassification(input string) (*EventMessageInferenceClassification, error) {
	vals := map[string]EventMessageInferenceClassification{
		"focused": EventMessageInferenceClassificationfocused,
		"other":   EventMessageInferenceClassificationother,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageInferenceClassification(input)
	return &out, nil
}

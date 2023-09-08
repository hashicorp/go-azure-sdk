package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarSharingMessageInferenceClassification string

const (
	CalendarSharingMessageInferenceClassificationfocused CalendarSharingMessageInferenceClassification = "Focused"
	CalendarSharingMessageInferenceClassificationother   CalendarSharingMessageInferenceClassification = "Other"
)

func PossibleValuesForCalendarSharingMessageInferenceClassification() []string {
	return []string{
		string(CalendarSharingMessageInferenceClassificationfocused),
		string(CalendarSharingMessageInferenceClassificationother),
	}
}

func parseCalendarSharingMessageInferenceClassification(input string) (*CalendarSharingMessageInferenceClassification, error) {
	vals := map[string]CalendarSharingMessageInferenceClassification{
		"focused": CalendarSharingMessageInferenceClassificationfocused,
		"other":   CalendarSharingMessageInferenceClassificationother,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarSharingMessageInferenceClassification(input)
	return &out, nil
}

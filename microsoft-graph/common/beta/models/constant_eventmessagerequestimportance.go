package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageRequestImportance string

const (
	EventMessageRequestImportancehigh   EventMessageRequestImportance = "High"
	EventMessageRequestImportancelow    EventMessageRequestImportance = "Low"
	EventMessageRequestImportancenormal EventMessageRequestImportance = "Normal"
)

func PossibleValuesForEventMessageRequestImportance() []string {
	return []string{
		string(EventMessageRequestImportancehigh),
		string(EventMessageRequestImportancelow),
		string(EventMessageRequestImportancenormal),
	}
}

func parseEventMessageRequestImportance(input string) (*EventMessageRequestImportance, error) {
	vals := map[string]EventMessageRequestImportance{
		"high":   EventMessageRequestImportancehigh,
		"low":    EventMessageRequestImportancelow,
		"normal": EventMessageRequestImportancenormal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageRequestImportance(input)
	return &out, nil
}

package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageImportance string

const (
	EventMessageImportancehigh   EventMessageImportance = "High"
	EventMessageImportancelow    EventMessageImportance = "Low"
	EventMessageImportancenormal EventMessageImportance = "Normal"
)

func PossibleValuesForEventMessageImportance() []string {
	return []string{
		string(EventMessageImportancehigh),
		string(EventMessageImportancelow),
		string(EventMessageImportancenormal),
	}
}

func parseEventMessageImportance(input string) (*EventMessageImportance, error) {
	vals := map[string]EventMessageImportance{
		"high":   EventMessageImportancehigh,
		"low":    EventMessageImportancelow,
		"normal": EventMessageImportancenormal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageImportance(input)
	return &out, nil
}

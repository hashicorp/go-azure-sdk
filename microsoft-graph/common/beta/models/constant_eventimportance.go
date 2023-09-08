package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventImportance string

const (
	EventImportancehigh   EventImportance = "High"
	EventImportancelow    EventImportance = "Low"
	EventImportancenormal EventImportance = "Normal"
)

func PossibleValuesForEventImportance() []string {
	return []string{
		string(EventImportancehigh),
		string(EventImportancelow),
		string(EventImportancenormal),
	}
}

func parseEventImportance(input string) (*EventImportance, error) {
	vals := map[string]EventImportance{
		"high":   EventImportancehigh,
		"low":    EventImportancelow,
		"normal": EventImportancenormal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventImportance(input)
	return &out, nil
}

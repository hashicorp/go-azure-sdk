package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecurrencePatternIndex string

const (
	RecurrencePatternIndexfirst  RecurrencePatternIndex = "First"
	RecurrencePatternIndexfourth RecurrencePatternIndex = "Fourth"
	RecurrencePatternIndexlast   RecurrencePatternIndex = "Last"
	RecurrencePatternIndexsecond RecurrencePatternIndex = "Second"
	RecurrencePatternIndexthird  RecurrencePatternIndex = "Third"
)

func PossibleValuesForRecurrencePatternIndex() []string {
	return []string{
		string(RecurrencePatternIndexfirst),
		string(RecurrencePatternIndexfourth),
		string(RecurrencePatternIndexlast),
		string(RecurrencePatternIndexsecond),
		string(RecurrencePatternIndexthird),
	}
}

func parseRecurrencePatternIndex(input string) (*RecurrencePatternIndex, error) {
	vals := map[string]RecurrencePatternIndex{
		"first":  RecurrencePatternIndexfirst,
		"fourth": RecurrencePatternIndexfourth,
		"last":   RecurrencePatternIndexlast,
		"second": RecurrencePatternIndexsecond,
		"third":  RecurrencePatternIndexthird,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecurrencePatternIndex(input)
	return &out, nil
}

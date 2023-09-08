package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecurrencePatternType string

const (
	RecurrencePatternTypeabsoluteMonthly RecurrencePatternType = "AbsoluteMonthly"
	RecurrencePatternTypeabsoluteYearly  RecurrencePatternType = "AbsoluteYearly"
	RecurrencePatternTypedaily           RecurrencePatternType = "Daily"
	RecurrencePatternTyperelativeMonthly RecurrencePatternType = "RelativeMonthly"
	RecurrencePatternTyperelativeYearly  RecurrencePatternType = "RelativeYearly"
	RecurrencePatternTypeweekly          RecurrencePatternType = "Weekly"
)

func PossibleValuesForRecurrencePatternType() []string {
	return []string{
		string(RecurrencePatternTypeabsoluteMonthly),
		string(RecurrencePatternTypeabsoluteYearly),
		string(RecurrencePatternTypedaily),
		string(RecurrencePatternTyperelativeMonthly),
		string(RecurrencePatternTyperelativeYearly),
		string(RecurrencePatternTypeweekly),
	}
}

func parseRecurrencePatternType(input string) (*RecurrencePatternType, error) {
	vals := map[string]RecurrencePatternType{
		"absolutemonthly": RecurrencePatternTypeabsoluteMonthly,
		"absoluteyearly":  RecurrencePatternTypeabsoluteYearly,
		"daily":           RecurrencePatternTypedaily,
		"relativemonthly": RecurrencePatternTyperelativeMonthly,
		"relativeyearly":  RecurrencePatternTyperelativeYearly,
		"weekly":          RecurrencePatternTypeweekly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecurrencePatternType(input)
	return &out, nil
}

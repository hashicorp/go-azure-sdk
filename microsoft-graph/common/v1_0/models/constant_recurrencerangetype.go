package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecurrenceRangeType string

const (
	RecurrenceRangeTypeendDate  RecurrenceRangeType = "EndDate"
	RecurrenceRangeTypenoEnd    RecurrenceRangeType = "NoEnd"
	RecurrenceRangeTypenumbered RecurrenceRangeType = "Numbered"
)

func PossibleValuesForRecurrenceRangeType() []string {
	return []string{
		string(RecurrenceRangeTypeendDate),
		string(RecurrenceRangeTypenoEnd),
		string(RecurrenceRangeTypenumbered),
	}
}

func parseRecurrenceRangeType(input string) (*RecurrenceRangeType, error) {
	vals := map[string]RecurrenceRangeType{
		"enddate":  RecurrenceRangeTypeendDate,
		"noend":    RecurrenceRangeTypenoEnd,
		"numbered": RecurrenceRangeTypenumbered,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecurrenceRangeType(input)
	return &out, nil
}

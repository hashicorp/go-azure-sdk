package loganalytics

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntervalInMins string

const (
	IntervalInMinsFiveMins   IntervalInMins = "FiveMins"
	IntervalInMinsSixtyMins  IntervalInMins = "SixtyMins"
	IntervalInMinsThirtyMins IntervalInMins = "ThirtyMins"
	IntervalInMinsThreeMins  IntervalInMins = "ThreeMins"
)

func PossibleValuesForIntervalInMins() []string {
	return []string{
		string(IntervalInMinsFiveMins),
		string(IntervalInMinsSixtyMins),
		string(IntervalInMinsThirtyMins),
		string(IntervalInMinsThreeMins),
	}
}

func parseIntervalInMins(input string) (*IntervalInMins, error) {
	vals := map[string]IntervalInMins{
		"fivemins":   IntervalInMinsFiveMins,
		"sixtymins":  IntervalInMinsSixtyMins,
		"thirtymins": IntervalInMinsThirtyMins,
		"threemins":  IntervalInMinsThreeMins,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IntervalInMins(input)
	return &out, nil
}

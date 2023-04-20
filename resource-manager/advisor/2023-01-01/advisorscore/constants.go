package advisorscore

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Aggregated string

const (
	AggregatedDay   Aggregated = "day"
	AggregatedMonth Aggregated = "month"
	AggregatedWeek  Aggregated = "week"
)

func PossibleValuesForAggregated() []string {
	return []string{
		string(AggregatedDay),
		string(AggregatedMonth),
		string(AggregatedWeek),
	}
}

func parseAggregated(input string) (*Aggregated, error) {
	vals := map[string]Aggregated{
		"day":   AggregatedDay,
		"month": AggregatedMonth,
		"week":  AggregatedWeek,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Aggregated(input)
	return &out, nil
}

package advisorscore

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

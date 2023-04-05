package loganalytics

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

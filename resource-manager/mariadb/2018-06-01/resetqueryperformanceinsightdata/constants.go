package resetqueryperformanceinsightdata

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryPerformanceInsightResetDataResultState string

const (
	QueryPerformanceInsightResetDataResultStateFailed    QueryPerformanceInsightResetDataResultState = "Failed"
	QueryPerformanceInsightResetDataResultStateSucceeded QueryPerformanceInsightResetDataResultState = "Succeeded"
)

func PossibleValuesForQueryPerformanceInsightResetDataResultState() []string {
	return []string{
		string(QueryPerformanceInsightResetDataResultStateFailed),
		string(QueryPerformanceInsightResetDataResultStateSucceeded),
	}
}

func parseQueryPerformanceInsightResetDataResultState(input string) (*QueryPerformanceInsightResetDataResultState, error) {
	vals := map[string]QueryPerformanceInsightResetDataResultState{
		"failed":    QueryPerformanceInsightResetDataResultStateFailed,
		"succeeded": QueryPerformanceInsightResetDataResultStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := QueryPerformanceInsightResetDataResultState(input)
	return &out, nil
}

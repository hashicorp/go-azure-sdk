package resetqueryperformanceinsightdata

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

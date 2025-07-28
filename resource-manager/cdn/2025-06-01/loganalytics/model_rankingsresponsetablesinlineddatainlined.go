package loganalytics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RankingsResponseTablesInlinedDataInlined struct {
	Metrics *[]RankingsResponseTablesInlinedDataInlinedMetricsInlined `json:"metrics,omitempty"`
	Name    *string                                                   `json:"name,omitempty"`
}

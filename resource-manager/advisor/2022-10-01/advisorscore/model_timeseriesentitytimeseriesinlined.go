package advisorscore

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeSeriesEntityTimeSeriesInlined struct {
	AggregationLevel *Aggregated    `json:"aggregationLevel,omitempty"`
	ScoreHistory     *[]ScoreEntity `json:"scoreHistory,omitempty"`
}

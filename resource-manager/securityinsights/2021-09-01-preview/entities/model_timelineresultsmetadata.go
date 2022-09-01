package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimelineResultsMetadata struct {
	Aggregations []TimelineAggregation `json:"aggregations"`
	Errors       *[]TimelineError      `json:"errors,omitempty"`
	TotalCount   int64                 `json:"totalCount"`
}

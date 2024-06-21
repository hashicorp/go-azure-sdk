package workspaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchMetadata struct {
	AggregatedGroupingFields *string               `json:"aggregatedGroupingFields,omitempty"`
	AggregatedValueField     *string               `json:"aggregatedValueField,omitempty"`
	CoreSummaries            *[]CoreSummary        `json:"coreSummaries,omitempty"`
	ETag                     *string               `json:"eTag,omitempty"`
	Id                       *string               `json:"id,omitempty"`
	LastUpdated              *string               `json:"lastUpdated,omitempty"`
	Max                      *int64                `json:"max,omitempty"`
	RequestId                *string               `json:"requestId,omitempty"`
	RequestTime              *int64                `json:"requestTime,omitempty"`
	ResultType               *string               `json:"resultType,omitempty"`
	Schema                   *SearchMetadataSchema `json:"schema,omitempty"`
	Sort                     *[]SearchSort         `json:"sort,omitempty"`
	StartTime                *string               `json:"startTime,omitempty"`
	Status                   *string               `json:"status,omitempty"`
	Sum                      *int64                `json:"sum,omitempty"`
	Top                      *int64                `json:"top,omitempty"`
	Total                    *int64                `json:"total,omitempty"`
}

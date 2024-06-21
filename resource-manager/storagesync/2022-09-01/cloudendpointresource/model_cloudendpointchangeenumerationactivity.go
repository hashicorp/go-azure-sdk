package cloudendpointresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudEndpointChangeEnumerationActivity struct {
	DeletesProgressPercent    *int64                                          `json:"deletesProgressPercent,omitempty"`
	LastUpdatedTimestamp      *string                                         `json:"lastUpdatedTimestamp,omitempty"`
	MinutesRemaining          *int64                                          `json:"minutesRemaining,omitempty"`
	OperationState            *CloudEndpointChangeEnumerationActivityState    `json:"operationState,omitempty"`
	ProcessedDirectoriesCount *int64                                          `json:"processedDirectoriesCount,omitempty"`
	ProcessedFilesCount       *int64                                          `json:"processedFilesCount,omitempty"`
	ProgressPercent           *int64                                          `json:"progressPercent,omitempty"`
	StartedTimestamp          *string                                         `json:"startedTimestamp,omitempty"`
	StatusCode                *int64                                          `json:"statusCode,omitempty"`
	TotalCountsState          *CloudEndpointChangeEnumerationTotalCountsState `json:"totalCountsState,omitempty"`
	TotalDirectoriesCount     *int64                                          `json:"totalDirectoriesCount,omitempty"`
	TotalFilesCount           *int64                                          `json:"totalFilesCount,omitempty"`
	TotalSizeBytes            *int64                                          `json:"totalSizeBytes,omitempty"`
}

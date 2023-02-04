package hypervjobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobProperties struct {
	ActivityId      *string         `json:"activityId,omitempty"`
	ClientRequestId *string         `json:"clientRequestId,omitempty"`
	DisplayName     *string         `json:"displayName,omitempty"`
	EndTime         *string         `json:"endTime,omitempty"`
	Errors          *[]ErrorDetails `json:"errors,omitempty"`
	StartTime       *string         `json:"startTime,omitempty"`
	Status          *string         `json:"status,omitempty"`
}

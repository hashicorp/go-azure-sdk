package jobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Job struct {
	EndTime         *string          `json:"endTime,omitempty"`
	Error           *JobErrorDetails `json:"error,omitempty"`
	Id              *string          `json:"id,omitempty"`
	Name            *string          `json:"name,omitempty"`
	PercentComplete *int64           `json:"percentComplete,omitempty"`
	Properties      *JobProperties   `json:"properties,omitempty"`
	StartTime       *string          `json:"startTime,omitempty"`
	Status          *JobStatus       `json:"status,omitempty"`
	Type            *string          `json:"type,omitempty"`
}

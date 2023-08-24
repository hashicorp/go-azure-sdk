package replicationjobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobQueryParameter struct {
	AffectedObjectTypes *string                           `json:"affectedObjectTypes,omitempty"`
	EndTime             *string                           `json:"endTime,omitempty"`
	FabricId            *string                           `json:"fabricId,omitempty"`
	JobName             *string                           `json:"jobName,omitempty"`
	JobOutputType       *ExportJobOutputSerializationType `json:"jobOutputType,omitempty"`
	JobStatus           *string                           `json:"jobStatus,omitempty"`
	StartTime           *string                           `json:"startTime,omitempty"`
	TimezoneOffset      *float64                          `json:"timezoneOffset,omitempty"`
}

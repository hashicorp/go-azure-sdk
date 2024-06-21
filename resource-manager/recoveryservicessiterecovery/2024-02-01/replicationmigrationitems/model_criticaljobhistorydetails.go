package replicationmigrationitems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CriticalJobHistoryDetails struct {
	JobId     *string `json:"jobId,omitempty"`
	JobName   *string `json:"jobName,omitempty"`
	JobStatus *string `json:"jobStatus,omitempty"`
	StartTime *string `json:"startTime,omitempty"`
}

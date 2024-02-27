package replicationmigrationitems

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CriticalJobHistoryDetails struct {
	JobId     *string `json:"jobId,omitempty"`
	JobName   *string `json:"jobName,omitempty"`
	JobStatus *string `json:"jobStatus,omitempty"`
	StartTime *string `json:"startTime,omitempty"`
}

func (o *CriticalJobHistoryDetails) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

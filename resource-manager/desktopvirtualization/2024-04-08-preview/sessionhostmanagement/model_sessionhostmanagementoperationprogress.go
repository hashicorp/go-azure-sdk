package sessionhostmanagement

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SessionHostManagementOperationProgress struct {
	ExecutionStartTime         *string `json:"executionStartTime,omitempty"`
	SessionHostsCompleted      *int64  `json:"sessionHostsCompleted,omitempty"`
	SessionHostsInProgress     *int64  `json:"sessionHostsInProgress,omitempty"`
	SessionHostsRollbackFailed *int64  `json:"sessionHostsRollbackFailed,omitempty"`
	TotalSessionHosts          *int64  `json:"totalSessionHosts,omitempty"`
}

func (o *SessionHostManagementOperationProgress) GetExecutionStartTimeAsTime() (*time.Time, error) {
	if o.ExecutionStartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExecutionStartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *SessionHostManagementOperationProgress) SetExecutionStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExecutionStartTime = &formatted
}

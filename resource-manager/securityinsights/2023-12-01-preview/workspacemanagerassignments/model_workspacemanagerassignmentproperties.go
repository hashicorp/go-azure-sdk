package workspacemanagerassignments

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagerAssignmentProperties struct {
	Items                    []AssignmentItem   `json:"items"`
	LastJobEndTime           *string            `json:"lastJobEndTime,omitempty"`
	LastJobProvisioningState *ProvisioningState `json:"lastJobProvisioningState,omitempty"`
	TargetResourceName       string             `json:"targetResourceName"`
}

func (o *WorkspaceManagerAssignmentProperties) GetLastJobEndTimeAsTime() (*time.Time, error) {
	if o.LastJobEndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastJobEndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *WorkspaceManagerAssignmentProperties) SetLastJobEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastJobEndTime = &formatted
}

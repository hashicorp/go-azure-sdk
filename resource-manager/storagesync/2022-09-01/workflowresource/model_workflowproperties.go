package workflowresource

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowProperties struct {
	CommandName         *string             `json:"commandName,omitempty"`
	CreatedTimestamp    *string             `json:"createdTimestamp,omitempty"`
	LastOperationId     *string             `json:"lastOperationId,omitempty"`
	LastStatusTimestamp *string             `json:"lastStatusTimestamp,omitempty"`
	LastStepName        *string             `json:"lastStepName,omitempty"`
	Operation           *OperationDirection `json:"operation,omitempty"`
	Status              *WorkflowStatus     `json:"status,omitempty"`
	Steps               *string             `json:"steps,omitempty"`
}

func (o *WorkflowProperties) GetCreatedTimestampAsTime() (*time.Time, error) {
	if o.CreatedTimestamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedTimestamp, "2006-01-02T15:04:05Z07:00")
}

func (o *WorkflowProperties) SetCreatedTimestampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedTimestamp = &formatted
}

func (o *WorkflowProperties) GetLastStatusTimestampAsTime() (*time.Time, error) {
	if o.LastStatusTimestamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastStatusTimestamp, "2006-01-02T15:04:05Z07:00")
}

func (o *WorkflowProperties) SetLastStatusTimestampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastStatusTimestamp = &formatted
}

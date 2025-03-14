package experimentexecutions

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionStatus struct {
	ActionId   *string                                             `json:"actionId,omitempty"`
	ActionName *string                                             `json:"actionName,omitempty"`
	EndTime    *string                                             `json:"endTime,omitempty"`
	StartTime  *string                                             `json:"startTime,omitempty"`
	Status     *string                                             `json:"status,omitempty"`
	Targets    *[]ExperimentExecutionActionTargetDetailsProperties `json:"targets,omitempty"`
}

func (o *ActionStatus) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ActionStatus) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *ActionStatus) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ActionStatus) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
}

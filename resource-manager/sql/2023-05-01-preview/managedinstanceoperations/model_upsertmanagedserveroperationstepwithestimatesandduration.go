package managedinstanceoperations

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpsertManagedServerOperationStepWithEstimatesAndDuration struct {
	Name          *string `json:"name,omitempty"`
	Order         *int64  `json:"order,omitempty"`
	Status        *Status `json:"status,omitempty"`
	StepEndTime   *string `json:"stepEndTime,omitempty"`
	StepStartTime *string `json:"stepStartTime,omitempty"`
	TimeElapsed   *string `json:"timeElapsed,omitempty"`
}

func (o *UpsertManagedServerOperationStepWithEstimatesAndDuration) GetStepEndTimeAsTime() (*time.Time, error) {
	if o.StepEndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StepEndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *UpsertManagedServerOperationStepWithEstimatesAndDuration) SetStepEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StepEndTime = &formatted
}

func (o *UpsertManagedServerOperationStepWithEstimatesAndDuration) GetStepStartTimeAsTime() (*time.Time, error) {
	if o.StepStartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StepStartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *UpsertManagedServerOperationStepWithEstimatesAndDuration) SetStepStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StepStartTime = &formatted
}

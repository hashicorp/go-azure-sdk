package managedinstanceoperations

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstanceOperationProperties struct {
	Description             *string                                 `json:"description,omitempty"`
	ErrorCode               *int64                                  `json:"errorCode,omitempty"`
	ErrorDescription        *string                                 `json:"errorDescription,omitempty"`
	ErrorSeverity           *int64                                  `json:"errorSeverity,omitempty"`
	EstimatedCompletionTime *string                                 `json:"estimatedCompletionTime,omitempty"`
	IsCancellable           *bool                                   `json:"isCancellable,omitempty"`
	IsUserError             *bool                                   `json:"isUserError,omitempty"`
	ManagedInstanceName     *string                                 `json:"managedInstanceName,omitempty"`
	Operation               *string                                 `json:"operation,omitempty"`
	OperationFriendlyName   *string                                 `json:"operationFriendlyName,omitempty"`
	OperationParameters     *ManagedInstanceOperationParametersPair `json:"operationParameters,omitempty"`
	OperationSteps          *ManagedInstanceOperationSteps          `json:"operationSteps,omitempty"`
	PercentComplete         *int64                                  `json:"percentComplete,omitempty"`
	StartTime               *string                                 `json:"startTime,omitempty"`
	State                   *ManagementOperationState               `json:"state,omitempty"`
}

func (o *ManagedInstanceOperationProperties) GetEstimatedCompletionTimeAsTime() (*time.Time, error) {
	if o.EstimatedCompletionTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EstimatedCompletionTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ManagedInstanceOperationProperties) SetEstimatedCompletionTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EstimatedCompletionTime = &formatted
}

func (o *ManagedInstanceOperationProperties) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ManagedInstanceOperationProperties) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
}

package sparkjobdefinitions

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SparkBatchJobState struct {
	CurrentState       *string       `json:"currentState,omitempty"`
	DeadAt             *string       `json:"deadAt,omitempty"`
	JobCreationRequest *SparkRequest `json:"jobCreationRequest,omitempty"`
	KilledAt           *string       `json:"killedAt,omitempty"`
	NotStartedAt       *string       `json:"notStartedAt,omitempty"`
	RecoveringAt       *string       `json:"recoveringAt,omitempty"`
	RunningAt          *string       `json:"runningAt,omitempty"`
	StartingAt         *string       `json:"startingAt,omitempty"`
	SuccessAt          *string       `json:"successAt,omitempty"`
}

func (o *SparkBatchJobState) GetDeadAtAsTime() (*time.Time, error) {
	if o.DeadAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.DeadAt, "2006-01-02T15:04:05Z07:00")
}

func (o *SparkBatchJobState) SetDeadAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.DeadAt = &formatted
}

func (o *SparkBatchJobState) GetKilledAtAsTime() (*time.Time, error) {
	if o.KilledAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.KilledAt, "2006-01-02T15:04:05Z07:00")
}

func (o *SparkBatchJobState) SetKilledAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.KilledAt = &formatted
}

func (o *SparkBatchJobState) GetNotStartedAtAsTime() (*time.Time, error) {
	if o.NotStartedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.NotStartedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *SparkBatchJobState) SetNotStartedAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.NotStartedAt = &formatted
}

func (o *SparkBatchJobState) GetRecoveringAtAsTime() (*time.Time, error) {
	if o.RecoveringAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RecoveringAt, "2006-01-02T15:04:05Z07:00")
}

func (o *SparkBatchJobState) SetRecoveringAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RecoveringAt = &formatted
}

func (o *SparkBatchJobState) GetRunningAtAsTime() (*time.Time, error) {
	if o.RunningAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RunningAt, "2006-01-02T15:04:05Z07:00")
}

func (o *SparkBatchJobState) SetRunningAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RunningAt = &formatted
}

func (o *SparkBatchJobState) GetStartingAtAsTime() (*time.Time, error) {
	if o.StartingAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartingAt, "2006-01-02T15:04:05Z07:00")
}

func (o *SparkBatchJobState) SetStartingAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartingAt = &formatted
}

func (o *SparkBatchJobState) GetSuccessAtAsTime() (*time.Time, error) {
	if o.SuccessAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.SuccessAt, "2006-01-02T15:04:05Z07:00")
}

func (o *SparkBatchJobState) SetSuccessAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.SuccessAt = &formatted
}

package sparkjobdefinitions

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SparkScheduler struct {
	CancellationRequestedAt *string                `json:"cancellationRequestedAt,omitempty"`
	CurrentState            *SchedulerCurrentState `json:"currentState,omitempty"`
	EndedAt                 *string                `json:"endedAt,omitempty"`
	ScheduledAt             *string                `json:"scheduledAt,omitempty"`
	SubmittedAt             *string                `json:"submittedAt,omitempty"`
}

func (o *SparkScheduler) GetCancellationRequestedAtAsTime() (*time.Time, error) {
	if o.CancellationRequestedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CancellationRequestedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *SparkScheduler) SetCancellationRequestedAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CancellationRequestedAt = &formatted
}

func (o *SparkScheduler) GetEndedAtAsTime() (*time.Time, error) {
	if o.EndedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *SparkScheduler) SetEndedAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndedAt = &formatted
}

func (o *SparkScheduler) GetScheduledAtAsTime() (*time.Time, error) {
	if o.ScheduledAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ScheduledAt, "2006-01-02T15:04:05Z07:00")
}

func (o *SparkScheduler) SetScheduledAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ScheduledAt = &formatted
}

func (o *SparkScheduler) GetSubmittedAtAsTime() (*time.Time, error) {
	if o.SubmittedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.SubmittedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *SparkScheduler) SetSubmittedAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.SubmittedAt = &formatted
}

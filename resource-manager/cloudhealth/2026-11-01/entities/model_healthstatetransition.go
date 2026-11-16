package entities

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthStateTransition struct {
	NewState      HealthState `json:"newState"`
	OccurredAt    string      `json:"occurredAt"`
	PreviousState HealthState `json:"previousState"`
	Reason        *string     `json:"reason,omitempty"`
}

func (o *HealthStateTransition) GetOccurredAtAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.OccurredAt, "2006-01-02T15:04:05Z07:00")
}

func (o *HealthStateTransition) SetOccurredAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.OccurredAt = formatted
}

package jobschedules

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Schedule struct {
	DoNotRunAfter      *string `json:"doNotRunAfter,omitempty"`
	DoNotRunUntil      *string `json:"doNotRunUntil,omitempty"`
	RecurrenceInterval *string `json:"recurrenceInterval,omitempty"`
	StartWindow        *string `json:"startWindow,omitempty"`
}

func (o *Schedule) GetDoNotRunAfterAsTime() (*time.Time, error) {
	if o.DoNotRunAfter == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.DoNotRunAfter, "2006-01-02T15:04:05Z07:00")
}

func (o *Schedule) SetDoNotRunAfterAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.DoNotRunAfter = &formatted
}

func (o *Schedule) GetDoNotRunUntilAsTime() (*time.Time, error) {
	if o.DoNotRunUntil == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.DoNotRunUntil, "2006-01-02T15:04:05Z07:00")
}

func (o *Schedule) SetDoNotRunUntilAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.DoNotRunUntil = &formatted
}

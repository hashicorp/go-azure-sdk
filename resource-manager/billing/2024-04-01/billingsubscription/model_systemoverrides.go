package billingsubscription

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SystemOverrides struct {
	Cancellation               *Cancellation `json:"cancellation,omitempty"`
	CancellationAllowedEndDate *string       `json:"cancellationAllowedEndDate,omitempty"`
}

func (o *SystemOverrides) GetCancellationAllowedEndDateAsTime() (*time.Time, error) {
	if o.CancellationAllowedEndDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CancellationAllowedEndDate, "2006-01-02T15:04:05Z07:00")
}

func (o *SystemOverrides) SetCancellationAllowedEndDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CancellationAllowedEndDate = &formatted
}

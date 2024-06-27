package billingaccount

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransitionDetails struct {
	AnniversaryDay *int64  `json:"anniversaryDay,omitempty"`
	TransitionDate *string `json:"transitionDate,omitempty"`
}

func (o *TransitionDetails) GetTransitionDateAsTime() (*time.Time, error) {
	if o.TransitionDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.TransitionDate, "2006-01-02T15:04:05Z07:00")
}

func (o *TransitionDetails) SetTransitionDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.TransitionDate = &formatted
}

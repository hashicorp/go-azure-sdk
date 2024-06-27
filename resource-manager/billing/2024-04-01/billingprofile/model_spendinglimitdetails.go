package billingprofile

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SpendingLimitDetails struct {
	Amount    *float64             `json:"amount,omitempty"`
	Currency  *string              `json:"currency,omitempty"`
	EndDate   *string              `json:"endDate,omitempty"`
	StartDate *string              `json:"startDate,omitempty"`
	Status    *SpendingLimitStatus `json:"status,omitempty"`
	Type      *SpendingLimitType   `json:"type,omitempty"`
}

func (o *SpendingLimitDetails) GetEndDateAsTime() (*time.Time, error) {
	if o.EndDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndDate, "2006-01-02T15:04:05Z07:00")
}

func (o *SpendingLimitDetails) SetEndDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndDate = &formatted
}

func (o *SpendingLimitDetails) GetStartDateAsTime() (*time.Time, error) {
	if o.StartDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartDate, "2006-01-02T15:04:05Z07:00")
}

func (o *SpendingLimitDetails) SetStartDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartDate = &formatted
}

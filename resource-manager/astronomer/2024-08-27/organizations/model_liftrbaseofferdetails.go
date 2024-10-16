package organizations

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LiftrBaseOfferDetails struct {
	EndDate     *string      `json:"endDate,omitempty"`
	OfferId     string       `json:"offerId"`
	PlanId      string       `json:"planId"`
	PlanName    *string      `json:"planName,omitempty"`
	PublisherId string       `json:"publisherId"`
	RenewalMode *RenewalMode `json:"renewalMode,omitempty"`
	TermId      *string      `json:"termId,omitempty"`
	TermUnit    *string      `json:"termUnit,omitempty"`
}

func (o *LiftrBaseOfferDetails) GetEndDateAsTime() (*time.Time, error) {
	if o.EndDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndDate, "2006-01-02T15:04:05Z07:00")
}

func (o *LiftrBaseOfferDetails) SetEndDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndDate = &formatted
}

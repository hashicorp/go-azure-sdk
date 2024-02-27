package agreements

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgreementProperties struct {
	AcceptanceMode     *AcceptanceMode     `json:"acceptanceMode,omitempty"`
	AgreementLink      *string             `json:"agreementLink,omitempty"`
	BillingProfileInfo *BillingProfileInfo `json:"billingProfileInfo,omitempty"`
	Category           *Category           `json:"category,omitempty"`
	EffectiveDate      *string             `json:"effectiveDate,omitempty"`
	ExpirationDate     *string             `json:"expirationDate,omitempty"`
	Participants       *[]Participants     `json:"participants,omitempty"`
	Status             *string             `json:"status,omitempty"`
}

func (o *AgreementProperties) GetEffectiveDateAsTime() (*time.Time, error) {
	if o.EffectiveDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EffectiveDate, "2006-01-02T15:04:05Z07:00")
}

func (o *AgreementProperties) SetEffectiveDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EffectiveDate = &formatted
}

func (o *AgreementProperties) GetExpirationDateAsTime() (*time.Time, error) {
	if o.ExpirationDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpirationDate, "2006-01-02T15:04:05Z07:00")
}

func (o *AgreementProperties) SetExpirationDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpirationDate = &formatted
}

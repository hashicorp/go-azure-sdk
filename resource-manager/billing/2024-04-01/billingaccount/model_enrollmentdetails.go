package billingaccount

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentDetails struct {
	BillingCycle             *string                   `json:"billingCycle,omitempty"`
	Channel                  *string                   `json:"channel,omitempty"`
	Cloud                    *string                   `json:"cloud,omitempty"`
	CountryCode              *string                   `json:"countryCode,omitempty"`
	Currency                 *string                   `json:"currency,omitempty"`
	EndDate                  *string                   `json:"endDate,omitempty"`
	ExtendedTermOption       *ExtendedTermOption       `json:"extendedTermOption,omitempty"`
	IndirectRelationshipInfo *IndirectRelationshipInfo `json:"indirectRelationshipInfo,omitempty"`
	InvoiceRecipient         *string                   `json:"invoiceRecipient,omitempty"`
	Language                 *string                   `json:"language,omitempty"`
	MarkupStatus             *MarkupStatus             `json:"markupStatus,omitempty"`
	PoNumber                 *string                   `json:"poNumber,omitempty"`
	StartDate                *string                   `json:"startDate,omitempty"`
	SupportCoverage          *string                   `json:"supportCoverage,omitempty"`
	SupportLevel             *SupportLevel             `json:"supportLevel,omitempty"`
}

func (o *EnrollmentDetails) GetEndDateAsTime() (*time.Time, error) {
	if o.EndDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndDate, "2006-01-02T15:04:05Z07:00")
}

func (o *EnrollmentDetails) SetEndDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndDate = &formatted
}

func (o *EnrollmentDetails) GetStartDateAsTime() (*time.Time, error) {
	if o.StartDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartDate, "2006-01-02T15:04:05Z07:00")
}

func (o *EnrollmentDetails) SetStartDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartDate = &formatted
}

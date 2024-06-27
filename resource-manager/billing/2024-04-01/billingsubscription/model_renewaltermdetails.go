package billingsubscription

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RenewalTermDetails struct {
	BillingFrequency *string `json:"billingFrequency,omitempty"`
	ProductId        *string `json:"productId,omitempty"`
	ProductTypeId    *string `json:"productTypeId,omitempty"`
	Quantity         *int64  `json:"quantity,omitempty"`
	SkuId            *string `json:"skuId,omitempty"`
	TermDuration     *string `json:"termDuration,omitempty"`
	TermEndDate      *string `json:"termEndDate,omitempty"`
}

func (o *RenewalTermDetails) GetTermEndDateAsTime() (*time.Time, error) {
	if o.TermEndDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.TermEndDate, "2006-01-02T15:04:05Z07:00")
}

func (o *RenewalTermDetails) SetTermEndDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.TermEndDate = &formatted
}

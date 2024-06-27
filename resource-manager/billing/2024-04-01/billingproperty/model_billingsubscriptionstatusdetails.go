package billingproperty

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingSubscriptionStatusDetails struct {
	EffectiveDate *string                   `json:"effectiveDate,omitempty"`
	Reason        *SubscriptionStatusReason `json:"reason,omitempty"`
}

func (o *BillingSubscriptionStatusDetails) GetEffectiveDateAsTime() (*time.Time, error) {
	if o.EffectiveDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EffectiveDate, "2006-01-02T15:04:05Z07:00")
}

func (o *BillingSubscriptionStatusDetails) SetEffectiveDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EffectiveDate = &formatted
}

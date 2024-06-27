package availablebalance

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PaymentOnAccount struct {
	Amount                    *Amount              `json:"amount,omitempty"`
	BillingProfileDisplayName *string              `json:"billingProfileDisplayName,omitempty"`
	BillingProfileId          *string              `json:"billingProfileId,omitempty"`
	Date                      *string              `json:"date,omitempty"`
	InvoiceId                 *string              `json:"invoiceId,omitempty"`
	InvoiceName               *string              `json:"invoiceName,omitempty"`
	Type                      *PaymentMethodFamily `json:"type,omitempty"`
}

func (o *PaymentOnAccount) GetDateAsTime() (*time.Time, error) {
	if o.Date == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Date, "2006-01-02T15:04:05Z07:00")
}

func (o *PaymentOnAccount) SetDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Date = &formatted
}

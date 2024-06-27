package invoice

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FailedPayment struct {
	Date                *string              `json:"date,omitempty"`
	FailedPaymentReason *FailedPaymentReason `json:"failedPaymentReason,omitempty"`
}

func (o *FailedPayment) GetDateAsTime() (*time.Time, error) {
	if o.Date == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Date, "2006-01-02T15:04:05Z07:00")
}

func (o *FailedPayment) SetDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Date = &formatted
}

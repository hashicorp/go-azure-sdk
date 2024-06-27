package invoice

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RefundDetailsSummary struct {
	AmountRefunded    *Amount           `json:"amountRefunded,omitempty"`
	AmountRequested   *Amount           `json:"amountRequested,omitempty"`
	ApprovedOn        *string           `json:"approvedOn,omitempty"`
	CompletedOn       *string           `json:"completedOn,omitempty"`
	RebillInvoiceId   *string           `json:"rebillInvoiceId,omitempty"`
	RefundOperationId *string           `json:"refundOperationId,omitempty"`
	RefundReason      *RefundReasonCode `json:"refundReason,omitempty"`
	RefundStatus      *RefundStatus     `json:"refundStatus,omitempty"`
	RequestedOn       *string           `json:"requestedOn,omitempty"`
	TransactionCount  *int64            `json:"transactionCount,omitempty"`
}

func (o *RefundDetailsSummary) GetApprovedOnAsTime() (*time.Time, error) {
	if o.ApprovedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ApprovedOn, "2006-01-02T15:04:05Z07:00")
}

func (o *RefundDetailsSummary) SetApprovedOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ApprovedOn = &formatted
}

func (o *RefundDetailsSummary) GetCompletedOnAsTime() (*time.Time, error) {
	if o.CompletedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CompletedOn, "2006-01-02T15:04:05Z07:00")
}

func (o *RefundDetailsSummary) SetCompletedOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CompletedOn = &formatted
}

func (o *RefundDetailsSummary) GetRequestedOnAsTime() (*time.Time, error) {
	if o.RequestedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RequestedOn, "2006-01-02T15:04:05Z07:00")
}

func (o *RefundDetailsSummary) SetRequestedOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RequestedOn = &formatted
}

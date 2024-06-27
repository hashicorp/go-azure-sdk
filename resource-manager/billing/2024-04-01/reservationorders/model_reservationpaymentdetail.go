package reservationorders

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationPaymentDetail struct {
	BillingAccount       *string                        `json:"billingAccount,omitempty"`
	BillingCurrencyTotal *Price                         `json:"billingCurrencyTotal,omitempty"`
	DueDate              *string                        `json:"dueDate,omitempty"`
	ExtendedStatusInfo   *ReservationExtendedStatusInfo `json:"extendedStatusInfo,omitempty"`
	PaymentDate          *string                        `json:"paymentDate,omitempty"`
	PricingCurrencyTotal *Price                         `json:"pricingCurrencyTotal,omitempty"`
	Status               *PaymentStatus                 `json:"status,omitempty"`
}

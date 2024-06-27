package reservationorder

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationOrderBillingPlanInformation struct {
	NextPaymentDueDate   *string                     `json:"nextPaymentDueDate,omitempty"`
	PricingCurrencyTotal *Price                      `json:"pricingCurrencyTotal,omitempty"`
	StartDate            *string                     `json:"startDate,omitempty"`
	Transactions         *[]ReservationPaymentDetail `json:"transactions,omitempty"`
}

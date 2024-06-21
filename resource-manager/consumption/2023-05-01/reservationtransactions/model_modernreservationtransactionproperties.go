package reservationtransactions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ModernReservationTransactionProperties struct {
	Amount                     *float64 `json:"amount,omitempty"`
	ArmSkuName                 *string  `json:"armSkuName,omitempty"`
	BillingFrequency           *string  `json:"billingFrequency,omitempty"`
	BillingProfileId           *string  `json:"billingProfileId,omitempty"`
	BillingProfileName         *string  `json:"billingProfileName,omitempty"`
	Currency                   *string  `json:"currency,omitempty"`
	Description                *string  `json:"description,omitempty"`
	EventDate                  *string  `json:"eventDate,omitempty"`
	EventType                  *string  `json:"eventType,omitempty"`
	Invoice                    *string  `json:"invoice,omitempty"`
	InvoiceId                  *string  `json:"invoiceId,omitempty"`
	InvoiceSectionId           *string  `json:"invoiceSectionId,omitempty"`
	InvoiceSectionName         *string  `json:"invoiceSectionName,omitempty"`
	PurchasingSubscriptionGuid *string  `json:"purchasingSubscriptionGuid,omitempty"`
	PurchasingSubscriptionName *string  `json:"purchasingSubscriptionName,omitempty"`
	Quantity                   *float64 `json:"quantity,omitempty"`
	Region                     *string  `json:"region,omitempty"`
	ReservationOrderId         *string  `json:"reservationOrderId,omitempty"`
	ReservationOrderName       *string  `json:"reservationOrderName,omitempty"`
	Term                       *string  `json:"term,omitempty"`
}

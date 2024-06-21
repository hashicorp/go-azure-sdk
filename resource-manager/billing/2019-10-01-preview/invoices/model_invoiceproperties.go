package invoices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvoiceProperties struct {
	AmountDue                 *Amount               `json:"amountDue,omitempty"`
	AzurePrepaymentApplied    *Amount               `json:"azurePrepaymentApplied,omitempty"`
	BilledAmount              *Amount               `json:"billedAmount,omitempty"`
	BillingProfileDisplayName *string               `json:"billingProfileDisplayName,omitempty"`
	BillingProfileId          *string               `json:"billingProfileId,omitempty"`
	CreditAmount              *Amount               `json:"creditAmount,omitempty"`
	Documents                 *[]Document           `json:"documents,omitempty"`
	DueDate                   *string               `json:"dueDate,omitempty"`
	FreeAzureCreditApplied    *Amount               `json:"freeAzureCreditApplied,omitempty"`
	InvoiceDate               *string               `json:"invoiceDate,omitempty"`
	InvoicePeriodEndDate      *string               `json:"invoicePeriodEndDate,omitempty"`
	InvoicePeriodStartDate    *string               `json:"invoicePeriodStartDate,omitempty"`
	InvoiceType               *InvoiceType          `json:"invoiceType,omitempty"`
	IsMonthlyInvoice          *bool                 `json:"isMonthlyInvoice,omitempty"`
	Payments                  *[]PaymentProperties  `json:"payments,omitempty"`
	PurchaseOrderNumber       *string               `json:"purchaseOrderNumber,omitempty"`
	RebillDetails             *InvoiceRebillDetails `json:"rebillDetails,omitempty"`
	Status                    *InvoiceStatus        `json:"status,omitempty"`
	SubTotal                  *Amount               `json:"subTotal,omitempty"`
	SubscriptionId            *string               `json:"subscriptionId,omitempty"`
	TaxAmount                 *Amount               `json:"taxAmount,omitempty"`
	TotalAmount               *Amount               `json:"totalAmount,omitempty"`
}

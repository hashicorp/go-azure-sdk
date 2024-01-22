package invoices

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvoiceProperties struct {
	AmountDue                 *Amount                   `json:"amountDue,omitempty"`
	AzurePrepaymentApplied    *Amount                   `json:"azurePrepaymentApplied,omitempty"`
	BilledAmount              *Amount                   `json:"billedAmount,omitempty"`
	BilledDocumentId          *string                   `json:"billedDocumentId,omitempty"`
	BillingProfileDisplayName *string                   `json:"billingProfileDisplayName,omitempty"`
	BillingProfileId          *string                   `json:"billingProfileId,omitempty"`
	CreditAmount              *Amount                   `json:"creditAmount,omitempty"`
	CreditForDocumentId       *string                   `json:"creditForDocumentId,omitempty"`
	DocumentType              *InvoiceDocumentType      `json:"documentType,omitempty"`
	Documents                 *[]Document               `json:"documents,omitempty"`
	DueDate                   *string                   `json:"dueDate,omitempty"`
	FreeAzureCreditApplied    *Amount                   `json:"freeAzureCreditApplied,omitempty"`
	InvoiceDate               *string                   `json:"invoiceDate,omitempty"`
	InvoicePeriodEndDate      *string                   `json:"invoicePeriodEndDate,omitempty"`
	InvoicePeriodStartDate    *string                   `json:"invoicePeriodStartDate,omitempty"`
	InvoiceType               *InvoiceType              `json:"invoiceType,omitempty"`
	IsMonthlyInvoice          *bool                     `json:"isMonthlyInvoice,omitempty"`
	Payments                  *[]PaymentProperties      `json:"payments,omitempty"`
	PurchaseOrderNumber       *string                   `json:"purchaseOrderNumber,omitempty"`
	RebillDetails             *map[string]RebillDetails `json:"rebillDetails,omitempty"`
	Status                    *InvoiceStatus            `json:"status,omitempty"`
	SubTotal                  *Amount                   `json:"subTotal,omitempty"`
	SubscriptionId            *string                   `json:"subscriptionId,omitempty"`
	TaxAmount                 *Amount                   `json:"taxAmount,omitempty"`
	TotalAmount               *Amount                   `json:"totalAmount,omitempty"`
}

func (o *InvoiceProperties) GetDueDateAsTime() (*time.Time, error) {
	if o.DueDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.DueDate, "2006-01-02T15:04:05Z07:00")
}

func (o *InvoiceProperties) SetDueDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.DueDate = &formatted
}

func (o *InvoiceProperties) GetInvoiceDateAsTime() (*time.Time, error) {
	if o.InvoiceDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.InvoiceDate, "2006-01-02T15:04:05Z07:00")
}

func (o *InvoiceProperties) SetInvoiceDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.InvoiceDate = &formatted
}

func (o *InvoiceProperties) GetInvoicePeriodEndDateAsTime() (*time.Time, error) {
	if o.InvoicePeriodEndDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.InvoicePeriodEndDate, "2006-01-02T15:04:05Z07:00")
}

func (o *InvoiceProperties) SetInvoicePeriodEndDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.InvoicePeriodEndDate = &formatted
}

func (o *InvoiceProperties) GetInvoicePeriodStartDateAsTime() (*time.Time, error) {
	if o.InvoicePeriodStartDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.InvoicePeriodStartDate, "2006-01-02T15:04:05Z07:00")
}

func (o *InvoiceProperties) SetInvoicePeriodStartDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.InvoicePeriodStartDate = &formatted
}

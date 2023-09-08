package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SalesInvoiceOperationPredicate struct {
	BillToCustomerId               *string
	BillToCustomerNumber           *string
	BillToName                     *string
	CurrencyCode                   *string
	CurrencyId                     *string
	CustomerId                     *string
	CustomerName                   *string
	CustomerNumber                 *string
	CustomerPurchaseOrderReference *string
	DiscountAmount                 *float64
	DiscountAppliedBeforeTax       *bool
	DueDate                        *string
	Email                          *string
	ExternalDocumentNumber         *string
	Id                             *string
	InvoiceDate                    *string
	LastModifiedDateTime           *string
	Number                         *string
	ODataType                      *string
	OrderId                        *string
	OrderNumber                    *string
	PaymentTermsId                 *string
	PhoneNumber                    *string
	PricesIncludeTax               *bool
	Salesperson                    *string
	ShipToContact                  *string
	ShipToName                     *string
	ShipmentMethodId               *string
	Status                         *string
	TotalAmountExcludingTax        *float64
	TotalAmountIncludingTax        *float64
	TotalTaxAmount                 *float64
}

func (p SalesInvoiceOperationPredicate) Matches(input SalesInvoice) bool {

	if p.BillToCustomerId != nil && (input.BillToCustomerId == nil || *p.BillToCustomerId != *input.BillToCustomerId) {
		return false
	}

	if p.BillToCustomerNumber != nil && (input.BillToCustomerNumber == nil || *p.BillToCustomerNumber != *input.BillToCustomerNumber) {
		return false
	}

	if p.BillToName != nil && (input.BillToName == nil || *p.BillToName != *input.BillToName) {
		return false
	}

	if p.CurrencyCode != nil && (input.CurrencyCode == nil || *p.CurrencyCode != *input.CurrencyCode) {
		return false
	}

	if p.CurrencyId != nil && (input.CurrencyId == nil || *p.CurrencyId != *input.CurrencyId) {
		return false
	}

	if p.CustomerId != nil && (input.CustomerId == nil || *p.CustomerId != *input.CustomerId) {
		return false
	}

	if p.CustomerName != nil && (input.CustomerName == nil || *p.CustomerName != *input.CustomerName) {
		return false
	}

	if p.CustomerNumber != nil && (input.CustomerNumber == nil || *p.CustomerNumber != *input.CustomerNumber) {
		return false
	}

	if p.CustomerPurchaseOrderReference != nil && (input.CustomerPurchaseOrderReference == nil || *p.CustomerPurchaseOrderReference != *input.CustomerPurchaseOrderReference) {
		return false
	}

	if p.DiscountAmount != nil && (input.DiscountAmount == nil || *p.DiscountAmount != *input.DiscountAmount) {
		return false
	}

	if p.DiscountAppliedBeforeTax != nil && (input.DiscountAppliedBeforeTax == nil || *p.DiscountAppliedBeforeTax != *input.DiscountAppliedBeforeTax) {
		return false
	}

	if p.DueDate != nil && (input.DueDate == nil || *p.DueDate != *input.DueDate) {
		return false
	}

	if p.Email != nil && (input.Email == nil || *p.Email != *input.Email) {
		return false
	}

	if p.ExternalDocumentNumber != nil && (input.ExternalDocumentNumber == nil || *p.ExternalDocumentNumber != *input.ExternalDocumentNumber) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.InvoiceDate != nil && (input.InvoiceDate == nil || *p.InvoiceDate != *input.InvoiceDate) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.Number != nil && (input.Number == nil || *p.Number != *input.Number) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OrderId != nil && (input.OrderId == nil || *p.OrderId != *input.OrderId) {
		return false
	}

	if p.OrderNumber != nil && (input.OrderNumber == nil || *p.OrderNumber != *input.OrderNumber) {
		return false
	}

	if p.PaymentTermsId != nil && (input.PaymentTermsId == nil || *p.PaymentTermsId != *input.PaymentTermsId) {
		return false
	}

	if p.PhoneNumber != nil && (input.PhoneNumber == nil || *p.PhoneNumber != *input.PhoneNumber) {
		return false
	}

	if p.PricesIncludeTax != nil && (input.PricesIncludeTax == nil || *p.PricesIncludeTax != *input.PricesIncludeTax) {
		return false
	}

	if p.Salesperson != nil && (input.Salesperson == nil || *p.Salesperson != *input.Salesperson) {
		return false
	}

	if p.ShipToContact != nil && (input.ShipToContact == nil || *p.ShipToContact != *input.ShipToContact) {
		return false
	}

	if p.ShipToName != nil && (input.ShipToName == nil || *p.ShipToName != *input.ShipToName) {
		return false
	}

	if p.ShipmentMethodId != nil && (input.ShipmentMethodId == nil || *p.ShipmentMethodId != *input.ShipmentMethodId) {
		return false
	}

	if p.Status != nil && (input.Status == nil || *p.Status != *input.Status) {
		return false
	}

	if p.TotalAmountExcludingTax != nil && (input.TotalAmountExcludingTax == nil || *p.TotalAmountExcludingTax != *input.TotalAmountExcludingTax) {
		return false
	}

	if p.TotalAmountIncludingTax != nil && (input.TotalAmountIncludingTax == nil || *p.TotalAmountIncludingTax != *input.TotalAmountIncludingTax) {
		return false
	}

	if p.TotalTaxAmount != nil && (input.TotalTaxAmount == nil || *p.TotalTaxAmount != *input.TotalTaxAmount) {
		return false
	}

	return true
}

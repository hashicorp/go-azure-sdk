package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SalesOrder struct {
	BillToCustomerId         *string            `json:"billToCustomerId,omitempty"`
	BillToCustomerNumber     *string            `json:"billToCustomerNumber,omitempty"`
	BillToName               *string            `json:"billToName,omitempty"`
	BillingPostalAddress     *PostalAddressType `json:"billingPostalAddress,omitempty"`
	Currency                 *Currency          `json:"currency,omitempty"`
	CurrencyCode             *string            `json:"currencyCode,omitempty"`
	CurrencyId               *string            `json:"currencyId,omitempty"`
	Customer                 *Customer          `json:"customer,omitempty"`
	CustomerId               *string            `json:"customerId,omitempty"`
	CustomerName             *string            `json:"customerName,omitempty"`
	CustomerNumber           *string            `json:"customerNumber,omitempty"`
	DiscountAmount           *float64           `json:"discountAmount,omitempty"`
	DiscountAppliedBeforeTax *bool              `json:"discountAppliedBeforeTax,omitempty"`
	Email                    *string            `json:"email,omitempty"`
	ExternalDocumentNumber   *string            `json:"externalDocumentNumber,omitempty"`
	FullyShipped             *bool              `json:"fullyShipped,omitempty"`
	Id                       *string            `json:"id,omitempty"`
	LastModifiedDateTime     *string            `json:"lastModifiedDateTime,omitempty"`
	Number                   *string            `json:"number,omitempty"`
	ODataType                *string            `json:"@odata.type,omitempty"`
	OrderDate                *string            `json:"orderDate,omitempty"`
	PartialShipping          *bool              `json:"partialShipping,omitempty"`
	PaymentTerm              *PaymentTerm       `json:"paymentTerm,omitempty"`
	PaymentTermsId           *string            `json:"paymentTermsId,omitempty"`
	PhoneNumber              *string            `json:"phoneNumber,omitempty"`
	PricesIncludeTax         *bool              `json:"pricesIncludeTax,omitempty"`
	RequestedDeliveryDate    *string            `json:"requestedDeliveryDate,omitempty"`
	SalesOrderLines          *[]SalesOrderLine  `json:"salesOrderLines,omitempty"`
	Salesperson              *string            `json:"salesperson,omitempty"`
	SellingPostalAddress     *PostalAddressType `json:"sellingPostalAddress,omitempty"`
	ShipToContact            *string            `json:"shipToContact,omitempty"`
	ShipToName               *string            `json:"shipToName,omitempty"`
	ShippingPostalAddress    *PostalAddressType `json:"shippingPostalAddress,omitempty"`
	Status                   *string            `json:"status,omitempty"`
	TotalAmountExcludingTax  *float64           `json:"totalAmountExcludingTax,omitempty"`
	TotalAmountIncludingTax  *float64           `json:"totalAmountIncludingTax,omitempty"`
	TotalTaxAmount           *float64           `json:"totalTaxAmount,omitempty"`
}

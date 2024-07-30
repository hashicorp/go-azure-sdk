package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SalesInvoice struct {
	BillToCustomerId               *string             `json:"billToCustomerId,omitempty"`
	BillToCustomerNumber           *string             `json:"billToCustomerNumber,omitempty"`
	BillToName                     *string             `json:"billToName,omitempty"`
	BillingPostalAddress           *PostalAddressType  `json:"billingPostalAddress,omitempty"`
	Currency                       *Currency           `json:"currency,omitempty"`
	CurrencyCode                   *string             `json:"currencyCode,omitempty"`
	CurrencyId                     *string             `json:"currencyId,omitempty"`
	Customer                       *Customer           `json:"customer,omitempty"`
	CustomerId                     *string             `json:"customerId,omitempty"`
	CustomerName                   *string             `json:"customerName,omitempty"`
	CustomerNumber                 *string             `json:"customerNumber,omitempty"`
	CustomerPurchaseOrderReference *string             `json:"customerPurchaseOrderReference,omitempty"`
	DiscountAmount                 *float64            `json:"discountAmount,omitempty"`
	DiscountAppliedBeforeTax       *bool               `json:"discountAppliedBeforeTax,omitempty"`
	DueDate                        *string             `json:"dueDate,omitempty"`
	Email                          *string             `json:"email,omitempty"`
	ExternalDocumentNumber         *string             `json:"externalDocumentNumber,omitempty"`
	Id                             *string             `json:"id,omitempty"`
	InvoiceDate                    *string             `json:"invoiceDate,omitempty"`
	LastModifiedDateTime           *string             `json:"lastModifiedDateTime,omitempty"`
	Number                         *string             `json:"number,omitempty"`
	ODataType                      *string             `json:"@odata.type,omitempty"`
	OrderId                        *string             `json:"orderId,omitempty"`
	OrderNumber                    *string             `json:"orderNumber,omitempty"`
	PaymentTerm                    *PaymentTerm        `json:"paymentTerm,omitempty"`
	PaymentTermsId                 *string             `json:"paymentTermsId,omitempty"`
	PhoneNumber                    *string             `json:"phoneNumber,omitempty"`
	PricesIncludeTax               *bool               `json:"pricesIncludeTax,omitempty"`
	SalesInvoiceLines              *[]SalesInvoiceLine `json:"salesInvoiceLines,omitempty"`
	Salesperson                    *string             `json:"salesperson,omitempty"`
	SellingPostalAddress           *PostalAddressType  `json:"sellingPostalAddress,omitempty"`
	ShipToContact                  *string             `json:"shipToContact,omitempty"`
	ShipToName                     *string             `json:"shipToName,omitempty"`
	ShipmentMethod                 *ShipmentMethod     `json:"shipmentMethod,omitempty"`
	ShipmentMethodId               *string             `json:"shipmentMethodId,omitempty"`
	ShippingPostalAddress          *PostalAddressType  `json:"shippingPostalAddress,omitempty"`
	Status                         *string             `json:"status,omitempty"`
	TotalAmountExcludingTax        *float64            `json:"totalAmountExcludingTax,omitempty"`
	TotalAmountIncludingTax        *float64            `json:"totalAmountIncludingTax,omitempty"`
	TotalTaxAmount                 *float64            `json:"totalTaxAmount,omitempty"`
}

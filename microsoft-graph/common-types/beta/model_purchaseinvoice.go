package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PurchaseInvoice struct {
	BuyFromAddress           *PostalAddressType     `json:"buyFromAddress,omitempty"`
	Currency                 *Currency              `json:"currency,omitempty"`
	CurrencyCode             *string                `json:"currencyCode,omitempty"`
	CurrencyId               *string                `json:"currencyId,omitempty"`
	DiscountAmount           *float64               `json:"discountAmount,omitempty"`
	DiscountAppliedBeforeTax *bool                  `json:"discountAppliedBeforeTax,omitempty"`
	DueDate                  *string                `json:"dueDate,omitempty"`
	Id                       *string                `json:"id,omitempty"`
	InvoiceDate              *string                `json:"invoiceDate,omitempty"`
	LastModifiedDateTime     *string                `json:"lastModifiedDateTime,omitempty"`
	Number                   *string                `json:"number,omitempty"`
	ODataType                *string                `json:"@odata.type,omitempty"`
	PayToAddress             *PostalAddressType     `json:"payToAddress,omitempty"`
	PayToContact             *string                `json:"payToContact,omitempty"`
	PayToName                *string                `json:"payToName,omitempty"`
	PayToVendorId            *string                `json:"payToVendorId,omitempty"`
	PayToVendorNumber        *string                `json:"payToVendorNumber,omitempty"`
	PricesIncludeTax         *bool                  `json:"pricesIncludeTax,omitempty"`
	PurchaseInvoiceLines     *[]PurchaseInvoiceLine `json:"purchaseInvoiceLines,omitempty"`
	ShipToAddress            *PostalAddressType     `json:"shipToAddress,omitempty"`
	ShipToContact            *string                `json:"shipToContact,omitempty"`
	ShipToName               *string                `json:"shipToName,omitempty"`
	Status                   *string                `json:"status,omitempty"`
	TotalAmountExcludingTax  *float64               `json:"totalAmountExcludingTax,omitempty"`
	TotalAmountIncludingTax  *float64               `json:"totalAmountIncludingTax,omitempty"`
	TotalTaxAmount           *float64               `json:"totalTaxAmount,omitempty"`
	Vendor                   *Vendor                `json:"vendor,omitempty"`
	VendorId                 *string                `json:"vendorId,omitempty"`
	VendorInvoiceNumber      *string                `json:"vendorInvoiceNumber,omitempty"`
	VendorName               *string                `json:"vendorName,omitempty"`
	VendorNumber             *string                `json:"vendorNumber,omitempty"`
}

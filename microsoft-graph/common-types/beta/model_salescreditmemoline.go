package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SalesCreditMemoLine struct {
	Account                   *Account `json:"account,omitempty"`
	AccountId                 *string  `json:"accountId,omitempty"`
	AmountExcludingTax        *float64 `json:"amountExcludingTax,omitempty"`
	AmountIncludingTax        *float64 `json:"amountIncludingTax,omitempty"`
	Description               *string  `json:"description,omitempty"`
	DiscountAmount            *float64 `json:"discountAmount,omitempty"`
	DiscountAppliedBeforeTax  *bool    `json:"discountAppliedBeforeTax,omitempty"`
	DiscountPercent           *float64 `json:"discountPercent,omitempty"`
	DocumentId                *string  `json:"documentId,omitempty"`
	Id                        *string  `json:"id,omitempty"`
	InvoiceDiscountAllocation *float64 `json:"invoiceDiscountAllocation,omitempty"`
	Item                      *Item    `json:"item,omitempty"`
	ItemId                    *string  `json:"itemId,omitempty"`
	LineType                  *string  `json:"lineType,omitempty"`
	NetAmount                 *float64 `json:"netAmount,omitempty"`
	NetAmountIncludingTax     *float64 `json:"netAmountIncludingTax,omitempty"`
	NetTaxAmount              *float64 `json:"netTaxAmount,omitempty"`
	ODataType                 *string  `json:"@odata.type,omitempty"`
	Quantity                  *float64 `json:"quantity,omitempty"`
	Sequence                  *int64   `json:"sequence,omitempty"`
	ShipmentDate              *string  `json:"shipmentDate,omitempty"`
	TaxCode                   *string  `json:"taxCode,omitempty"`
	TaxPercent                *float64 `json:"taxPercent,omitempty"`
	TotalTaxAmount            *float64 `json:"totalTaxAmount,omitempty"`
	UnitOfMeasureId           *string  `json:"unitOfMeasureId,omitempty"`
	UnitPrice                 *float64 `json:"unitPrice,omitempty"`
}

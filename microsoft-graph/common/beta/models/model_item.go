package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Item struct {
	BaseUnitOfMeasureId  *string       `json:"baseUnitOfMeasureId,omitempty"`
	Blocked              *bool         `json:"blocked,omitempty"`
	DisplayName          *string       `json:"displayName,omitempty"`
	Gtin                 *string       `json:"gtin,omitempty"`
	Id                   *string       `json:"id,omitempty"`
	Inventory            *float64      `json:"inventory,omitempty"`
	ItemCategory         *ItemCategory `json:"itemCategory,omitempty"`
	ItemCategoryCode     *string       `json:"itemCategoryCode,omitempty"`
	ItemCategoryId       *string       `json:"itemCategoryId,omitempty"`
	LastModifiedDateTime *string       `json:"lastModifiedDateTime,omitempty"`
	Number               *string       `json:"number,omitempty"`
	ODataType            *string       `json:"@odata.type,omitempty"`
	Picture              *[]Picture    `json:"picture,omitempty"`
	PriceIncludesTax     *bool         `json:"priceIncludesTax,omitempty"`
	TaxGroupCode         *string       `json:"taxGroupCode,omitempty"`
	TaxGroupId           *string       `json:"taxGroupId,omitempty"`
	Type                 *string       `json:"type,omitempty"`
	UnitCost             *float64      `json:"unitCost,omitempty"`
	UnitPrice            *float64      `json:"unitPrice,omitempty"`
}

package pricesheet

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PriceSheetModel struct {
	Download    *MeterDetails           `json:"download"`
	NextLink    *string                 `json:"nextLink,omitempty"`
	Pricesheets *[]PriceSheetProperties `json:"pricesheets,omitempty"`
}

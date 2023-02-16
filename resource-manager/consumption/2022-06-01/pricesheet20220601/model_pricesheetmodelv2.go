package pricesheet20220601

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PriceSheetModelV2 struct {
	Download    *MeterDetailsV2           `json:"download,omitempty"`
	NextLink    *string                   `json:"nextLink,omitempty"`
	Pricesheets *[]PriceSheetPropertiesV2 `json:"pricesheets,omitempty"`
}

package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookChartAxis struct {
	Format         *WorkbookChartAxisFormat `json:"format,omitempty"`
	Id             *string                  `json:"id,omitempty"`
	MajorGridlines *WorkbookChartGridlines  `json:"majorGridlines,omitempty"`
	MajorUnit      *Json                    `json:"majorUnit,omitempty"`
	Maximum        *Json                    `json:"maximum,omitempty"`
	Minimum        *Json                    `json:"minimum,omitempty"`
	MinorGridlines *WorkbookChartGridlines  `json:"minorGridlines,omitempty"`
	MinorUnit      *Json                    `json:"minorUnit,omitempty"`
	ODataType      *string                  `json:"@odata.type,omitempty"`
	Title          *WorkbookChartAxisTitle  `json:"title,omitempty"`
}

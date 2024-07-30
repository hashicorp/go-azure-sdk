package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookChart struct {
	Axes       *WorkbookChartAxes       `json:"axes,omitempty"`
	DataLabels *WorkbookChartDataLabels `json:"dataLabels,omitempty"`
	Format     *WorkbookChartAreaFormat `json:"format,omitempty"`
	Height     *float64                 `json:"height,omitempty"`
	Id         *string                  `json:"id,omitempty"`
	Left       *float64                 `json:"left,omitempty"`
	Legend     *WorkbookChartLegend     `json:"legend,omitempty"`
	Name       *string                  `json:"name,omitempty"`
	ODataType  *string                  `json:"@odata.type,omitempty"`
	Series     *[]WorkbookChartSeries   `json:"series,omitempty"`
	Title      *WorkbookChartTitle      `json:"title,omitempty"`
	Top        *float64                 `json:"top,omitempty"`
	Width      *float64                 `json:"width,omitempty"`
	Worksheet  *WorkbookWorksheet       `json:"worksheet,omitempty"`
}

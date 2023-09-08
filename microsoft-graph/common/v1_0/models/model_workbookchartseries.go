package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookChartSeries struct {
	Format    *WorkbookChartSeriesFormat `json:"format,omitempty"`
	Id        *string                    `json:"id,omitempty"`
	Name      *string                    `json:"name,omitempty"`
	ODataType *string                    `json:"@odata.type,omitempty"`
	Points    *[]WorkbookChartPoint      `json:"points,omitempty"`
}

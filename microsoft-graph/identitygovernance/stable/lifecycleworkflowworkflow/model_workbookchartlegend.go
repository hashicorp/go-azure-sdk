package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookChartLegend struct {
	Format    *WorkbookChartLegendFormat `json:"format,omitempty"`
	Id        *string                    `json:"id,omitempty"`
	ODataType *string                    `json:"@odata.type,omitempty"`
	Overlay   *bool                      `json:"overlay,omitempty"`
	Position  *string                    `json:"position,omitempty"`
	Visible   *bool                      `json:"visible,omitempty"`
}

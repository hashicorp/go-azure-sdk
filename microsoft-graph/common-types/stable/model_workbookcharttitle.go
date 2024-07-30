package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookChartTitle struct {
	Format    *WorkbookChartTitleFormat `json:"format,omitempty"`
	Id        *string                   `json:"id,omitempty"`
	ODataType *string                   `json:"@odata.type,omitempty"`
	Overlay   *bool                     `json:"overlay,omitempty"`
	Text      *string                   `json:"text,omitempty"`
	Visible   *bool                     `json:"visible,omitempty"`
}

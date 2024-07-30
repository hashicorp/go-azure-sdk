package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookChartGridlines struct {
	Format    *WorkbookChartGridlinesFormat `json:"format,omitempty"`
	Id        *string                       `json:"id,omitempty"`
	ODataType *string                       `json:"@odata.type,omitempty"`
	Visible   *bool                         `json:"visible,omitempty"`
}

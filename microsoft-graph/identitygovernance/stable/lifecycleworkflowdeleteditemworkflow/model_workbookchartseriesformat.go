package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookChartSeriesFormat struct {
	Fill      *WorkbookChartFill       `json:"fill,omitempty"`
	Id        *string                  `json:"id,omitempty"`
	Line      *WorkbookChartLineFormat `json:"line,omitempty"`
	ODataType *string                  `json:"@odata.type,omitempty"`
}

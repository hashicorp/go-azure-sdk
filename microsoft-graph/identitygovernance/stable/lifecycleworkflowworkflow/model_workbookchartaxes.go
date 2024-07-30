package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookChartAxes struct {
	CategoryAxis *WorkbookChartAxis `json:"categoryAxis,omitempty"`
	Id           *string            `json:"id,omitempty"`
	ODataType    *string            `json:"@odata.type,omitempty"`
	SeriesAxis   *WorkbookChartAxis `json:"seriesAxis,omitempty"`
	ValueAxis    *WorkbookChartAxis `json:"valueAxis,omitempty"`
}

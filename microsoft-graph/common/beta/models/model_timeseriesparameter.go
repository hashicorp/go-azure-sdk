package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeSeriesParameter struct {
	EndDateTime   *string `json:"endDateTime,omitempty"`
	MetricName    *string `json:"metricName,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
	StartDateTime *string `json:"startDateTime,omitempty"`
}

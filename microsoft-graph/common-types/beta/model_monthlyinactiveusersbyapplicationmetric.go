package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonthlyInactiveUsersByApplicationMetric struct {
	AppId                      *string `json:"appId,omitempty"`
	FactDate                   *string `json:"factDate,omitempty"`
	Id                         *string `json:"id,omitempty"`
	Inactive30DayCount         *int64  `json:"inactive30DayCount,omitempty"`
	Inactive60DayCount         *int64  `json:"inactive60DayCount,omitempty"`
	Inactive90DayCount         *int64  `json:"inactive90DayCount,omitempty"`
	InactiveCalendarMonthCount *int64  `json:"inactiveCalendarMonthCount,omitempty"`
	ODataType                  *string `json:"@odata.type,omitempty"`
}

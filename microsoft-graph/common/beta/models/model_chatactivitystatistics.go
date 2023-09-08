package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatActivityStatistics struct {
	Activity     *ChatActivityStatisticsActivity `json:"activity,omitempty"`
	AfterHours   *string                         `json:"afterHours,omitempty"`
	Duration     *string                         `json:"duration,omitempty"`
	EndDate      *string                         `json:"endDate,omitempty"`
	Id           *string                         `json:"id,omitempty"`
	ODataType    *string                         `json:"@odata.type,omitempty"`
	StartDate    *string                         `json:"startDate,omitempty"`
	TimeZoneUsed *string                         `json:"timeZoneUsed,omitempty"`
}

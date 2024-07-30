package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingActivityStatistics struct {
	Activity     *MeetingActivityStatisticsActivity `json:"activity,omitempty"`
	AfterHours   *string                            `json:"afterHours,omitempty"`
	Conflicting  *string                            `json:"conflicting,omitempty"`
	Duration     *string                            `json:"duration,omitempty"`
	EndDate      *string                            `json:"endDate,omitempty"`
	Id           *string                            `json:"id,omitempty"`
	Long         *string                            `json:"long,omitempty"`
	Multitasking *string                            `json:"multitasking,omitempty"`
	ODataType    *string                            `json:"@odata.type,omitempty"`
	Organized    *string                            `json:"organized,omitempty"`
	Recurring    *string                            `json:"recurring,omitempty"`
	StartDate    *string                            `json:"startDate,omitempty"`
	TimeZoneUsed *string                            `json:"timeZoneUsed,omitempty"`
}

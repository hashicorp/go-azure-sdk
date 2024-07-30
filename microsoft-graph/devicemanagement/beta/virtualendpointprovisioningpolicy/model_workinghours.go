package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkingHours struct {
	DaysOfWeek *WorkingHoursDaysOfWeek `json:"daysOfWeek,omitempty"`
	EndTime    *string                 `json:"endTime,omitempty"`
	ODataType  *string                 `json:"@odata.type,omitempty"`
	StartTime  *string                 `json:"startTime,omitempty"`
	TimeZone   *TimeZoneBase           `json:"timeZone,omitempty"`
}

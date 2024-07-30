package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDateTimeConfiguration struct {
	DateFormat           *string `json:"dateFormat,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
	OfficeHoursEndTime   *string `json:"officeHoursEndTime,omitempty"`
	OfficeHoursStartTime *string `json:"officeHoursStartTime,omitempty"`
	TimeFormat           *string `json:"timeFormat,omitempty"`
	TimeZone             *string `json:"timeZone,omitempty"`
}

package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomUpdateTimeWindow struct {
	EndDay    *CustomUpdateTimeWindowEndDay   `json:"endDay,omitempty"`
	EndTime   *string                         `json:"endTime,omitempty"`
	ODataType *string                         `json:"@odata.type,omitempty"`
	StartDay  *CustomUpdateTimeWindowStartDay `json:"startDay,omitempty"`
	StartTime *string                         `json:"startTime,omitempty"`
}

package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StandardTimeZoneOffset struct {
	DayOccurrence *int64                           `json:"dayOccurrence,omitempty"`
	DayOfWeek     *StandardTimeZoneOffsetDayOfWeek `json:"dayOfWeek,omitempty"`
	Month         *int64                           `json:"month,omitempty"`
	ODataType     *string                          `json:"@odata.type,omitempty"`
	Time          *string                          `json:"time,omitempty"`
	Year          *int64                           `json:"year,omitempty"`
}

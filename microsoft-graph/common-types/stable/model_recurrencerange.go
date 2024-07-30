package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecurrenceRange struct {
	EndDate             *string              `json:"endDate,omitempty"`
	NumberOfOccurrences *int64               `json:"numberOfOccurrences,omitempty"`
	ODataType           *string              `json:"@odata.type,omitempty"`
	RecurrenceTimeZone  *string              `json:"recurrenceTimeZone,omitempty"`
	StartDate           *string              `json:"startDate,omitempty"`
	Type                *RecurrenceRangeType `json:"type,omitempty"`
}

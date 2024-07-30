package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShiftAvailability struct {
	ODataType  *string              `json:"@odata.type,omitempty"`
	Recurrence *PatternedRecurrence `json:"recurrence,omitempty"`
	TimeSlots  *[]TimeRange         `json:"timeSlots,omitempty"`
	TimeZone   *string              `json:"timeZone,omitempty"`
}

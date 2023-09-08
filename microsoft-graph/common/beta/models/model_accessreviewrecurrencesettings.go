package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewRecurrenceSettings struct {
	DurationInDays    *int64  `json:"durationInDays,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	RecurrenceCount   *int64  `json:"recurrenceCount,omitempty"`
	RecurrenceEndType *string `json:"recurrenceEndType,omitempty"`
	RecurrenceType    *string `json:"recurrenceType,omitempty"`
}

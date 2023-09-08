package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeCardEntry struct {
	Breaks        *[]TimeCardBreak `json:"breaks,omitempty"`
	ClockInEvent  *TimeCardEvent   `json:"clockInEvent,omitempty"`
	ClockOutEvent *TimeCardEvent   `json:"clockOutEvent,omitempty"`
	ODataType     *string          `json:"@odata.type,omitempty"`
}

package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingWorkTimeSlot struct {
	End       *string `json:"end,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Start     *string `json:"start,omitempty"`
}

package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingSchedulingPolicy struct {
	AllowStaffSelection      *bool   `json:"allowStaffSelection,omitempty"`
	MaximumAdvance           *string `json:"maximumAdvance,omitempty"`
	MinimumLeadTime          *string `json:"minimumLeadTime,omitempty"`
	ODataType                *string `json:"@odata.type,omitempty"`
	SendConfirmationsToOwner *bool   `json:"sendConfirmationsToOwner,omitempty"`
	TimeSlotInterval         *string `json:"timeSlotInterval,omitempty"`
}

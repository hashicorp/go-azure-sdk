package me

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeConstraint struct {
	ActivityDomain *TimeConstraintActivityDomain `json:"activityDomain,omitempty"`
	ODataType      *string                       `json:"@odata.type,omitempty"`
	TimeSlots      *[]TimeSlot                   `json:"timeSlots,omitempty"`
}

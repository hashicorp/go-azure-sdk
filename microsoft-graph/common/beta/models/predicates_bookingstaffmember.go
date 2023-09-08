package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingStaffMemberOperationPredicate struct {
	AvailabilityIsAffectedByPersonalCalendar *bool
	ColorIndex                               *int64
	DisplayName                              *string
	EmailAddress                             *string
	Id                                       *string
	IsEmailNotificationEnabled               *bool
	ODataType                                *string
	TimeZone                                 *string
	UseBusinessHours                         *bool
}

func (p BookingStaffMemberOperationPredicate) Matches(input BookingStaffMember) bool {

	if p.AvailabilityIsAffectedByPersonalCalendar != nil && (input.AvailabilityIsAffectedByPersonalCalendar == nil || *p.AvailabilityIsAffectedByPersonalCalendar != *input.AvailabilityIsAffectedByPersonalCalendar) {
		return false
	}

	if p.ColorIndex != nil && (input.ColorIndex == nil || *p.ColorIndex != *input.ColorIndex) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.EmailAddress != nil && (input.EmailAddress == nil || *p.EmailAddress != *input.EmailAddress) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsEmailNotificationEnabled != nil && (input.IsEmailNotificationEnabled == nil || *p.IsEmailNotificationEnabled != *input.IsEmailNotificationEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TimeZone != nil && (input.TimeZone == nil || *p.TimeZone != *input.TimeZone) {
		return false
	}

	if p.UseBusinessHours != nil && (input.UseBusinessHours == nil || *p.UseBusinessHours != *input.UseBusinessHours) {
		return false
	}

	return true
}

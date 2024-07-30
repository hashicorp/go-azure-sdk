package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingStaffMember struct {
	AvailabilityIsAffectedByPersonalCalendar *bool                               `json:"availabilityIsAffectedByPersonalCalendar,omitempty"`
	ColorIndex                               *int64                              `json:"colorIndex,omitempty"`
	DisplayName                              *string                             `json:"displayName,omitempty"`
	EmailAddress                             *string                             `json:"emailAddress,omitempty"`
	Id                                       *string                             `json:"id,omitempty"`
	IsEmailNotificationEnabled               *bool                               `json:"isEmailNotificationEnabled,omitempty"`
	MembershipStatus                         *BookingStaffMemberMembershipStatus `json:"membershipStatus,omitempty"`
	ODataType                                *string                             `json:"@odata.type,omitempty"`
	Role                                     *BookingStaffMemberRole             `json:"role,omitempty"`
	TimeZone                                 *string                             `json:"timeZone,omitempty"`
	UseBusinessHours                         *bool                               `json:"useBusinessHours,omitempty"`
	WorkingHours                             *[]BookingWorkHours                 `json:"workingHours,omitempty"`
}

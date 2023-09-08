package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingStaffMemberRole string

const (
	BookingStaffMemberRoleadministrator BookingStaffMemberRole = "Administrator"
	BookingStaffMemberRoleexternalGuest BookingStaffMemberRole = "ExternalGuest"
	BookingStaffMemberRoleguest         BookingStaffMemberRole = "Guest"
	BookingStaffMemberRolescheduler     BookingStaffMemberRole = "Scheduler"
	BookingStaffMemberRoleteamMember    BookingStaffMemberRole = "TeamMember"
	BookingStaffMemberRoleviewer        BookingStaffMemberRole = "Viewer"
)

func PossibleValuesForBookingStaffMemberRole() []string {
	return []string{
		string(BookingStaffMemberRoleadministrator),
		string(BookingStaffMemberRoleexternalGuest),
		string(BookingStaffMemberRoleguest),
		string(BookingStaffMemberRolescheduler),
		string(BookingStaffMemberRoleteamMember),
		string(BookingStaffMemberRoleviewer),
	}
}

func parseBookingStaffMemberRole(input string) (*BookingStaffMemberRole, error) {
	vals := map[string]BookingStaffMemberRole{
		"administrator": BookingStaffMemberRoleadministrator,
		"externalguest": BookingStaffMemberRoleexternalGuest,
		"guest":         BookingStaffMemberRoleguest,
		"scheduler":     BookingStaffMemberRolescheduler,
		"teammember":    BookingStaffMemberRoleteamMember,
		"viewer":        BookingStaffMemberRoleviewer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingStaffMemberRole(input)
	return &out, nil
}

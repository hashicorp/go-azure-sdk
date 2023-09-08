package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingStaffMemberMembershipStatus string

const (
	BookingStaffMemberMembershipStatusactive            BookingStaffMemberMembershipStatus = "Active"
	BookingStaffMemberMembershipStatuspendingAcceptance BookingStaffMemberMembershipStatus = "PendingAcceptance"
	BookingStaffMemberMembershipStatusrejectedByStaff   BookingStaffMemberMembershipStatus = "RejectedByStaff"
)

func PossibleValuesForBookingStaffMemberMembershipStatus() []string {
	return []string{
		string(BookingStaffMemberMembershipStatusactive),
		string(BookingStaffMemberMembershipStatuspendingAcceptance),
		string(BookingStaffMemberMembershipStatusrejectedByStaff),
	}
}

func parseBookingStaffMemberMembershipStatus(input string) (*BookingStaffMemberMembershipStatus, error) {
	vals := map[string]BookingStaffMemberMembershipStatus{
		"active":            BookingStaffMemberMembershipStatusactive,
		"pendingacceptance": BookingStaffMemberMembershipStatuspendingAcceptance,
		"rejectedbystaff":   BookingStaffMemberMembershipStatusrejectedByStaff,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingStaffMemberMembershipStatus(input)
	return &out, nil
}

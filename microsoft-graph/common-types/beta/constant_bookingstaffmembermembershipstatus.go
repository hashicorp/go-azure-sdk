package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingStaffMemberMembershipStatus string

const (
	BookingStaffMemberMembershipStatus_Active            BookingStaffMemberMembershipStatus = "active"
	BookingStaffMemberMembershipStatus_PendingAcceptance BookingStaffMemberMembershipStatus = "pendingAcceptance"
	BookingStaffMemberMembershipStatus_RejectedByStaff   BookingStaffMemberMembershipStatus = "rejectedByStaff"
)

func PossibleValuesForBookingStaffMemberMembershipStatus() []string {
	return []string{
		string(BookingStaffMemberMembershipStatus_Active),
		string(BookingStaffMemberMembershipStatus_PendingAcceptance),
		string(BookingStaffMemberMembershipStatus_RejectedByStaff),
	}
}

func (s *BookingStaffMemberMembershipStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBookingStaffMemberMembershipStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBookingStaffMemberMembershipStatus(input string) (*BookingStaffMemberMembershipStatus, error) {
	vals := map[string]BookingStaffMemberMembershipStatus{
		"active":            BookingStaffMemberMembershipStatus_Active,
		"pendingacceptance": BookingStaffMemberMembershipStatus_PendingAcceptance,
		"rejectedbystaff":   BookingStaffMemberMembershipStatus_RejectedByStaff,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingStaffMemberMembershipStatus(input)
	return &out, nil
}

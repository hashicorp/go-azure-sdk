package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingStaffMemberRole string

const (
	BookingStaffMemberRole_Administrator BookingStaffMemberRole = "administrator"
	BookingStaffMemberRole_ExternalGuest BookingStaffMemberRole = "externalGuest"
	BookingStaffMemberRole_Guest         BookingStaffMemberRole = "guest"
	BookingStaffMemberRole_Scheduler     BookingStaffMemberRole = "scheduler"
	BookingStaffMemberRole_TeamMember    BookingStaffMemberRole = "teamMember"
	BookingStaffMemberRole_Viewer        BookingStaffMemberRole = "viewer"
)

func PossibleValuesForBookingStaffMemberRole() []string {
	return []string{
		string(BookingStaffMemberRole_Administrator),
		string(BookingStaffMemberRole_ExternalGuest),
		string(BookingStaffMemberRole_Guest),
		string(BookingStaffMemberRole_Scheduler),
		string(BookingStaffMemberRole_TeamMember),
		string(BookingStaffMemberRole_Viewer),
	}
}

func (s *BookingStaffMemberRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBookingStaffMemberRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBookingStaffMemberRole(input string) (*BookingStaffMemberRole, error) {
	vals := map[string]BookingStaffMemberRole{
		"administrator": BookingStaffMemberRole_Administrator,
		"externalguest": BookingStaffMemberRole_ExternalGuest,
		"guest":         BookingStaffMemberRole_Guest,
		"scheduler":     BookingStaffMemberRole_Scheduler,
		"teammember":    BookingStaffMemberRole_TeamMember,
		"viewer":        BookingStaffMemberRole_Viewer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingStaffMemberRole(input)
	return &out, nil
}

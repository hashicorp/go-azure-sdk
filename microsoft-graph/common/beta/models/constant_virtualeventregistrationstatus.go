package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventRegistrationStatus string

const (
	VirtualEventRegistrationStatuscanceled            VirtualEventRegistrationStatus = "Canceled"
	VirtualEventRegistrationStatuspendingApproval     VirtualEventRegistrationStatus = "PendingApproval"
	VirtualEventRegistrationStatusregistered          VirtualEventRegistrationStatus = "Registered"
	VirtualEventRegistrationStatusrejectedByOrganizer VirtualEventRegistrationStatus = "RejectedByOrganizer"
	VirtualEventRegistrationStatuswaitlisted          VirtualEventRegistrationStatus = "Waitlisted"
)

func PossibleValuesForVirtualEventRegistrationStatus() []string {
	return []string{
		string(VirtualEventRegistrationStatuscanceled),
		string(VirtualEventRegistrationStatuspendingApproval),
		string(VirtualEventRegistrationStatusregistered),
		string(VirtualEventRegistrationStatusrejectedByOrganizer),
		string(VirtualEventRegistrationStatuswaitlisted),
	}
}

func parseVirtualEventRegistrationStatus(input string) (*VirtualEventRegistrationStatus, error) {
	vals := map[string]VirtualEventRegistrationStatus{
		"canceled":            VirtualEventRegistrationStatuscanceled,
		"pendingapproval":     VirtualEventRegistrationStatuspendingApproval,
		"registered":          VirtualEventRegistrationStatusregistered,
		"rejectedbyorganizer": VirtualEventRegistrationStatusrejectedByOrganizer,
		"waitlisted":          VirtualEventRegistrationStatuswaitlisted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventRegistrationStatus(input)
	return &out, nil
}

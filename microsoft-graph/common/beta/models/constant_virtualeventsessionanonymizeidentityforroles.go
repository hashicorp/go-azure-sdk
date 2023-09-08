package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventSessionAnonymizeIdentityForRoles string

const (
	VirtualEventSessionAnonymizeIdentityForRolesattendee    VirtualEventSessionAnonymizeIdentityForRoles = "Attendee"
	VirtualEventSessionAnonymizeIdentityForRolescoorganizer VirtualEventSessionAnonymizeIdentityForRoles = "Coorganizer"
	VirtualEventSessionAnonymizeIdentityForRolespresenter   VirtualEventSessionAnonymizeIdentityForRoles = "Presenter"
	VirtualEventSessionAnonymizeIdentityForRolesproducer    VirtualEventSessionAnonymizeIdentityForRoles = "Producer"
)

func PossibleValuesForVirtualEventSessionAnonymizeIdentityForRoles() []string {
	return []string{
		string(VirtualEventSessionAnonymizeIdentityForRolesattendee),
		string(VirtualEventSessionAnonymizeIdentityForRolescoorganizer),
		string(VirtualEventSessionAnonymizeIdentityForRolespresenter),
		string(VirtualEventSessionAnonymizeIdentityForRolesproducer),
	}
}

func parseVirtualEventSessionAnonymizeIdentityForRoles(input string) (*VirtualEventSessionAnonymizeIdentityForRoles, error) {
	vals := map[string]VirtualEventSessionAnonymizeIdentityForRoles{
		"attendee":    VirtualEventSessionAnonymizeIdentityForRolesattendee,
		"coorganizer": VirtualEventSessionAnonymizeIdentityForRolescoorganizer,
		"presenter":   VirtualEventSessionAnonymizeIdentityForRolespresenter,
		"producer":    VirtualEventSessionAnonymizeIdentityForRolesproducer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventSessionAnonymizeIdentityForRoles(input)
	return &out, nil
}

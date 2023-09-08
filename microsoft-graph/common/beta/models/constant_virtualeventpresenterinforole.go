package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventPresenterInfoRole string

const (
	VirtualEventPresenterInfoRoleattendee    VirtualEventPresenterInfoRole = "Attendee"
	VirtualEventPresenterInfoRolecoorganizer VirtualEventPresenterInfoRole = "Coorganizer"
	VirtualEventPresenterInfoRolepresenter   VirtualEventPresenterInfoRole = "Presenter"
	VirtualEventPresenterInfoRoleproducer    VirtualEventPresenterInfoRole = "Producer"
)

func PossibleValuesForVirtualEventPresenterInfoRole() []string {
	return []string{
		string(VirtualEventPresenterInfoRoleattendee),
		string(VirtualEventPresenterInfoRolecoorganizer),
		string(VirtualEventPresenterInfoRolepresenter),
		string(VirtualEventPresenterInfoRoleproducer),
	}
}

func parseVirtualEventPresenterInfoRole(input string) (*VirtualEventPresenterInfoRole, error) {
	vals := map[string]VirtualEventPresenterInfoRole{
		"attendee":    VirtualEventPresenterInfoRoleattendee,
		"coorganizer": VirtualEventPresenterInfoRolecoorganizer,
		"presenter":   VirtualEventPresenterInfoRolepresenter,
		"producer":    VirtualEventPresenterInfoRoleproducer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventPresenterInfoRole(input)
	return &out, nil
}

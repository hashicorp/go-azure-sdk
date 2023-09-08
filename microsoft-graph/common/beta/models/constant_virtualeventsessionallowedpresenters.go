package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventSessionAllowedPresenters string

const (
	VirtualEventSessionAllowedPresenterseveryone        VirtualEventSessionAllowedPresenters = "Everyone"
	VirtualEventSessionAllowedPresentersorganization    VirtualEventSessionAllowedPresenters = "Organization"
	VirtualEventSessionAllowedPresentersorganizer       VirtualEventSessionAllowedPresenters = "Organizer"
	VirtualEventSessionAllowedPresentersroleIsPresenter VirtualEventSessionAllowedPresenters = "RoleIsPresenter"
)

func PossibleValuesForVirtualEventSessionAllowedPresenters() []string {
	return []string{
		string(VirtualEventSessionAllowedPresenterseveryone),
		string(VirtualEventSessionAllowedPresentersorganization),
		string(VirtualEventSessionAllowedPresentersorganizer),
		string(VirtualEventSessionAllowedPresentersroleIsPresenter),
	}
}

func parseVirtualEventSessionAllowedPresenters(input string) (*VirtualEventSessionAllowedPresenters, error) {
	vals := map[string]VirtualEventSessionAllowedPresenters{
		"everyone":        VirtualEventSessionAllowedPresenterseveryone,
		"organization":    VirtualEventSessionAllowedPresentersorganization,
		"organizer":       VirtualEventSessionAllowedPresentersorganizer,
		"roleispresenter": VirtualEventSessionAllowedPresentersroleIsPresenter,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventSessionAllowedPresenters(input)
	return &out, nil
}

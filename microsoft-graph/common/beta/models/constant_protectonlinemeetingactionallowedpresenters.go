package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectOnlineMeetingActionAllowedPresenters string

const (
	ProtectOnlineMeetingActionAllowedPresenterseveryone        ProtectOnlineMeetingActionAllowedPresenters = "Everyone"
	ProtectOnlineMeetingActionAllowedPresentersorganization    ProtectOnlineMeetingActionAllowedPresenters = "Organization"
	ProtectOnlineMeetingActionAllowedPresentersorganizer       ProtectOnlineMeetingActionAllowedPresenters = "Organizer"
	ProtectOnlineMeetingActionAllowedPresentersroleIsPresenter ProtectOnlineMeetingActionAllowedPresenters = "RoleIsPresenter"
)

func PossibleValuesForProtectOnlineMeetingActionAllowedPresenters() []string {
	return []string{
		string(ProtectOnlineMeetingActionAllowedPresenterseveryone),
		string(ProtectOnlineMeetingActionAllowedPresentersorganization),
		string(ProtectOnlineMeetingActionAllowedPresentersorganizer),
		string(ProtectOnlineMeetingActionAllowedPresentersroleIsPresenter),
	}
}

func parseProtectOnlineMeetingActionAllowedPresenters(input string) (*ProtectOnlineMeetingActionAllowedPresenters, error) {
	vals := map[string]ProtectOnlineMeetingActionAllowedPresenters{
		"everyone":        ProtectOnlineMeetingActionAllowedPresenterseveryone,
		"organization":    ProtectOnlineMeetingActionAllowedPresentersorganization,
		"organizer":       ProtectOnlineMeetingActionAllowedPresentersorganizer,
		"roleispresenter": ProtectOnlineMeetingActionAllowedPresentersroleIsPresenter,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectOnlineMeetingActionAllowedPresenters(input)
	return &out, nil
}

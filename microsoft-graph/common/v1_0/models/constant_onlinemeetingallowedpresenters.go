package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingAllowedPresenters string

const (
	OnlineMeetingAllowedPresenterseveryone        OnlineMeetingAllowedPresenters = "Everyone"
	OnlineMeetingAllowedPresentersorganization    OnlineMeetingAllowedPresenters = "Organization"
	OnlineMeetingAllowedPresentersorganizer       OnlineMeetingAllowedPresenters = "Organizer"
	OnlineMeetingAllowedPresentersroleIsPresenter OnlineMeetingAllowedPresenters = "RoleIsPresenter"
)

func PossibleValuesForOnlineMeetingAllowedPresenters() []string {
	return []string{
		string(OnlineMeetingAllowedPresenterseveryone),
		string(OnlineMeetingAllowedPresentersorganization),
		string(OnlineMeetingAllowedPresentersorganizer),
		string(OnlineMeetingAllowedPresentersroleIsPresenter),
	}
}

func parseOnlineMeetingAllowedPresenters(input string) (*OnlineMeetingAllowedPresenters, error) {
	vals := map[string]OnlineMeetingAllowedPresenters{
		"everyone":        OnlineMeetingAllowedPresenterseveryone,
		"organization":    OnlineMeetingAllowedPresentersorganization,
		"organizer":       OnlineMeetingAllowedPresentersorganizer,
		"roleispresenter": OnlineMeetingAllowedPresentersroleIsPresenter,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingAllowedPresenters(input)
	return &out, nil
}

package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectOnlineMeetingActionAllowedPresenters string

const (
	ProtectOnlineMeetingActionAllowedPresenters_Everyone        ProtectOnlineMeetingActionAllowedPresenters = "everyone"
	ProtectOnlineMeetingActionAllowedPresenters_Organization    ProtectOnlineMeetingActionAllowedPresenters = "organization"
	ProtectOnlineMeetingActionAllowedPresenters_Organizer       ProtectOnlineMeetingActionAllowedPresenters = "organizer"
	ProtectOnlineMeetingActionAllowedPresenters_RoleIsPresenter ProtectOnlineMeetingActionAllowedPresenters = "roleIsPresenter"
)

func PossibleValuesForProtectOnlineMeetingActionAllowedPresenters() []string {
	return []string{
		string(ProtectOnlineMeetingActionAllowedPresenters_Everyone),
		string(ProtectOnlineMeetingActionAllowedPresenters_Organization),
		string(ProtectOnlineMeetingActionAllowedPresenters_Organizer),
		string(ProtectOnlineMeetingActionAllowedPresenters_RoleIsPresenter),
	}
}

func (s *ProtectOnlineMeetingActionAllowedPresenters) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProtectOnlineMeetingActionAllowedPresenters(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProtectOnlineMeetingActionAllowedPresenters(input string) (*ProtectOnlineMeetingActionAllowedPresenters, error) {
	vals := map[string]ProtectOnlineMeetingActionAllowedPresenters{
		"everyone":        ProtectOnlineMeetingActionAllowedPresenters_Everyone,
		"organization":    ProtectOnlineMeetingActionAllowedPresenters_Organization,
		"organizer":       ProtectOnlineMeetingActionAllowedPresenters_Organizer,
		"roleispresenter": ProtectOnlineMeetingActionAllowedPresenters_RoleIsPresenter,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectOnlineMeetingActionAllowedPresenters(input)
	return &out, nil
}

package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventSessionAllowedPresenters string

const (
	VirtualEventSessionAllowedPresenters_Everyone        VirtualEventSessionAllowedPresenters = "everyone"
	VirtualEventSessionAllowedPresenters_Organization    VirtualEventSessionAllowedPresenters = "organization"
	VirtualEventSessionAllowedPresenters_Organizer       VirtualEventSessionAllowedPresenters = "organizer"
	VirtualEventSessionAllowedPresenters_RoleIsPresenter VirtualEventSessionAllowedPresenters = "roleIsPresenter"
)

func PossibleValuesForVirtualEventSessionAllowedPresenters() []string {
	return []string{
		string(VirtualEventSessionAllowedPresenters_Everyone),
		string(VirtualEventSessionAllowedPresenters_Organization),
		string(VirtualEventSessionAllowedPresenters_Organizer),
		string(VirtualEventSessionAllowedPresenters_RoleIsPresenter),
	}
}

func (s *VirtualEventSessionAllowedPresenters) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualEventSessionAllowedPresenters(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualEventSessionAllowedPresenters(input string) (*VirtualEventSessionAllowedPresenters, error) {
	vals := map[string]VirtualEventSessionAllowedPresenters{
		"everyone":        VirtualEventSessionAllowedPresenters_Everyone,
		"organization":    VirtualEventSessionAllowedPresenters_Organization,
		"organizer":       VirtualEventSessionAllowedPresenters_Organizer,
		"roleispresenter": VirtualEventSessionAllowedPresenters_RoleIsPresenter,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventSessionAllowedPresenters(input)
	return &out, nil
}

package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventSessionAnonymizeIdentityForRoles string

const (
	VirtualEventSessionAnonymizeIdentityForRoles_Attendee    VirtualEventSessionAnonymizeIdentityForRoles = "attendee"
	VirtualEventSessionAnonymizeIdentityForRoles_Coorganizer VirtualEventSessionAnonymizeIdentityForRoles = "coorganizer"
	VirtualEventSessionAnonymizeIdentityForRoles_Presenter   VirtualEventSessionAnonymizeIdentityForRoles = "presenter"
	VirtualEventSessionAnonymizeIdentityForRoles_Producer    VirtualEventSessionAnonymizeIdentityForRoles = "producer"
)

func PossibleValuesForVirtualEventSessionAnonymizeIdentityForRoles() []string {
	return []string{
		string(VirtualEventSessionAnonymizeIdentityForRoles_Attendee),
		string(VirtualEventSessionAnonymizeIdentityForRoles_Coorganizer),
		string(VirtualEventSessionAnonymizeIdentityForRoles_Presenter),
		string(VirtualEventSessionAnonymizeIdentityForRoles_Producer),
	}
}

func (s *VirtualEventSessionAnonymizeIdentityForRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualEventSessionAnonymizeIdentityForRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualEventSessionAnonymizeIdentityForRoles(input string) (*VirtualEventSessionAnonymizeIdentityForRoles, error) {
	vals := map[string]VirtualEventSessionAnonymizeIdentityForRoles{
		"attendee":    VirtualEventSessionAnonymizeIdentityForRoles_Attendee,
		"coorganizer": VirtualEventSessionAnonymizeIdentityForRoles_Coorganizer,
		"presenter":   VirtualEventSessionAnonymizeIdentityForRoles_Presenter,
		"producer":    VirtualEventSessionAnonymizeIdentityForRoles_Producer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventSessionAnonymizeIdentityForRoles(input)
	return &out, nil
}

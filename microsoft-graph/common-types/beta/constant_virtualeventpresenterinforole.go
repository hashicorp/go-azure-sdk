package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventPresenterInfoRole string

const (
	VirtualEventPresenterInfoRole_Attendee    VirtualEventPresenterInfoRole = "attendee"
	VirtualEventPresenterInfoRole_Coorganizer VirtualEventPresenterInfoRole = "coorganizer"
	VirtualEventPresenterInfoRole_Presenter   VirtualEventPresenterInfoRole = "presenter"
	VirtualEventPresenterInfoRole_Producer    VirtualEventPresenterInfoRole = "producer"
)

func PossibleValuesForVirtualEventPresenterInfoRole() []string {
	return []string{
		string(VirtualEventPresenterInfoRole_Attendee),
		string(VirtualEventPresenterInfoRole_Coorganizer),
		string(VirtualEventPresenterInfoRole_Presenter),
		string(VirtualEventPresenterInfoRole_Producer),
	}
}

func (s *VirtualEventPresenterInfoRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualEventPresenterInfoRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualEventPresenterInfoRole(input string) (*VirtualEventPresenterInfoRole, error) {
	vals := map[string]VirtualEventPresenterInfoRole{
		"attendee":    VirtualEventPresenterInfoRole_Attendee,
		"coorganizer": VirtualEventPresenterInfoRole_Coorganizer,
		"presenter":   VirtualEventPresenterInfoRole_Presenter,
		"producer":    VirtualEventPresenterInfoRole_Producer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventPresenterInfoRole(input)
	return &out, nil
}

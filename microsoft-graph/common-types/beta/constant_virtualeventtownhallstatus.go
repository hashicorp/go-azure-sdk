package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventTownhallStatus string

const (
	VirtualEventTownhallStatus_Canceled  VirtualEventTownhallStatus = "canceled"
	VirtualEventTownhallStatus_Draft     VirtualEventTownhallStatus = "draft"
	VirtualEventTownhallStatus_Published VirtualEventTownhallStatus = "published"
)

func PossibleValuesForVirtualEventTownhallStatus() []string {
	return []string{
		string(VirtualEventTownhallStatus_Canceled),
		string(VirtualEventTownhallStatus_Draft),
		string(VirtualEventTownhallStatus_Published),
	}
}

func (s *VirtualEventTownhallStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualEventTownhallStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualEventTownhallStatus(input string) (*VirtualEventTownhallStatus, error) {
	vals := map[string]VirtualEventTownhallStatus{
		"canceled":  VirtualEventTownhallStatus_Canceled,
		"draft":     VirtualEventTownhallStatus_Draft,
		"published": VirtualEventTownhallStatus_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventTownhallStatus(input)
	return &out, nil
}

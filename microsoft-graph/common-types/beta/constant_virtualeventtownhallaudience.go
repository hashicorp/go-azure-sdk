package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventTownhallAudience string

const (
	VirtualEventTownhallAudience_Everyone     VirtualEventTownhallAudience = "everyone"
	VirtualEventTownhallAudience_Organization VirtualEventTownhallAudience = "organization"
)

func PossibleValuesForVirtualEventTownhallAudience() []string {
	return []string{
		string(VirtualEventTownhallAudience_Everyone),
		string(VirtualEventTownhallAudience_Organization),
	}
}

func (s *VirtualEventTownhallAudience) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualEventTownhallAudience(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualEventTownhallAudience(input string) (*VirtualEventTownhallAudience, error) {
	vals := map[string]VirtualEventTownhallAudience{
		"everyone":     VirtualEventTownhallAudience_Everyone,
		"organization": VirtualEventTownhallAudience_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventTownhallAudience(input)
	return &out, nil
}

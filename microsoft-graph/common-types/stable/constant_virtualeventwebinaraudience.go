package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventWebinarAudience string

const (
	VirtualEventWebinarAudience_Everyone     VirtualEventWebinarAudience = "everyone"
	VirtualEventWebinarAudience_Organization VirtualEventWebinarAudience = "organization"
)

func PossibleValuesForVirtualEventWebinarAudience() []string {
	return []string{
		string(VirtualEventWebinarAudience_Everyone),
		string(VirtualEventWebinarAudience_Organization),
	}
}

func (s *VirtualEventWebinarAudience) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualEventWebinarAudience(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualEventWebinarAudience(input string) (*VirtualEventWebinarAudience, error) {
	vals := map[string]VirtualEventWebinarAudience{
		"everyone":     VirtualEventWebinarAudience_Everyone,
		"organization": VirtualEventWebinarAudience_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventWebinarAudience(input)
	return &out, nil
}

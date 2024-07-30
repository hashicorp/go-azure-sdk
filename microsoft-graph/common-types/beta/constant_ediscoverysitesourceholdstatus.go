package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoverySiteSourceHoldStatus string

const (
	EdiscoverySiteSourceHoldStatus_Applied    EdiscoverySiteSourceHoldStatus = "applied"
	EdiscoverySiteSourceHoldStatus_Applying   EdiscoverySiteSourceHoldStatus = "applying"
	EdiscoverySiteSourceHoldStatus_NotApplied EdiscoverySiteSourceHoldStatus = "notApplied"
	EdiscoverySiteSourceHoldStatus_Partial    EdiscoverySiteSourceHoldStatus = "partial"
	EdiscoverySiteSourceHoldStatus_Removing   EdiscoverySiteSourceHoldStatus = "removing"
)

func PossibleValuesForEdiscoverySiteSourceHoldStatus() []string {
	return []string{
		string(EdiscoverySiteSourceHoldStatus_Applied),
		string(EdiscoverySiteSourceHoldStatus_Applying),
		string(EdiscoverySiteSourceHoldStatus_NotApplied),
		string(EdiscoverySiteSourceHoldStatus_Partial),
		string(EdiscoverySiteSourceHoldStatus_Removing),
	}
}

func (s *EdiscoverySiteSourceHoldStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoverySiteSourceHoldStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoverySiteSourceHoldStatus(input string) (*EdiscoverySiteSourceHoldStatus, error) {
	vals := map[string]EdiscoverySiteSourceHoldStatus{
		"applied":    EdiscoverySiteSourceHoldStatus_Applied,
		"applying":   EdiscoverySiteSourceHoldStatus_Applying,
		"notapplied": EdiscoverySiteSourceHoldStatus_NotApplied,
		"partial":    EdiscoverySiteSourceHoldStatus_Partial,
		"removing":   EdiscoverySiteSourceHoldStatus_Removing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoverySiteSourceHoldStatus(input)
	return &out, nil
}

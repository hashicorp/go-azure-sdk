package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryUserSourceHoldStatus string

const (
	EdiscoveryUserSourceHoldStatus_Applied    EdiscoveryUserSourceHoldStatus = "applied"
	EdiscoveryUserSourceHoldStatus_Applying   EdiscoveryUserSourceHoldStatus = "applying"
	EdiscoveryUserSourceHoldStatus_NotApplied EdiscoveryUserSourceHoldStatus = "notApplied"
	EdiscoveryUserSourceHoldStatus_Partial    EdiscoveryUserSourceHoldStatus = "partial"
	EdiscoveryUserSourceHoldStatus_Removing   EdiscoveryUserSourceHoldStatus = "removing"
)

func PossibleValuesForEdiscoveryUserSourceHoldStatus() []string {
	return []string{
		string(EdiscoveryUserSourceHoldStatus_Applied),
		string(EdiscoveryUserSourceHoldStatus_Applying),
		string(EdiscoveryUserSourceHoldStatus_NotApplied),
		string(EdiscoveryUserSourceHoldStatus_Partial),
		string(EdiscoveryUserSourceHoldStatus_Removing),
	}
}

func (s *EdiscoveryUserSourceHoldStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryUserSourceHoldStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryUserSourceHoldStatus(input string) (*EdiscoveryUserSourceHoldStatus, error) {
	vals := map[string]EdiscoveryUserSourceHoldStatus{
		"applied":    EdiscoveryUserSourceHoldStatus_Applied,
		"applying":   EdiscoveryUserSourceHoldStatus_Applying,
		"notapplied": EdiscoveryUserSourceHoldStatus_NotApplied,
		"partial":    EdiscoveryUserSourceHoldStatus_Partial,
		"removing":   EdiscoveryUserSourceHoldStatus_Removing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryUserSourceHoldStatus(input)
	return &out, nil
}

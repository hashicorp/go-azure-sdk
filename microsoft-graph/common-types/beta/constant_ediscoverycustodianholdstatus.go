package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCustodianHoldStatus string

const (
	EdiscoveryCustodianHoldStatus_Applied    EdiscoveryCustodianHoldStatus = "applied"
	EdiscoveryCustodianHoldStatus_Applying   EdiscoveryCustodianHoldStatus = "applying"
	EdiscoveryCustodianHoldStatus_NotApplied EdiscoveryCustodianHoldStatus = "notApplied"
	EdiscoveryCustodianHoldStatus_Partial    EdiscoveryCustodianHoldStatus = "partial"
	EdiscoveryCustodianHoldStatus_Removing   EdiscoveryCustodianHoldStatus = "removing"
)

func PossibleValuesForEdiscoveryCustodianHoldStatus() []string {
	return []string{
		string(EdiscoveryCustodianHoldStatus_Applied),
		string(EdiscoveryCustodianHoldStatus_Applying),
		string(EdiscoveryCustodianHoldStatus_NotApplied),
		string(EdiscoveryCustodianHoldStatus_Partial),
		string(EdiscoveryCustodianHoldStatus_Removing),
	}
}

func (s *EdiscoveryCustodianHoldStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryCustodianHoldStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryCustodianHoldStatus(input string) (*EdiscoveryCustodianHoldStatus, error) {
	vals := map[string]EdiscoveryCustodianHoldStatus{
		"applied":    EdiscoveryCustodianHoldStatus_Applied,
		"applying":   EdiscoveryCustodianHoldStatus_Applying,
		"notapplied": EdiscoveryCustodianHoldStatus_NotApplied,
		"partial":    EdiscoveryCustodianHoldStatus_Partial,
		"removing":   EdiscoveryCustodianHoldStatus_Removing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCustodianHoldStatus(input)
	return &out, nil
}

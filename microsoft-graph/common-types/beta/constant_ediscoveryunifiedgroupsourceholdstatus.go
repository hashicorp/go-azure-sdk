package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryUnifiedGroupSourceHoldStatus string

const (
	EdiscoveryUnifiedGroupSourceHoldStatus_Applied    EdiscoveryUnifiedGroupSourceHoldStatus = "applied"
	EdiscoveryUnifiedGroupSourceHoldStatus_Applying   EdiscoveryUnifiedGroupSourceHoldStatus = "applying"
	EdiscoveryUnifiedGroupSourceHoldStatus_NotApplied EdiscoveryUnifiedGroupSourceHoldStatus = "notApplied"
	EdiscoveryUnifiedGroupSourceHoldStatus_Partial    EdiscoveryUnifiedGroupSourceHoldStatus = "partial"
	EdiscoveryUnifiedGroupSourceHoldStatus_Removing   EdiscoveryUnifiedGroupSourceHoldStatus = "removing"
)

func PossibleValuesForEdiscoveryUnifiedGroupSourceHoldStatus() []string {
	return []string{
		string(EdiscoveryUnifiedGroupSourceHoldStatus_Applied),
		string(EdiscoveryUnifiedGroupSourceHoldStatus_Applying),
		string(EdiscoveryUnifiedGroupSourceHoldStatus_NotApplied),
		string(EdiscoveryUnifiedGroupSourceHoldStatus_Partial),
		string(EdiscoveryUnifiedGroupSourceHoldStatus_Removing),
	}
}

func (s *EdiscoveryUnifiedGroupSourceHoldStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryUnifiedGroupSourceHoldStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryUnifiedGroupSourceHoldStatus(input string) (*EdiscoveryUnifiedGroupSourceHoldStatus, error) {
	vals := map[string]EdiscoveryUnifiedGroupSourceHoldStatus{
		"applied":    EdiscoveryUnifiedGroupSourceHoldStatus_Applied,
		"applying":   EdiscoveryUnifiedGroupSourceHoldStatus_Applying,
		"notapplied": EdiscoveryUnifiedGroupSourceHoldStatus_NotApplied,
		"partial":    EdiscoveryUnifiedGroupSourceHoldStatus_Partial,
		"removing":   EdiscoveryUnifiedGroupSourceHoldStatus_Removing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryUnifiedGroupSourceHoldStatus(input)
	return &out, nil
}

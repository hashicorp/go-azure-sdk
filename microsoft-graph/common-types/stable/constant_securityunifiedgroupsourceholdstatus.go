package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUnifiedGroupSourceHoldStatus string

const (
	SecurityUnifiedGroupSourceHoldStatus_Applied    SecurityUnifiedGroupSourceHoldStatus = "applied"
	SecurityUnifiedGroupSourceHoldStatus_Applying   SecurityUnifiedGroupSourceHoldStatus = "applying"
	SecurityUnifiedGroupSourceHoldStatus_NotApplied SecurityUnifiedGroupSourceHoldStatus = "notApplied"
	SecurityUnifiedGroupSourceHoldStatus_Partial    SecurityUnifiedGroupSourceHoldStatus = "partial"
	SecurityUnifiedGroupSourceHoldStatus_Removing   SecurityUnifiedGroupSourceHoldStatus = "removing"
)

func PossibleValuesForSecurityUnifiedGroupSourceHoldStatus() []string {
	return []string{
		string(SecurityUnifiedGroupSourceHoldStatus_Applied),
		string(SecurityUnifiedGroupSourceHoldStatus_Applying),
		string(SecurityUnifiedGroupSourceHoldStatus_NotApplied),
		string(SecurityUnifiedGroupSourceHoldStatus_Partial),
		string(SecurityUnifiedGroupSourceHoldStatus_Removing),
	}
}

func (s *SecurityUnifiedGroupSourceHoldStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityUnifiedGroupSourceHoldStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityUnifiedGroupSourceHoldStatus(input string) (*SecurityUnifiedGroupSourceHoldStatus, error) {
	vals := map[string]SecurityUnifiedGroupSourceHoldStatus{
		"applied":    SecurityUnifiedGroupSourceHoldStatus_Applied,
		"applying":   SecurityUnifiedGroupSourceHoldStatus_Applying,
		"notapplied": SecurityUnifiedGroupSourceHoldStatus_NotApplied,
		"partial":    SecurityUnifiedGroupSourceHoldStatus_Partial,
		"removing":   SecurityUnifiedGroupSourceHoldStatus_Removing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUnifiedGroupSourceHoldStatus(input)
	return &out, nil
}

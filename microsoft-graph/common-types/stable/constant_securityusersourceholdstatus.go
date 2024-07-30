package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUserSourceHoldStatus string

const (
	SecurityUserSourceHoldStatus_Applied    SecurityUserSourceHoldStatus = "applied"
	SecurityUserSourceHoldStatus_Applying   SecurityUserSourceHoldStatus = "applying"
	SecurityUserSourceHoldStatus_NotApplied SecurityUserSourceHoldStatus = "notApplied"
	SecurityUserSourceHoldStatus_Partial    SecurityUserSourceHoldStatus = "partial"
	SecurityUserSourceHoldStatus_Removing   SecurityUserSourceHoldStatus = "removing"
)

func PossibleValuesForSecurityUserSourceHoldStatus() []string {
	return []string{
		string(SecurityUserSourceHoldStatus_Applied),
		string(SecurityUserSourceHoldStatus_Applying),
		string(SecurityUserSourceHoldStatus_NotApplied),
		string(SecurityUserSourceHoldStatus_Partial),
		string(SecurityUserSourceHoldStatus_Removing),
	}
}

func (s *SecurityUserSourceHoldStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityUserSourceHoldStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityUserSourceHoldStatus(input string) (*SecurityUserSourceHoldStatus, error) {
	vals := map[string]SecurityUserSourceHoldStatus{
		"applied":    SecurityUserSourceHoldStatus_Applied,
		"applying":   SecurityUserSourceHoldStatus_Applying,
		"notapplied": SecurityUserSourceHoldStatus_NotApplied,
		"partial":    SecurityUserSourceHoldStatus_Partial,
		"removing":   SecurityUserSourceHoldStatus_Removing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUserSourceHoldStatus(input)
	return &out, nil
}

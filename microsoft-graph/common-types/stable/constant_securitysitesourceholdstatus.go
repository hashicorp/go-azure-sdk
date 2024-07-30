package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySiteSourceHoldStatus string

const (
	SecuritySiteSourceHoldStatus_Applied    SecuritySiteSourceHoldStatus = "applied"
	SecuritySiteSourceHoldStatus_Applying   SecuritySiteSourceHoldStatus = "applying"
	SecuritySiteSourceHoldStatus_NotApplied SecuritySiteSourceHoldStatus = "notApplied"
	SecuritySiteSourceHoldStatus_Partial    SecuritySiteSourceHoldStatus = "partial"
	SecuritySiteSourceHoldStatus_Removing   SecuritySiteSourceHoldStatus = "removing"
)

func PossibleValuesForSecuritySiteSourceHoldStatus() []string {
	return []string{
		string(SecuritySiteSourceHoldStatus_Applied),
		string(SecuritySiteSourceHoldStatus_Applying),
		string(SecuritySiteSourceHoldStatus_NotApplied),
		string(SecuritySiteSourceHoldStatus_Partial),
		string(SecuritySiteSourceHoldStatus_Removing),
	}
}

func (s *SecuritySiteSourceHoldStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySiteSourceHoldStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySiteSourceHoldStatus(input string) (*SecuritySiteSourceHoldStatus, error) {
	vals := map[string]SecuritySiteSourceHoldStatus{
		"applied":    SecuritySiteSourceHoldStatus_Applied,
		"applying":   SecuritySiteSourceHoldStatus_Applying,
		"notapplied": SecuritySiteSourceHoldStatus_NotApplied,
		"partial":    SecuritySiteSourceHoldStatus_Partial,
		"removing":   SecuritySiteSourceHoldStatus_Removing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySiteSourceHoldStatus(input)
	return &out, nil
}

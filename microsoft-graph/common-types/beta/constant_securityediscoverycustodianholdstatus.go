package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryCustodianHoldStatus string

const (
	SecurityEdiscoveryCustodianHoldStatus_Applied    SecurityEdiscoveryCustodianHoldStatus = "applied"
	SecurityEdiscoveryCustodianHoldStatus_Applying   SecurityEdiscoveryCustodianHoldStatus = "applying"
	SecurityEdiscoveryCustodianHoldStatus_NotApplied SecurityEdiscoveryCustodianHoldStatus = "notApplied"
	SecurityEdiscoveryCustodianHoldStatus_Partial    SecurityEdiscoveryCustodianHoldStatus = "partial"
	SecurityEdiscoveryCustodianHoldStatus_Removing   SecurityEdiscoveryCustodianHoldStatus = "removing"
)

func PossibleValuesForSecurityEdiscoveryCustodianHoldStatus() []string {
	return []string{
		string(SecurityEdiscoveryCustodianHoldStatus_Applied),
		string(SecurityEdiscoveryCustodianHoldStatus_Applying),
		string(SecurityEdiscoveryCustodianHoldStatus_NotApplied),
		string(SecurityEdiscoveryCustodianHoldStatus_Partial),
		string(SecurityEdiscoveryCustodianHoldStatus_Removing),
	}
}

func (s *SecurityEdiscoveryCustodianHoldStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryCustodianHoldStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryCustodianHoldStatus(input string) (*SecurityEdiscoveryCustodianHoldStatus, error) {
	vals := map[string]SecurityEdiscoveryCustodianHoldStatus{
		"applied":    SecurityEdiscoveryCustodianHoldStatus_Applied,
		"applying":   SecurityEdiscoveryCustodianHoldStatus_Applying,
		"notapplied": SecurityEdiscoveryCustodianHoldStatus_NotApplied,
		"partial":    SecurityEdiscoveryCustodianHoldStatus_Partial,
		"removing":   SecurityEdiscoveryCustodianHoldStatus_Removing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryCustodianHoldStatus(input)
	return &out, nil
}

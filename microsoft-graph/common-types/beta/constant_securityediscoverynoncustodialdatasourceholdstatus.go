package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryNoncustodialDataSourceHoldStatus string

const (
	SecurityEdiscoveryNoncustodialDataSourceHoldStatus_Applied    SecurityEdiscoveryNoncustodialDataSourceHoldStatus = "applied"
	SecurityEdiscoveryNoncustodialDataSourceHoldStatus_Applying   SecurityEdiscoveryNoncustodialDataSourceHoldStatus = "applying"
	SecurityEdiscoveryNoncustodialDataSourceHoldStatus_NotApplied SecurityEdiscoveryNoncustodialDataSourceHoldStatus = "notApplied"
	SecurityEdiscoveryNoncustodialDataSourceHoldStatus_Partial    SecurityEdiscoveryNoncustodialDataSourceHoldStatus = "partial"
	SecurityEdiscoveryNoncustodialDataSourceHoldStatus_Removing   SecurityEdiscoveryNoncustodialDataSourceHoldStatus = "removing"
)

func PossibleValuesForSecurityEdiscoveryNoncustodialDataSourceHoldStatus() []string {
	return []string{
		string(SecurityEdiscoveryNoncustodialDataSourceHoldStatus_Applied),
		string(SecurityEdiscoveryNoncustodialDataSourceHoldStatus_Applying),
		string(SecurityEdiscoveryNoncustodialDataSourceHoldStatus_NotApplied),
		string(SecurityEdiscoveryNoncustodialDataSourceHoldStatus_Partial),
		string(SecurityEdiscoveryNoncustodialDataSourceHoldStatus_Removing),
	}
}

func (s *SecurityEdiscoveryNoncustodialDataSourceHoldStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryNoncustodialDataSourceHoldStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryNoncustodialDataSourceHoldStatus(input string) (*SecurityEdiscoveryNoncustodialDataSourceHoldStatus, error) {
	vals := map[string]SecurityEdiscoveryNoncustodialDataSourceHoldStatus{
		"applied":    SecurityEdiscoveryNoncustodialDataSourceHoldStatus_Applied,
		"applying":   SecurityEdiscoveryNoncustodialDataSourceHoldStatus_Applying,
		"notapplied": SecurityEdiscoveryNoncustodialDataSourceHoldStatus_NotApplied,
		"partial":    SecurityEdiscoveryNoncustodialDataSourceHoldStatus_Partial,
		"removing":   SecurityEdiscoveryNoncustodialDataSourceHoldStatus_Removing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryNoncustodialDataSourceHoldStatus(input)
	return &out, nil
}

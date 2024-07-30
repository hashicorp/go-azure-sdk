package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDataSourceContainerHoldStatus string

const (
	SecurityDataSourceContainerHoldStatus_Applied    SecurityDataSourceContainerHoldStatus = "applied"
	SecurityDataSourceContainerHoldStatus_Applying   SecurityDataSourceContainerHoldStatus = "applying"
	SecurityDataSourceContainerHoldStatus_NotApplied SecurityDataSourceContainerHoldStatus = "notApplied"
	SecurityDataSourceContainerHoldStatus_Partial    SecurityDataSourceContainerHoldStatus = "partial"
	SecurityDataSourceContainerHoldStatus_Removing   SecurityDataSourceContainerHoldStatus = "removing"
)

func PossibleValuesForSecurityDataSourceContainerHoldStatus() []string {
	return []string{
		string(SecurityDataSourceContainerHoldStatus_Applied),
		string(SecurityDataSourceContainerHoldStatus_Applying),
		string(SecurityDataSourceContainerHoldStatus_NotApplied),
		string(SecurityDataSourceContainerHoldStatus_Partial),
		string(SecurityDataSourceContainerHoldStatus_Removing),
	}
}

func (s *SecurityDataSourceContainerHoldStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDataSourceContainerHoldStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDataSourceContainerHoldStatus(input string) (*SecurityDataSourceContainerHoldStatus, error) {
	vals := map[string]SecurityDataSourceContainerHoldStatus{
		"applied":    SecurityDataSourceContainerHoldStatus_Applied,
		"applying":   SecurityDataSourceContainerHoldStatus_Applying,
		"notapplied": SecurityDataSourceContainerHoldStatus_NotApplied,
		"partial":    SecurityDataSourceContainerHoldStatus_Partial,
		"removing":   SecurityDataSourceContainerHoldStatus_Removing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDataSourceContainerHoldStatus(input)
	return &out, nil
}

package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryDataSourceContainerHoldStatus string

const (
	EdiscoveryDataSourceContainerHoldStatus_Applied    EdiscoveryDataSourceContainerHoldStatus = "applied"
	EdiscoveryDataSourceContainerHoldStatus_Applying   EdiscoveryDataSourceContainerHoldStatus = "applying"
	EdiscoveryDataSourceContainerHoldStatus_NotApplied EdiscoveryDataSourceContainerHoldStatus = "notApplied"
	EdiscoveryDataSourceContainerHoldStatus_Partial    EdiscoveryDataSourceContainerHoldStatus = "partial"
	EdiscoveryDataSourceContainerHoldStatus_Removing   EdiscoveryDataSourceContainerHoldStatus = "removing"
)

func PossibleValuesForEdiscoveryDataSourceContainerHoldStatus() []string {
	return []string{
		string(EdiscoveryDataSourceContainerHoldStatus_Applied),
		string(EdiscoveryDataSourceContainerHoldStatus_Applying),
		string(EdiscoveryDataSourceContainerHoldStatus_NotApplied),
		string(EdiscoveryDataSourceContainerHoldStatus_Partial),
		string(EdiscoveryDataSourceContainerHoldStatus_Removing),
	}
}

func (s *EdiscoveryDataSourceContainerHoldStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryDataSourceContainerHoldStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryDataSourceContainerHoldStatus(input string) (*EdiscoveryDataSourceContainerHoldStatus, error) {
	vals := map[string]EdiscoveryDataSourceContainerHoldStatus{
		"applied":    EdiscoveryDataSourceContainerHoldStatus_Applied,
		"applying":   EdiscoveryDataSourceContainerHoldStatus_Applying,
		"notapplied": EdiscoveryDataSourceContainerHoldStatus_NotApplied,
		"partial":    EdiscoveryDataSourceContainerHoldStatus_Partial,
		"removing":   EdiscoveryDataSourceContainerHoldStatus_Removing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryDataSourceContainerHoldStatus(input)
	return &out, nil
}

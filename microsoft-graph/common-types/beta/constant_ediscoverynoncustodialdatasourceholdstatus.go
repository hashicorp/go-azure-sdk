package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryNoncustodialDataSourceHoldStatus string

const (
	EdiscoveryNoncustodialDataSourceHoldStatus_Applied    EdiscoveryNoncustodialDataSourceHoldStatus = "applied"
	EdiscoveryNoncustodialDataSourceHoldStatus_Applying   EdiscoveryNoncustodialDataSourceHoldStatus = "applying"
	EdiscoveryNoncustodialDataSourceHoldStatus_NotApplied EdiscoveryNoncustodialDataSourceHoldStatus = "notApplied"
	EdiscoveryNoncustodialDataSourceHoldStatus_Partial    EdiscoveryNoncustodialDataSourceHoldStatus = "partial"
	EdiscoveryNoncustodialDataSourceHoldStatus_Removing   EdiscoveryNoncustodialDataSourceHoldStatus = "removing"
)

func PossibleValuesForEdiscoveryNoncustodialDataSourceHoldStatus() []string {
	return []string{
		string(EdiscoveryNoncustodialDataSourceHoldStatus_Applied),
		string(EdiscoveryNoncustodialDataSourceHoldStatus_Applying),
		string(EdiscoveryNoncustodialDataSourceHoldStatus_NotApplied),
		string(EdiscoveryNoncustodialDataSourceHoldStatus_Partial),
		string(EdiscoveryNoncustodialDataSourceHoldStatus_Removing),
	}
}

func (s *EdiscoveryNoncustodialDataSourceHoldStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryNoncustodialDataSourceHoldStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryNoncustodialDataSourceHoldStatus(input string) (*EdiscoveryNoncustodialDataSourceHoldStatus, error) {
	vals := map[string]EdiscoveryNoncustodialDataSourceHoldStatus{
		"applied":    EdiscoveryNoncustodialDataSourceHoldStatus_Applied,
		"applying":   EdiscoveryNoncustodialDataSourceHoldStatus_Applying,
		"notapplied": EdiscoveryNoncustodialDataSourceHoldStatus_NotApplied,
		"partial":    EdiscoveryNoncustodialDataSourceHoldStatus_Partial,
		"removing":   EdiscoveryNoncustodialDataSourceHoldStatus_Removing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryNoncustodialDataSourceHoldStatus(input)
	return &out, nil
}

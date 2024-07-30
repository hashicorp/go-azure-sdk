package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnmanagedDeviceDiscoveryTaskCategory string

const (
	UnmanagedDeviceDiscoveryTaskCategory_AdvancedThreatProtection UnmanagedDeviceDiscoveryTaskCategory = "advancedThreatProtection"
	UnmanagedDeviceDiscoveryTaskCategory_Unknown                  UnmanagedDeviceDiscoveryTaskCategory = "unknown"
)

func PossibleValuesForUnmanagedDeviceDiscoveryTaskCategory() []string {
	return []string{
		string(UnmanagedDeviceDiscoveryTaskCategory_AdvancedThreatProtection),
		string(UnmanagedDeviceDiscoveryTaskCategory_Unknown),
	}
}

func (s *UnmanagedDeviceDiscoveryTaskCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUnmanagedDeviceDiscoveryTaskCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUnmanagedDeviceDiscoveryTaskCategory(input string) (*UnmanagedDeviceDiscoveryTaskCategory, error) {
	vals := map[string]UnmanagedDeviceDiscoveryTaskCategory{
		"advancedthreatprotection": UnmanagedDeviceDiscoveryTaskCategory_AdvancedThreatProtection,
		"unknown":                  UnmanagedDeviceDiscoveryTaskCategory_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnmanagedDeviceDiscoveryTaskCategory(input)
	return &out, nil
}

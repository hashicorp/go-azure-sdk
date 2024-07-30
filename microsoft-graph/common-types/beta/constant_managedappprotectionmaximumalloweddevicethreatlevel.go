package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionMaximumAllowedDeviceThreatLevel string

const (
	ManagedAppProtectionMaximumAllowedDeviceThreatLevel_High          ManagedAppProtectionMaximumAllowedDeviceThreatLevel = "high"
	ManagedAppProtectionMaximumAllowedDeviceThreatLevel_Low           ManagedAppProtectionMaximumAllowedDeviceThreatLevel = "low"
	ManagedAppProtectionMaximumAllowedDeviceThreatLevel_Medium        ManagedAppProtectionMaximumAllowedDeviceThreatLevel = "medium"
	ManagedAppProtectionMaximumAllowedDeviceThreatLevel_NotConfigured ManagedAppProtectionMaximumAllowedDeviceThreatLevel = "notConfigured"
	ManagedAppProtectionMaximumAllowedDeviceThreatLevel_Secured       ManagedAppProtectionMaximumAllowedDeviceThreatLevel = "secured"
)

func PossibleValuesForManagedAppProtectionMaximumAllowedDeviceThreatLevel() []string {
	return []string{
		string(ManagedAppProtectionMaximumAllowedDeviceThreatLevel_High),
		string(ManagedAppProtectionMaximumAllowedDeviceThreatLevel_Low),
		string(ManagedAppProtectionMaximumAllowedDeviceThreatLevel_Medium),
		string(ManagedAppProtectionMaximumAllowedDeviceThreatLevel_NotConfigured),
		string(ManagedAppProtectionMaximumAllowedDeviceThreatLevel_Secured),
	}
}

func (s *ManagedAppProtectionMaximumAllowedDeviceThreatLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppProtectionMaximumAllowedDeviceThreatLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppProtectionMaximumAllowedDeviceThreatLevel(input string) (*ManagedAppProtectionMaximumAllowedDeviceThreatLevel, error) {
	vals := map[string]ManagedAppProtectionMaximumAllowedDeviceThreatLevel{
		"high":          ManagedAppProtectionMaximumAllowedDeviceThreatLevel_High,
		"low":           ManagedAppProtectionMaximumAllowedDeviceThreatLevel_Low,
		"medium":        ManagedAppProtectionMaximumAllowedDeviceThreatLevel_Medium,
		"notconfigured": ManagedAppProtectionMaximumAllowedDeviceThreatLevel_NotConfigured,
		"secured":       ManagedAppProtectionMaximumAllowedDeviceThreatLevel_Secured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionMaximumAllowedDeviceThreatLevel(input)
	return &out, nil
}

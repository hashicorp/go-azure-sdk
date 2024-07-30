package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel string

const (
	TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel_High          TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel = "high"
	TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel_Low           TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel = "low"
	TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel_Medium        TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel = "medium"
	TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel_NotConfigured TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel = "notConfigured"
	TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel_Secured       TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel = "secured"
)

func PossibleValuesForTargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel() []string {
	return []string{
		string(TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel_High),
		string(TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel_Low),
		string(TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel_Medium),
		string(TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel_NotConfigured),
		string(TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel_Secured),
	}
}

func (s *TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel(input string) (*TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel, error) {
	vals := map[string]TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel{
		"high":          TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel_High,
		"low":           TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel_Low,
		"medium":        TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel_Medium,
		"notconfigured": TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel_NotConfigured,
		"secured":       TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel_Secured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel(input)
	return &out, nil
}

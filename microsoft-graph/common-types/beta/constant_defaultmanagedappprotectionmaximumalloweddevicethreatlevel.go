package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel string

const (
	DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel_High          DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel = "high"
	DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel_Low           DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel = "low"
	DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel_Medium        DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel = "medium"
	DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel_NotConfigured DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel = "notConfigured"
	DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel_Secured       DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel = "secured"
)

func PossibleValuesForDefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel() []string {
	return []string{
		string(DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel_High),
		string(DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel_Low),
		string(DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel_Medium),
		string(DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel_NotConfigured),
		string(DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel_Secured),
	}
}

func (s *DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel(input string) (*DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel, error) {
	vals := map[string]DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel{
		"high":          DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel_High,
		"low":           DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel_Low,
		"medium":        DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel_Medium,
		"notconfigured": DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel_NotConfigured,
		"secured":       DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel_Secured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel(input)
	return &out, nil
}

package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel string

const (
	WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel_High          WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel = "high"
	WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel_Low           WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel = "low"
	WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel_Medium        WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel = "medium"
	WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel_NotConfigured WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel = "notConfigured"
	WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel_Secured       WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel = "secured"
)

func PossibleValuesForWindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel() []string {
	return []string{
		string(WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel_High),
		string(WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel_Low),
		string(WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel_Medium),
		string(WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel_NotConfigured),
		string(WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel_Secured),
	}
}

func (s *WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel(input string) (*WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel, error) {
	vals := map[string]WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel{
		"high":          WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel_High,
		"low":           WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel_Low,
		"medium":        WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel_Medium,
		"notconfigured": WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel_NotConfigured,
		"secured":       WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel_Secured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel(input)
	return &out, nil
}

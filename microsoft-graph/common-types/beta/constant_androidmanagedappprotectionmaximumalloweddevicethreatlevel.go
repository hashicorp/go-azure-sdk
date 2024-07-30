package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel string

const (
	AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel_High          AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel = "high"
	AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel_Low           AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel = "low"
	AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel_Medium        AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel = "medium"
	AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel_NotConfigured AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel = "notConfigured"
	AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel_Secured       AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel = "secured"
)

func PossibleValuesForAndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel() []string {
	return []string{
		string(AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel_High),
		string(AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel_Low),
		string(AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel_Medium),
		string(AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel_NotConfigured),
		string(AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel_Secured),
	}
}

func (s *AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel(input string) (*AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel, error) {
	vals := map[string]AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel{
		"high":          AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel_High,
		"low":           AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel_Low,
		"medium":        AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel_Medium,
		"notconfigured": AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel_NotConfigured,
		"secured":       AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel_Secured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionMaximumAllowedDeviceThreatLevel(input)
	return &out, nil
}

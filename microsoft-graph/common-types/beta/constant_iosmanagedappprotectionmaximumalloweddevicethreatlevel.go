package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionMaximumAllowedDeviceThreatLevel string

const (
	IosManagedAppProtectionMaximumAllowedDeviceThreatLevel_High          IosManagedAppProtectionMaximumAllowedDeviceThreatLevel = "high"
	IosManagedAppProtectionMaximumAllowedDeviceThreatLevel_Low           IosManagedAppProtectionMaximumAllowedDeviceThreatLevel = "low"
	IosManagedAppProtectionMaximumAllowedDeviceThreatLevel_Medium        IosManagedAppProtectionMaximumAllowedDeviceThreatLevel = "medium"
	IosManagedAppProtectionMaximumAllowedDeviceThreatLevel_NotConfigured IosManagedAppProtectionMaximumAllowedDeviceThreatLevel = "notConfigured"
	IosManagedAppProtectionMaximumAllowedDeviceThreatLevel_Secured       IosManagedAppProtectionMaximumAllowedDeviceThreatLevel = "secured"
)

func PossibleValuesForIosManagedAppProtectionMaximumAllowedDeviceThreatLevel() []string {
	return []string{
		string(IosManagedAppProtectionMaximumAllowedDeviceThreatLevel_High),
		string(IosManagedAppProtectionMaximumAllowedDeviceThreatLevel_Low),
		string(IosManagedAppProtectionMaximumAllowedDeviceThreatLevel_Medium),
		string(IosManagedAppProtectionMaximumAllowedDeviceThreatLevel_NotConfigured),
		string(IosManagedAppProtectionMaximumAllowedDeviceThreatLevel_Secured),
	}
}

func (s *IosManagedAppProtectionMaximumAllowedDeviceThreatLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosManagedAppProtectionMaximumAllowedDeviceThreatLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosManagedAppProtectionMaximumAllowedDeviceThreatLevel(input string) (*IosManagedAppProtectionMaximumAllowedDeviceThreatLevel, error) {
	vals := map[string]IosManagedAppProtectionMaximumAllowedDeviceThreatLevel{
		"high":          IosManagedAppProtectionMaximumAllowedDeviceThreatLevel_High,
		"low":           IosManagedAppProtectionMaximumAllowedDeviceThreatLevel_Low,
		"medium":        IosManagedAppProtectionMaximumAllowedDeviceThreatLevel_Medium,
		"notconfigured": IosManagedAppProtectionMaximumAllowedDeviceThreatLevel_NotConfigured,
		"secured":       IosManagedAppProtectionMaximumAllowedDeviceThreatLevel_Secured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionMaximumAllowedDeviceThreatLevel(input)
	return &out, nil
}

package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationAppLockerApplicationControl string

const (
	Windows10EndpointProtectionConfigurationAppLockerApplicationControl_AuditComponentsAndStoreApps              Windows10EndpointProtectionConfigurationAppLockerApplicationControl = "auditComponentsAndStoreApps"
	Windows10EndpointProtectionConfigurationAppLockerApplicationControl_AuditComponentsStoreAppsAndSmartlocker   Windows10EndpointProtectionConfigurationAppLockerApplicationControl = "auditComponentsStoreAppsAndSmartlocker"
	Windows10EndpointProtectionConfigurationAppLockerApplicationControl_EnforceComponentsAndStoreApps            Windows10EndpointProtectionConfigurationAppLockerApplicationControl = "enforceComponentsAndStoreApps"
	Windows10EndpointProtectionConfigurationAppLockerApplicationControl_EnforceComponentsStoreAppsAndSmartlocker Windows10EndpointProtectionConfigurationAppLockerApplicationControl = "enforceComponentsStoreAppsAndSmartlocker"
	Windows10EndpointProtectionConfigurationAppLockerApplicationControl_NotConfigured                            Windows10EndpointProtectionConfigurationAppLockerApplicationControl = "notConfigured"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationAppLockerApplicationControl() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationAppLockerApplicationControl_AuditComponentsAndStoreApps),
		string(Windows10EndpointProtectionConfigurationAppLockerApplicationControl_AuditComponentsStoreAppsAndSmartlocker),
		string(Windows10EndpointProtectionConfigurationAppLockerApplicationControl_EnforceComponentsAndStoreApps),
		string(Windows10EndpointProtectionConfigurationAppLockerApplicationControl_EnforceComponentsStoreAppsAndSmartlocker),
		string(Windows10EndpointProtectionConfigurationAppLockerApplicationControl_NotConfigured),
	}
}

func (s *Windows10EndpointProtectionConfigurationAppLockerApplicationControl) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationAppLockerApplicationControl(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationAppLockerApplicationControl(input string) (*Windows10EndpointProtectionConfigurationAppLockerApplicationControl, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationAppLockerApplicationControl{
		"auditcomponentsandstoreapps":              Windows10EndpointProtectionConfigurationAppLockerApplicationControl_AuditComponentsAndStoreApps,
		"auditcomponentsstoreappsandsmartlocker":   Windows10EndpointProtectionConfigurationAppLockerApplicationControl_AuditComponentsStoreAppsAndSmartlocker,
		"enforcecomponentsandstoreapps":            Windows10EndpointProtectionConfigurationAppLockerApplicationControl_EnforceComponentsAndStoreApps,
		"enforcecomponentsstoreappsandsmartlocker": Windows10EndpointProtectionConfigurationAppLockerApplicationControl_EnforceComponentsStoreAppsAndSmartlocker,
		"notconfigured":                            Windows10EndpointProtectionConfigurationAppLockerApplicationControl_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationAppLockerApplicationControl(input)
	return &out, nil
}

package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationAppLockerApplicationControl string

const (
	Windows10EndpointProtectionConfigurationAppLockerApplicationControlauditComponentsAndStoreApps              Windows10EndpointProtectionConfigurationAppLockerApplicationControl = "AuditComponentsAndStoreApps"
	Windows10EndpointProtectionConfigurationAppLockerApplicationControlauditComponentsStoreAppsAndSmartlocker   Windows10EndpointProtectionConfigurationAppLockerApplicationControl = "AuditComponentsStoreAppsAndSmartlocker"
	Windows10EndpointProtectionConfigurationAppLockerApplicationControlenforceComponentsAndStoreApps            Windows10EndpointProtectionConfigurationAppLockerApplicationControl = "EnforceComponentsAndStoreApps"
	Windows10EndpointProtectionConfigurationAppLockerApplicationControlenforceComponentsStoreAppsAndSmartlocker Windows10EndpointProtectionConfigurationAppLockerApplicationControl = "EnforceComponentsStoreAppsAndSmartlocker"
	Windows10EndpointProtectionConfigurationAppLockerApplicationControlnotConfigured                            Windows10EndpointProtectionConfigurationAppLockerApplicationControl = "NotConfigured"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationAppLockerApplicationControl() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationAppLockerApplicationControlauditComponentsAndStoreApps),
		string(Windows10EndpointProtectionConfigurationAppLockerApplicationControlauditComponentsStoreAppsAndSmartlocker),
		string(Windows10EndpointProtectionConfigurationAppLockerApplicationControlenforceComponentsAndStoreApps),
		string(Windows10EndpointProtectionConfigurationAppLockerApplicationControlenforceComponentsStoreAppsAndSmartlocker),
		string(Windows10EndpointProtectionConfigurationAppLockerApplicationControlnotConfigured),
	}
}

func parseWindows10EndpointProtectionConfigurationAppLockerApplicationControl(input string) (*Windows10EndpointProtectionConfigurationAppLockerApplicationControl, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationAppLockerApplicationControl{
		"auditcomponentsandstoreapps":              Windows10EndpointProtectionConfigurationAppLockerApplicationControlauditComponentsAndStoreApps,
		"auditcomponentsstoreappsandsmartlocker":   Windows10EndpointProtectionConfigurationAppLockerApplicationControlauditComponentsStoreAppsAndSmartlocker,
		"enforcecomponentsandstoreapps":            Windows10EndpointProtectionConfigurationAppLockerApplicationControlenforceComponentsAndStoreApps,
		"enforcecomponentsstoreappsandsmartlocker": Windows10EndpointProtectionConfigurationAppLockerApplicationControlenforceComponentsStoreAppsAndSmartlocker,
		"notconfigured":                            Windows10EndpointProtectionConfigurationAppLockerApplicationControlnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationAppLockerApplicationControl(input)
	return &out, nil
}

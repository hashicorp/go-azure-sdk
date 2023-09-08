package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionDialerRestrictionLevel string

const (
	IosManagedAppProtectionDialerRestrictionLevelallApps     IosManagedAppProtectionDialerRestrictionLevel = "AllApps"
	IosManagedAppProtectionDialerRestrictionLevelblocked     IosManagedAppProtectionDialerRestrictionLevel = "Blocked"
	IosManagedAppProtectionDialerRestrictionLevelcustomApp   IosManagedAppProtectionDialerRestrictionLevel = "CustomApp"
	IosManagedAppProtectionDialerRestrictionLevelmanagedApps IosManagedAppProtectionDialerRestrictionLevel = "ManagedApps"
)

func PossibleValuesForIosManagedAppProtectionDialerRestrictionLevel() []string {
	return []string{
		string(IosManagedAppProtectionDialerRestrictionLevelallApps),
		string(IosManagedAppProtectionDialerRestrictionLevelblocked),
		string(IosManagedAppProtectionDialerRestrictionLevelcustomApp),
		string(IosManagedAppProtectionDialerRestrictionLevelmanagedApps),
	}
}

func parseIosManagedAppProtectionDialerRestrictionLevel(input string) (*IosManagedAppProtectionDialerRestrictionLevel, error) {
	vals := map[string]IosManagedAppProtectionDialerRestrictionLevel{
		"allapps":     IosManagedAppProtectionDialerRestrictionLevelallApps,
		"blocked":     IosManagedAppProtectionDialerRestrictionLevelblocked,
		"customapp":   IosManagedAppProtectionDialerRestrictionLevelcustomApp,
		"managedapps": IosManagedAppProtectionDialerRestrictionLevelmanagedApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionDialerRestrictionLevel(input)
	return &out, nil
}

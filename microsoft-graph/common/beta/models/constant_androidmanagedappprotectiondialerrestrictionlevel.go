package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionDialerRestrictionLevel string

const (
	AndroidManagedAppProtectionDialerRestrictionLevelallApps     AndroidManagedAppProtectionDialerRestrictionLevel = "AllApps"
	AndroidManagedAppProtectionDialerRestrictionLevelblocked     AndroidManagedAppProtectionDialerRestrictionLevel = "Blocked"
	AndroidManagedAppProtectionDialerRestrictionLevelcustomApp   AndroidManagedAppProtectionDialerRestrictionLevel = "CustomApp"
	AndroidManagedAppProtectionDialerRestrictionLevelmanagedApps AndroidManagedAppProtectionDialerRestrictionLevel = "ManagedApps"
)

func PossibleValuesForAndroidManagedAppProtectionDialerRestrictionLevel() []string {
	return []string{
		string(AndroidManagedAppProtectionDialerRestrictionLevelallApps),
		string(AndroidManagedAppProtectionDialerRestrictionLevelblocked),
		string(AndroidManagedAppProtectionDialerRestrictionLevelcustomApp),
		string(AndroidManagedAppProtectionDialerRestrictionLevelmanagedApps),
	}
}

func parseAndroidManagedAppProtectionDialerRestrictionLevel(input string) (*AndroidManagedAppProtectionDialerRestrictionLevel, error) {
	vals := map[string]AndroidManagedAppProtectionDialerRestrictionLevel{
		"allapps":     AndroidManagedAppProtectionDialerRestrictionLevelallApps,
		"blocked":     AndroidManagedAppProtectionDialerRestrictionLevelblocked,
		"customapp":   AndroidManagedAppProtectionDialerRestrictionLevelcustomApp,
		"managedapps": AndroidManagedAppProtectionDialerRestrictionLevelmanagedApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionDialerRestrictionLevel(input)
	return &out, nil
}

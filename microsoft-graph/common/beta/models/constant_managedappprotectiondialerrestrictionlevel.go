package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionDialerRestrictionLevel string

const (
	ManagedAppProtectionDialerRestrictionLevelallApps     ManagedAppProtectionDialerRestrictionLevel = "AllApps"
	ManagedAppProtectionDialerRestrictionLevelblocked     ManagedAppProtectionDialerRestrictionLevel = "Blocked"
	ManagedAppProtectionDialerRestrictionLevelcustomApp   ManagedAppProtectionDialerRestrictionLevel = "CustomApp"
	ManagedAppProtectionDialerRestrictionLevelmanagedApps ManagedAppProtectionDialerRestrictionLevel = "ManagedApps"
)

func PossibleValuesForManagedAppProtectionDialerRestrictionLevel() []string {
	return []string{
		string(ManagedAppProtectionDialerRestrictionLevelallApps),
		string(ManagedAppProtectionDialerRestrictionLevelblocked),
		string(ManagedAppProtectionDialerRestrictionLevelcustomApp),
		string(ManagedAppProtectionDialerRestrictionLevelmanagedApps),
	}
}

func parseManagedAppProtectionDialerRestrictionLevel(input string) (*ManagedAppProtectionDialerRestrictionLevel, error) {
	vals := map[string]ManagedAppProtectionDialerRestrictionLevel{
		"allapps":     ManagedAppProtectionDialerRestrictionLevelallApps,
		"blocked":     ManagedAppProtectionDialerRestrictionLevelblocked,
		"customapp":   ManagedAppProtectionDialerRestrictionLevelcustomApp,
		"managedapps": ManagedAppProtectionDialerRestrictionLevelmanagedApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionDialerRestrictionLevel(input)
	return &out, nil
}

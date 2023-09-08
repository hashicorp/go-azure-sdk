package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionDialerRestrictionLevel string

const (
	DefaultManagedAppProtectionDialerRestrictionLevelallApps     DefaultManagedAppProtectionDialerRestrictionLevel = "AllApps"
	DefaultManagedAppProtectionDialerRestrictionLevelblocked     DefaultManagedAppProtectionDialerRestrictionLevel = "Blocked"
	DefaultManagedAppProtectionDialerRestrictionLevelcustomApp   DefaultManagedAppProtectionDialerRestrictionLevel = "CustomApp"
	DefaultManagedAppProtectionDialerRestrictionLevelmanagedApps DefaultManagedAppProtectionDialerRestrictionLevel = "ManagedApps"
)

func PossibleValuesForDefaultManagedAppProtectionDialerRestrictionLevel() []string {
	return []string{
		string(DefaultManagedAppProtectionDialerRestrictionLevelallApps),
		string(DefaultManagedAppProtectionDialerRestrictionLevelblocked),
		string(DefaultManagedAppProtectionDialerRestrictionLevelcustomApp),
		string(DefaultManagedAppProtectionDialerRestrictionLevelmanagedApps),
	}
}

func parseDefaultManagedAppProtectionDialerRestrictionLevel(input string) (*DefaultManagedAppProtectionDialerRestrictionLevel, error) {
	vals := map[string]DefaultManagedAppProtectionDialerRestrictionLevel{
		"allapps":     DefaultManagedAppProtectionDialerRestrictionLevelallApps,
		"blocked":     DefaultManagedAppProtectionDialerRestrictionLevelblocked,
		"customapp":   DefaultManagedAppProtectionDialerRestrictionLevelcustomApp,
		"managedapps": DefaultManagedAppProtectionDialerRestrictionLevelmanagedApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionDialerRestrictionLevel(input)
	return &out, nil
}

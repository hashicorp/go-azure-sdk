package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppGroupType string

const (
	AndroidManagedAppProtectionAppGroupTypeallApps              AndroidManagedAppProtectionAppGroupType = "AllApps"
	AndroidManagedAppProtectionAppGroupTypeallCoreMicrosoftApps AndroidManagedAppProtectionAppGroupType = "AllCoreMicrosoftApps"
	AndroidManagedAppProtectionAppGroupTypeallMicrosoftApps     AndroidManagedAppProtectionAppGroupType = "AllMicrosoftApps"
	AndroidManagedAppProtectionAppGroupTypeselectedPublicApps   AndroidManagedAppProtectionAppGroupType = "SelectedPublicApps"
)

func PossibleValuesForAndroidManagedAppProtectionAppGroupType() []string {
	return []string{
		string(AndroidManagedAppProtectionAppGroupTypeallApps),
		string(AndroidManagedAppProtectionAppGroupTypeallCoreMicrosoftApps),
		string(AndroidManagedAppProtectionAppGroupTypeallMicrosoftApps),
		string(AndroidManagedAppProtectionAppGroupTypeselectedPublicApps),
	}
}

func parseAndroidManagedAppProtectionAppGroupType(input string) (*AndroidManagedAppProtectionAppGroupType, error) {
	vals := map[string]AndroidManagedAppProtectionAppGroupType{
		"allapps":              AndroidManagedAppProtectionAppGroupTypeallApps,
		"allcoremicrosoftapps": AndroidManagedAppProtectionAppGroupTypeallCoreMicrosoftApps,
		"allmicrosoftapps":     AndroidManagedAppProtectionAppGroupTypeallMicrosoftApps,
		"selectedpublicapps":   AndroidManagedAppProtectionAppGroupTypeselectedPublicApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppGroupType(input)
	return &out, nil
}

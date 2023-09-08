package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppConfigurationAppGroupType string

const (
	TargetedManagedAppConfigurationAppGroupTypeallApps              TargetedManagedAppConfigurationAppGroupType = "AllApps"
	TargetedManagedAppConfigurationAppGroupTypeallCoreMicrosoftApps TargetedManagedAppConfigurationAppGroupType = "AllCoreMicrosoftApps"
	TargetedManagedAppConfigurationAppGroupTypeallMicrosoftApps     TargetedManagedAppConfigurationAppGroupType = "AllMicrosoftApps"
	TargetedManagedAppConfigurationAppGroupTypeselectedPublicApps   TargetedManagedAppConfigurationAppGroupType = "SelectedPublicApps"
)

func PossibleValuesForTargetedManagedAppConfigurationAppGroupType() []string {
	return []string{
		string(TargetedManagedAppConfigurationAppGroupTypeallApps),
		string(TargetedManagedAppConfigurationAppGroupTypeallCoreMicrosoftApps),
		string(TargetedManagedAppConfigurationAppGroupTypeallMicrosoftApps),
		string(TargetedManagedAppConfigurationAppGroupTypeselectedPublicApps),
	}
}

func parseTargetedManagedAppConfigurationAppGroupType(input string) (*TargetedManagedAppConfigurationAppGroupType, error) {
	vals := map[string]TargetedManagedAppConfigurationAppGroupType{
		"allapps":              TargetedManagedAppConfigurationAppGroupTypeallApps,
		"allcoremicrosoftapps": TargetedManagedAppConfigurationAppGroupTypeallCoreMicrosoftApps,
		"allmicrosoftapps":     TargetedManagedAppConfigurationAppGroupTypeallMicrosoftApps,
		"selectedpublicapps":   TargetedManagedAppConfigurationAppGroupTypeselectedPublicApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppConfigurationAppGroupType(input)
	return &out, nil
}

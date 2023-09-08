package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutOfBoxExperienceSettingsDeviceUsageType string

const (
	OutOfBoxExperienceSettingsDeviceUsageTypeshared     OutOfBoxExperienceSettingsDeviceUsageType = "Shared"
	OutOfBoxExperienceSettingsDeviceUsageTypesingleUser OutOfBoxExperienceSettingsDeviceUsageType = "SingleUser"
)

func PossibleValuesForOutOfBoxExperienceSettingsDeviceUsageType() []string {
	return []string{
		string(OutOfBoxExperienceSettingsDeviceUsageTypeshared),
		string(OutOfBoxExperienceSettingsDeviceUsageTypesingleUser),
	}
}

func parseOutOfBoxExperienceSettingsDeviceUsageType(input string) (*OutOfBoxExperienceSettingsDeviceUsageType, error) {
	vals := map[string]OutOfBoxExperienceSettingsDeviceUsageType{
		"shared":     OutOfBoxExperienceSettingsDeviceUsageTypeshared,
		"singleuser": OutOfBoxExperienceSettingsDeviceUsageTypesingleUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutOfBoxExperienceSettingsDeviceUsageType(input)
	return &out, nil
}

package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingComparisonComparisonResult string

const (
	DeviceManagementSettingComparisonComparisonResultadded    DeviceManagementSettingComparisonComparisonResult = "Added"
	DeviceManagementSettingComparisonComparisonResultequal    DeviceManagementSettingComparisonComparisonResult = "Equal"
	DeviceManagementSettingComparisonComparisonResultnotEqual DeviceManagementSettingComparisonComparisonResult = "NotEqual"
	DeviceManagementSettingComparisonComparisonResultremoved  DeviceManagementSettingComparisonComparisonResult = "Removed"
	DeviceManagementSettingComparisonComparisonResultunknown  DeviceManagementSettingComparisonComparisonResult = "Unknown"
)

func PossibleValuesForDeviceManagementSettingComparisonComparisonResult() []string {
	return []string{
		string(DeviceManagementSettingComparisonComparisonResultadded),
		string(DeviceManagementSettingComparisonComparisonResultequal),
		string(DeviceManagementSettingComparisonComparisonResultnotEqual),
		string(DeviceManagementSettingComparisonComparisonResultremoved),
		string(DeviceManagementSettingComparisonComparisonResultunknown),
	}
}

func parseDeviceManagementSettingComparisonComparisonResult(input string) (*DeviceManagementSettingComparisonComparisonResult, error) {
	vals := map[string]DeviceManagementSettingComparisonComparisonResult{
		"added":    DeviceManagementSettingComparisonComparisonResultadded,
		"equal":    DeviceManagementSettingComparisonComparisonResultequal,
		"notequal": DeviceManagementSettingComparisonComparisonResultnotEqual,
		"removed":  DeviceManagementSettingComparisonComparisonResultremoved,
		"unknown":  DeviceManagementSettingComparisonComparisonResultunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementSettingComparisonComparisonResult(input)
	return &out, nil
}

package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementUserRightsSettingState string

const (
	DeviceManagementUserRightsSettingStateallowed       DeviceManagementUserRightsSettingState = "Allowed"
	DeviceManagementUserRightsSettingStateblocked       DeviceManagementUserRightsSettingState = "Blocked"
	DeviceManagementUserRightsSettingStatenotConfigured DeviceManagementUserRightsSettingState = "NotConfigured"
)

func PossibleValuesForDeviceManagementUserRightsSettingState() []string {
	return []string{
		string(DeviceManagementUserRightsSettingStateallowed),
		string(DeviceManagementUserRightsSettingStateblocked),
		string(DeviceManagementUserRightsSettingStatenotConfigured),
	}
}

func parseDeviceManagementUserRightsSettingState(input string) (*DeviceManagementUserRightsSettingState, error) {
	vals := map[string]DeviceManagementUserRightsSettingState{
		"allowed":       DeviceManagementUserRightsSettingStateallowed,
		"blocked":       DeviceManagementUserRightsSettingStateblocked,
		"notconfigured": DeviceManagementUserRightsSettingStatenotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementUserRightsSettingState(input)
	return &out, nil
}

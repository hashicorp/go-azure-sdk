package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode string

const (
	DeviceManagementConfigurationApplicationSettingApplicabilityDeviceModekiosk DeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode = "Kiosk"
	DeviceManagementConfigurationApplicationSettingApplicabilityDeviceModenone  DeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode = "None"
)

func PossibleValuesForDeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode() []string {
	return []string{
		string(DeviceManagementConfigurationApplicationSettingApplicabilityDeviceModekiosk),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityDeviceModenone),
	}
}

func parseDeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode(input string) (*DeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode, error) {
	vals := map[string]DeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode{
		"kiosk": DeviceManagementConfigurationApplicationSettingApplicabilityDeviceModekiosk,
		"none":  DeviceManagementConfigurationApplicationSettingApplicabilityDeviceModenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode(input)
	return &out, nil
}

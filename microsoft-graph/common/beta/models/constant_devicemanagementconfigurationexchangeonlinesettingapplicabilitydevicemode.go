package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode string

const (
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceModekiosk DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode = "Kiosk"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceModenone  DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode = "None"
)

func PossibleValuesForDeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode() []string {
	return []string{
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceModekiosk),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceModenone),
	}
}

func parseDeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode(input string) (*DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode, error) {
	vals := map[string]DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode{
		"kiosk": DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceModekiosk,
		"none":  DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceModenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode(input)
	return &out, nil
}

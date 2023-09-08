package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies string

const (
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesappleRemoteManagement       DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies = "AppleRemoteManagement"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesconfigManager               DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies = "ConfigManager"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesendpointPrivilegeManagement DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies = "EndpointPrivilegeManagement"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesenrollment                  DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies = "Enrollment"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesexchangeOnline              DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies = "ExchangeOnline"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologieslinuxMdm                    DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies = "LinuxMdm"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesmdm                         DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies = "Mdm"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesmicrosoftSense              DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies = "MicrosoftSense"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesnone                        DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies = "None"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologieswindows10XManagement        DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies = "Windows10XManagement"
)

func PossibleValuesForDeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies() []string {
	return []string{
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesappleRemoteManagement),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesconfigManager),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesendpointPrivilegeManagement),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesenrollment),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesexchangeOnline),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologieslinuxMdm),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesmdm),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesmicrosoftSense),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesnone),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologieswindows10XManagement),
	}
}

func parseDeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies(input string) (*DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies, error) {
	vals := map[string]DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies{
		"appleremotemanagement":       DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesappleRemoteManagement,
		"configmanager":               DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesconfigManager,
		"endpointprivilegemanagement": DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesendpointPrivilegeManagement,
		"enrollment":                  DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesenrollment,
		"exchangeonline":              DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesexchangeOnline,
		"linuxmdm":                    DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologieslinuxMdm,
		"mdm":                         DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesmdm,
		"microsoftsense":              DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesmicrosoftSense,
		"none":                        DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologiesnone,
		"windows10xmanagement":        DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologieswindows10XManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies(input)
	return &out, nil
}

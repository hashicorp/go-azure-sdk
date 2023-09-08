package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies string

const (
	DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesappleRemoteManagement       DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies = "AppleRemoteManagement"
	DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesconfigManager               DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies = "ConfigManager"
	DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesendpointPrivilegeManagement DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies = "EndpointPrivilegeManagement"
	DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesenrollment                  DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies = "Enrollment"
	DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesexchangeOnline              DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies = "ExchangeOnline"
	DeviceManagementConfigurationWindowsSettingApplicabilityTechnologieslinuxMdm                    DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies = "LinuxMdm"
	DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesmdm                         DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies = "Mdm"
	DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesmicrosoftSense              DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies = "MicrosoftSense"
	DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesnone                        DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies = "None"
	DeviceManagementConfigurationWindowsSettingApplicabilityTechnologieswindows10XManagement        DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies = "Windows10XManagement"
)

func PossibleValuesForDeviceManagementConfigurationWindowsSettingApplicabilityTechnologies() []string {
	return []string{
		string(DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesappleRemoteManagement),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesconfigManager),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesendpointPrivilegeManagement),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesenrollment),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesexchangeOnline),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityTechnologieslinuxMdm),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesmdm),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesmicrosoftSense),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesnone),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityTechnologieswindows10XManagement),
	}
}

func parseDeviceManagementConfigurationWindowsSettingApplicabilityTechnologies(input string) (*DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies, error) {
	vals := map[string]DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies{
		"appleremotemanagement":       DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesappleRemoteManagement,
		"configmanager":               DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesconfigManager,
		"endpointprivilegemanagement": DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesendpointPrivilegeManagement,
		"enrollment":                  DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesenrollment,
		"exchangeonline":              DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesexchangeOnline,
		"linuxmdm":                    DeviceManagementConfigurationWindowsSettingApplicabilityTechnologieslinuxMdm,
		"mdm":                         DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesmdm,
		"microsoftsense":              DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesmicrosoftSense,
		"none":                        DeviceManagementConfigurationWindowsSettingApplicabilityTechnologiesnone,
		"windows10xmanagement":        DeviceManagementConfigurationWindowsSettingApplicabilityTechnologieswindows10XManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies(input)
	return &out, nil
}

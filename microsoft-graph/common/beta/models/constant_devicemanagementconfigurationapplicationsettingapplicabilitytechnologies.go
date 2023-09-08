package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies string

const (
	DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesappleRemoteManagement       DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies = "AppleRemoteManagement"
	DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesconfigManager               DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies = "ConfigManager"
	DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesendpointPrivilegeManagement DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies = "EndpointPrivilegeManagement"
	DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesenrollment                  DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies = "Enrollment"
	DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesexchangeOnline              DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies = "ExchangeOnline"
	DeviceManagementConfigurationApplicationSettingApplicabilityTechnologieslinuxMdm                    DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies = "LinuxMdm"
	DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesmdm                         DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies = "Mdm"
	DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesmicrosoftSense              DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies = "MicrosoftSense"
	DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesnone                        DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies = "None"
	DeviceManagementConfigurationApplicationSettingApplicabilityTechnologieswindows10XManagement        DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies = "Windows10XManagement"
)

func PossibleValuesForDeviceManagementConfigurationApplicationSettingApplicabilityTechnologies() []string {
	return []string{
		string(DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesappleRemoteManagement),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesconfigManager),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesendpointPrivilegeManagement),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesenrollment),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesexchangeOnline),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityTechnologieslinuxMdm),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesmdm),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesmicrosoftSense),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesnone),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityTechnologieswindows10XManagement),
	}
}

func parseDeviceManagementConfigurationApplicationSettingApplicabilityTechnologies(input string) (*DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies, error) {
	vals := map[string]DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies{
		"appleremotemanagement":       DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesappleRemoteManagement,
		"configmanager":               DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesconfigManager,
		"endpointprivilegemanagement": DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesendpointPrivilegeManagement,
		"enrollment":                  DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesenrollment,
		"exchangeonline":              DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesexchangeOnline,
		"linuxmdm":                    DeviceManagementConfigurationApplicationSettingApplicabilityTechnologieslinuxMdm,
		"mdm":                         DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesmdm,
		"microsoftsense":              DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesmicrosoftSense,
		"none":                        DeviceManagementConfigurationApplicationSettingApplicabilityTechnologiesnone,
		"windows10xmanagement":        DeviceManagementConfigurationApplicationSettingApplicabilityTechnologieswindows10XManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies(input)
	return &out, nil
}

package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingApplicabilityTechnologies string

const (
	DeviceManagementConfigurationSettingApplicabilityTechnologiesappleRemoteManagement       DeviceManagementConfigurationSettingApplicabilityTechnologies = "AppleRemoteManagement"
	DeviceManagementConfigurationSettingApplicabilityTechnologiesconfigManager               DeviceManagementConfigurationSettingApplicabilityTechnologies = "ConfigManager"
	DeviceManagementConfigurationSettingApplicabilityTechnologiesendpointPrivilegeManagement DeviceManagementConfigurationSettingApplicabilityTechnologies = "EndpointPrivilegeManagement"
	DeviceManagementConfigurationSettingApplicabilityTechnologiesenrollment                  DeviceManagementConfigurationSettingApplicabilityTechnologies = "Enrollment"
	DeviceManagementConfigurationSettingApplicabilityTechnologiesexchangeOnline              DeviceManagementConfigurationSettingApplicabilityTechnologies = "ExchangeOnline"
	DeviceManagementConfigurationSettingApplicabilityTechnologieslinuxMdm                    DeviceManagementConfigurationSettingApplicabilityTechnologies = "LinuxMdm"
	DeviceManagementConfigurationSettingApplicabilityTechnologiesmdm                         DeviceManagementConfigurationSettingApplicabilityTechnologies = "Mdm"
	DeviceManagementConfigurationSettingApplicabilityTechnologiesmicrosoftSense              DeviceManagementConfigurationSettingApplicabilityTechnologies = "MicrosoftSense"
	DeviceManagementConfigurationSettingApplicabilityTechnologiesnone                        DeviceManagementConfigurationSettingApplicabilityTechnologies = "None"
	DeviceManagementConfigurationSettingApplicabilityTechnologieswindows10XManagement        DeviceManagementConfigurationSettingApplicabilityTechnologies = "Windows10XManagement"
)

func PossibleValuesForDeviceManagementConfigurationSettingApplicabilityTechnologies() []string {
	return []string{
		string(DeviceManagementConfigurationSettingApplicabilityTechnologiesappleRemoteManagement),
		string(DeviceManagementConfigurationSettingApplicabilityTechnologiesconfigManager),
		string(DeviceManagementConfigurationSettingApplicabilityTechnologiesendpointPrivilegeManagement),
		string(DeviceManagementConfigurationSettingApplicabilityTechnologiesenrollment),
		string(DeviceManagementConfigurationSettingApplicabilityTechnologiesexchangeOnline),
		string(DeviceManagementConfigurationSettingApplicabilityTechnologieslinuxMdm),
		string(DeviceManagementConfigurationSettingApplicabilityTechnologiesmdm),
		string(DeviceManagementConfigurationSettingApplicabilityTechnologiesmicrosoftSense),
		string(DeviceManagementConfigurationSettingApplicabilityTechnologiesnone),
		string(DeviceManagementConfigurationSettingApplicabilityTechnologieswindows10XManagement),
	}
}

func parseDeviceManagementConfigurationSettingApplicabilityTechnologies(input string) (*DeviceManagementConfigurationSettingApplicabilityTechnologies, error) {
	vals := map[string]DeviceManagementConfigurationSettingApplicabilityTechnologies{
		"appleremotemanagement":       DeviceManagementConfigurationSettingApplicabilityTechnologiesappleRemoteManagement,
		"configmanager":               DeviceManagementConfigurationSettingApplicabilityTechnologiesconfigManager,
		"endpointprivilegemanagement": DeviceManagementConfigurationSettingApplicabilityTechnologiesendpointPrivilegeManagement,
		"enrollment":                  DeviceManagementConfigurationSettingApplicabilityTechnologiesenrollment,
		"exchangeonline":              DeviceManagementConfigurationSettingApplicabilityTechnologiesexchangeOnline,
		"linuxmdm":                    DeviceManagementConfigurationSettingApplicabilityTechnologieslinuxMdm,
		"mdm":                         DeviceManagementConfigurationSettingApplicabilityTechnologiesmdm,
		"microsoftsense":              DeviceManagementConfigurationSettingApplicabilityTechnologiesmicrosoftSense,
		"none":                        DeviceManagementConfigurationSettingApplicabilityTechnologiesnone,
		"windows10xmanagement":        DeviceManagementConfigurationSettingApplicabilityTechnologieswindows10XManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingApplicabilityTechnologies(input)
	return &out, nil
}

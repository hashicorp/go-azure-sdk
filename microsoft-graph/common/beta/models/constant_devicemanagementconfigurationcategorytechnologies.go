package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationCategoryTechnologies string

const (
	DeviceManagementConfigurationCategoryTechnologiesappleRemoteManagement       DeviceManagementConfigurationCategoryTechnologies = "AppleRemoteManagement"
	DeviceManagementConfigurationCategoryTechnologiesconfigManager               DeviceManagementConfigurationCategoryTechnologies = "ConfigManager"
	DeviceManagementConfigurationCategoryTechnologiesendpointPrivilegeManagement DeviceManagementConfigurationCategoryTechnologies = "EndpointPrivilegeManagement"
	DeviceManagementConfigurationCategoryTechnologiesenrollment                  DeviceManagementConfigurationCategoryTechnologies = "Enrollment"
	DeviceManagementConfigurationCategoryTechnologiesexchangeOnline              DeviceManagementConfigurationCategoryTechnologies = "ExchangeOnline"
	DeviceManagementConfigurationCategoryTechnologieslinuxMdm                    DeviceManagementConfigurationCategoryTechnologies = "LinuxMdm"
	DeviceManagementConfigurationCategoryTechnologiesmdm                         DeviceManagementConfigurationCategoryTechnologies = "Mdm"
	DeviceManagementConfigurationCategoryTechnologiesmicrosoftSense              DeviceManagementConfigurationCategoryTechnologies = "MicrosoftSense"
	DeviceManagementConfigurationCategoryTechnologiesnone                        DeviceManagementConfigurationCategoryTechnologies = "None"
	DeviceManagementConfigurationCategoryTechnologieswindows10XManagement        DeviceManagementConfigurationCategoryTechnologies = "Windows10XManagement"
)

func PossibleValuesForDeviceManagementConfigurationCategoryTechnologies() []string {
	return []string{
		string(DeviceManagementConfigurationCategoryTechnologiesappleRemoteManagement),
		string(DeviceManagementConfigurationCategoryTechnologiesconfigManager),
		string(DeviceManagementConfigurationCategoryTechnologiesendpointPrivilegeManagement),
		string(DeviceManagementConfigurationCategoryTechnologiesenrollment),
		string(DeviceManagementConfigurationCategoryTechnologiesexchangeOnline),
		string(DeviceManagementConfigurationCategoryTechnologieslinuxMdm),
		string(DeviceManagementConfigurationCategoryTechnologiesmdm),
		string(DeviceManagementConfigurationCategoryTechnologiesmicrosoftSense),
		string(DeviceManagementConfigurationCategoryTechnologiesnone),
		string(DeviceManagementConfigurationCategoryTechnologieswindows10XManagement),
	}
}

func parseDeviceManagementConfigurationCategoryTechnologies(input string) (*DeviceManagementConfigurationCategoryTechnologies, error) {
	vals := map[string]DeviceManagementConfigurationCategoryTechnologies{
		"appleremotemanagement":       DeviceManagementConfigurationCategoryTechnologiesappleRemoteManagement,
		"configmanager":               DeviceManagementConfigurationCategoryTechnologiesconfigManager,
		"endpointprivilegemanagement": DeviceManagementConfigurationCategoryTechnologiesendpointPrivilegeManagement,
		"enrollment":                  DeviceManagementConfigurationCategoryTechnologiesenrollment,
		"exchangeonline":              DeviceManagementConfigurationCategoryTechnologiesexchangeOnline,
		"linuxmdm":                    DeviceManagementConfigurationCategoryTechnologieslinuxMdm,
		"mdm":                         DeviceManagementConfigurationCategoryTechnologiesmdm,
		"microsoftsense":              DeviceManagementConfigurationCategoryTechnologiesmicrosoftSense,
		"none":                        DeviceManagementConfigurationCategoryTechnologiesnone,
		"windows10xmanagement":        DeviceManagementConfigurationCategoryTechnologieswindows10XManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationCategoryTechnologies(input)
	return &out, nil
}

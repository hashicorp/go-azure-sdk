package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyTechnologies string

const (
	DeviceManagementConfigurationPolicyTechnologiesappleRemoteManagement       DeviceManagementConfigurationPolicyTechnologies = "AppleRemoteManagement"
	DeviceManagementConfigurationPolicyTechnologiesconfigManager               DeviceManagementConfigurationPolicyTechnologies = "ConfigManager"
	DeviceManagementConfigurationPolicyTechnologiesendpointPrivilegeManagement DeviceManagementConfigurationPolicyTechnologies = "EndpointPrivilegeManagement"
	DeviceManagementConfigurationPolicyTechnologiesenrollment                  DeviceManagementConfigurationPolicyTechnologies = "Enrollment"
	DeviceManagementConfigurationPolicyTechnologiesexchangeOnline              DeviceManagementConfigurationPolicyTechnologies = "ExchangeOnline"
	DeviceManagementConfigurationPolicyTechnologieslinuxMdm                    DeviceManagementConfigurationPolicyTechnologies = "LinuxMdm"
	DeviceManagementConfigurationPolicyTechnologiesmdm                         DeviceManagementConfigurationPolicyTechnologies = "Mdm"
	DeviceManagementConfigurationPolicyTechnologiesmicrosoftSense              DeviceManagementConfigurationPolicyTechnologies = "MicrosoftSense"
	DeviceManagementConfigurationPolicyTechnologiesnone                        DeviceManagementConfigurationPolicyTechnologies = "None"
	DeviceManagementConfigurationPolicyTechnologieswindows10XManagement        DeviceManagementConfigurationPolicyTechnologies = "Windows10XManagement"
)

func PossibleValuesForDeviceManagementConfigurationPolicyTechnologies() []string {
	return []string{
		string(DeviceManagementConfigurationPolicyTechnologiesappleRemoteManagement),
		string(DeviceManagementConfigurationPolicyTechnologiesconfigManager),
		string(DeviceManagementConfigurationPolicyTechnologiesendpointPrivilegeManagement),
		string(DeviceManagementConfigurationPolicyTechnologiesenrollment),
		string(DeviceManagementConfigurationPolicyTechnologiesexchangeOnline),
		string(DeviceManagementConfigurationPolicyTechnologieslinuxMdm),
		string(DeviceManagementConfigurationPolicyTechnologiesmdm),
		string(DeviceManagementConfigurationPolicyTechnologiesmicrosoftSense),
		string(DeviceManagementConfigurationPolicyTechnologiesnone),
		string(DeviceManagementConfigurationPolicyTechnologieswindows10XManagement),
	}
}

func parseDeviceManagementConfigurationPolicyTechnologies(input string) (*DeviceManagementConfigurationPolicyTechnologies, error) {
	vals := map[string]DeviceManagementConfigurationPolicyTechnologies{
		"appleremotemanagement":       DeviceManagementConfigurationPolicyTechnologiesappleRemoteManagement,
		"configmanager":               DeviceManagementConfigurationPolicyTechnologiesconfigManager,
		"endpointprivilegemanagement": DeviceManagementConfigurationPolicyTechnologiesendpointPrivilegeManagement,
		"enrollment":                  DeviceManagementConfigurationPolicyTechnologiesenrollment,
		"exchangeonline":              DeviceManagementConfigurationPolicyTechnologiesexchangeOnline,
		"linuxmdm":                    DeviceManagementConfigurationPolicyTechnologieslinuxMdm,
		"mdm":                         DeviceManagementConfigurationPolicyTechnologiesmdm,
		"microsoftsense":              DeviceManagementConfigurationPolicyTechnologiesmicrosoftSense,
		"none":                        DeviceManagementConfigurationPolicyTechnologiesnone,
		"windows10xmanagement":        DeviceManagementConfigurationPolicyTechnologieswindows10XManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationPolicyTechnologies(input)
	return &out, nil
}

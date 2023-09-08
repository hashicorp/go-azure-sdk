package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyTemplateTechnologies string

const (
	DeviceManagementConfigurationPolicyTemplateTechnologiesappleRemoteManagement       DeviceManagementConfigurationPolicyTemplateTechnologies = "AppleRemoteManagement"
	DeviceManagementConfigurationPolicyTemplateTechnologiesconfigManager               DeviceManagementConfigurationPolicyTemplateTechnologies = "ConfigManager"
	DeviceManagementConfigurationPolicyTemplateTechnologiesendpointPrivilegeManagement DeviceManagementConfigurationPolicyTemplateTechnologies = "EndpointPrivilegeManagement"
	DeviceManagementConfigurationPolicyTemplateTechnologiesenrollment                  DeviceManagementConfigurationPolicyTemplateTechnologies = "Enrollment"
	DeviceManagementConfigurationPolicyTemplateTechnologiesexchangeOnline              DeviceManagementConfigurationPolicyTemplateTechnologies = "ExchangeOnline"
	DeviceManagementConfigurationPolicyTemplateTechnologieslinuxMdm                    DeviceManagementConfigurationPolicyTemplateTechnologies = "LinuxMdm"
	DeviceManagementConfigurationPolicyTemplateTechnologiesmdm                         DeviceManagementConfigurationPolicyTemplateTechnologies = "Mdm"
	DeviceManagementConfigurationPolicyTemplateTechnologiesmicrosoftSense              DeviceManagementConfigurationPolicyTemplateTechnologies = "MicrosoftSense"
	DeviceManagementConfigurationPolicyTemplateTechnologiesnone                        DeviceManagementConfigurationPolicyTemplateTechnologies = "None"
	DeviceManagementConfigurationPolicyTemplateTechnologieswindows10XManagement        DeviceManagementConfigurationPolicyTemplateTechnologies = "Windows10XManagement"
)

func PossibleValuesForDeviceManagementConfigurationPolicyTemplateTechnologies() []string {
	return []string{
		string(DeviceManagementConfigurationPolicyTemplateTechnologiesappleRemoteManagement),
		string(DeviceManagementConfigurationPolicyTemplateTechnologiesconfigManager),
		string(DeviceManagementConfigurationPolicyTemplateTechnologiesendpointPrivilegeManagement),
		string(DeviceManagementConfigurationPolicyTemplateTechnologiesenrollment),
		string(DeviceManagementConfigurationPolicyTemplateTechnologiesexchangeOnline),
		string(DeviceManagementConfigurationPolicyTemplateTechnologieslinuxMdm),
		string(DeviceManagementConfigurationPolicyTemplateTechnologiesmdm),
		string(DeviceManagementConfigurationPolicyTemplateTechnologiesmicrosoftSense),
		string(DeviceManagementConfigurationPolicyTemplateTechnologiesnone),
		string(DeviceManagementConfigurationPolicyTemplateTechnologieswindows10XManagement),
	}
}

func parseDeviceManagementConfigurationPolicyTemplateTechnologies(input string) (*DeviceManagementConfigurationPolicyTemplateTechnologies, error) {
	vals := map[string]DeviceManagementConfigurationPolicyTemplateTechnologies{
		"appleremotemanagement":       DeviceManagementConfigurationPolicyTemplateTechnologiesappleRemoteManagement,
		"configmanager":               DeviceManagementConfigurationPolicyTemplateTechnologiesconfigManager,
		"endpointprivilegemanagement": DeviceManagementConfigurationPolicyTemplateTechnologiesendpointPrivilegeManagement,
		"enrollment":                  DeviceManagementConfigurationPolicyTemplateTechnologiesenrollment,
		"exchangeonline":              DeviceManagementConfigurationPolicyTemplateTechnologiesexchangeOnline,
		"linuxmdm":                    DeviceManagementConfigurationPolicyTemplateTechnologieslinuxMdm,
		"mdm":                         DeviceManagementConfigurationPolicyTemplateTechnologiesmdm,
		"microsoftsense":              DeviceManagementConfigurationPolicyTemplateTechnologiesmicrosoftSense,
		"none":                        DeviceManagementConfigurationPolicyTemplateTechnologiesnone,
		"windows10xmanagement":        DeviceManagementConfigurationPolicyTemplateTechnologieswindows10XManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationPolicyTemplateTechnologies(input)
	return &out, nil
}

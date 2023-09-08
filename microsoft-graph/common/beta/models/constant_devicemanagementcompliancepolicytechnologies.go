package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementCompliancePolicyTechnologies string

const (
	DeviceManagementCompliancePolicyTechnologiesappleRemoteManagement       DeviceManagementCompliancePolicyTechnologies = "AppleRemoteManagement"
	DeviceManagementCompliancePolicyTechnologiesconfigManager               DeviceManagementCompliancePolicyTechnologies = "ConfigManager"
	DeviceManagementCompliancePolicyTechnologiesendpointPrivilegeManagement DeviceManagementCompliancePolicyTechnologies = "EndpointPrivilegeManagement"
	DeviceManagementCompliancePolicyTechnologiesenrollment                  DeviceManagementCompliancePolicyTechnologies = "Enrollment"
	DeviceManagementCompliancePolicyTechnologiesexchangeOnline              DeviceManagementCompliancePolicyTechnologies = "ExchangeOnline"
	DeviceManagementCompliancePolicyTechnologieslinuxMdm                    DeviceManagementCompliancePolicyTechnologies = "LinuxMdm"
	DeviceManagementCompliancePolicyTechnologiesmdm                         DeviceManagementCompliancePolicyTechnologies = "Mdm"
	DeviceManagementCompliancePolicyTechnologiesmicrosoftSense              DeviceManagementCompliancePolicyTechnologies = "MicrosoftSense"
	DeviceManagementCompliancePolicyTechnologiesnone                        DeviceManagementCompliancePolicyTechnologies = "None"
	DeviceManagementCompliancePolicyTechnologieswindows10XManagement        DeviceManagementCompliancePolicyTechnologies = "Windows10XManagement"
)

func PossibleValuesForDeviceManagementCompliancePolicyTechnologies() []string {
	return []string{
		string(DeviceManagementCompliancePolicyTechnologiesappleRemoteManagement),
		string(DeviceManagementCompliancePolicyTechnologiesconfigManager),
		string(DeviceManagementCompliancePolicyTechnologiesendpointPrivilegeManagement),
		string(DeviceManagementCompliancePolicyTechnologiesenrollment),
		string(DeviceManagementCompliancePolicyTechnologiesexchangeOnline),
		string(DeviceManagementCompliancePolicyTechnologieslinuxMdm),
		string(DeviceManagementCompliancePolicyTechnologiesmdm),
		string(DeviceManagementCompliancePolicyTechnologiesmicrosoftSense),
		string(DeviceManagementCompliancePolicyTechnologiesnone),
		string(DeviceManagementCompliancePolicyTechnologieswindows10XManagement),
	}
}

func parseDeviceManagementCompliancePolicyTechnologies(input string) (*DeviceManagementCompliancePolicyTechnologies, error) {
	vals := map[string]DeviceManagementCompliancePolicyTechnologies{
		"appleremotemanagement":       DeviceManagementCompliancePolicyTechnologiesappleRemoteManagement,
		"configmanager":               DeviceManagementCompliancePolicyTechnologiesconfigManager,
		"endpointprivilegemanagement": DeviceManagementCompliancePolicyTechnologiesendpointPrivilegeManagement,
		"enrollment":                  DeviceManagementCompliancePolicyTechnologiesenrollment,
		"exchangeonline":              DeviceManagementCompliancePolicyTechnologiesexchangeOnline,
		"linuxmdm":                    DeviceManagementCompliancePolicyTechnologieslinuxMdm,
		"mdm":                         DeviceManagementCompliancePolicyTechnologiesmdm,
		"microsoftsense":              DeviceManagementCompliancePolicyTechnologiesmicrosoftSense,
		"none":                        DeviceManagementCompliancePolicyTechnologiesnone,
		"windows10xmanagement":        DeviceManagementCompliancePolicyTechnologieswindows10XManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementCompliancePolicyTechnologies(input)
	return &out, nil
}

package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType string

const (
	DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustTypeaddWorkAccount DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType = "AddWorkAccount"
	DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustTypeazureAdJoined  DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType = "AzureAdJoined"
	DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustTypemdmOnly        DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType = "MdmOnly"
	DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustTypenone           DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType = "None"
)

func PossibleValuesForDeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType() []string {
	return []string{
		string(DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustTypeaddWorkAccount),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustTypeazureAdJoined),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustTypemdmOnly),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustTypenone),
	}
}

func parseDeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType(input string) (*DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType, error) {
	vals := map[string]DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType{
		"addworkaccount": DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustTypeaddWorkAccount,
		"azureadjoined":  DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustTypeazureAdJoined,
		"mdmonly":        DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustTypemdmOnly,
		"none":           DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType(input)
	return &out, nil
}

package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingApplicabilityTechnologies string

const (
	DeviceManagementConfigurationSettingApplicabilityTechnologies_AppleRemoteManagement       DeviceManagementConfigurationSettingApplicabilityTechnologies = "appleRemoteManagement"
	DeviceManagementConfigurationSettingApplicabilityTechnologies_ConfigManager               DeviceManagementConfigurationSettingApplicabilityTechnologies = "configManager"
	DeviceManagementConfigurationSettingApplicabilityTechnologies_EndpointPrivilegeManagement DeviceManagementConfigurationSettingApplicabilityTechnologies = "endpointPrivilegeManagement"
	DeviceManagementConfigurationSettingApplicabilityTechnologies_Enrollment                  DeviceManagementConfigurationSettingApplicabilityTechnologies = "enrollment"
	DeviceManagementConfigurationSettingApplicabilityTechnologies_ExchangeOnline              DeviceManagementConfigurationSettingApplicabilityTechnologies = "exchangeOnline"
	DeviceManagementConfigurationSettingApplicabilityTechnologies_LinuxMdm                    DeviceManagementConfigurationSettingApplicabilityTechnologies = "linuxMdm"
	DeviceManagementConfigurationSettingApplicabilityTechnologies_Mdm                         DeviceManagementConfigurationSettingApplicabilityTechnologies = "mdm"
	DeviceManagementConfigurationSettingApplicabilityTechnologies_MicrosoftSense              DeviceManagementConfigurationSettingApplicabilityTechnologies = "microsoftSense"
	DeviceManagementConfigurationSettingApplicabilityTechnologies_MobileApplicationManagement DeviceManagementConfigurationSettingApplicabilityTechnologies = "mobileApplicationManagement"
	DeviceManagementConfigurationSettingApplicabilityTechnologies_None                        DeviceManagementConfigurationSettingApplicabilityTechnologies = "none"
	DeviceManagementConfigurationSettingApplicabilityTechnologies_Windows10XManagement        DeviceManagementConfigurationSettingApplicabilityTechnologies = "windows10XManagement"
)

func PossibleValuesForDeviceManagementConfigurationSettingApplicabilityTechnologies() []string {
	return []string{
		string(DeviceManagementConfigurationSettingApplicabilityTechnologies_AppleRemoteManagement),
		string(DeviceManagementConfigurationSettingApplicabilityTechnologies_ConfigManager),
		string(DeviceManagementConfigurationSettingApplicabilityTechnologies_EndpointPrivilegeManagement),
		string(DeviceManagementConfigurationSettingApplicabilityTechnologies_Enrollment),
		string(DeviceManagementConfigurationSettingApplicabilityTechnologies_ExchangeOnline),
		string(DeviceManagementConfigurationSettingApplicabilityTechnologies_LinuxMdm),
		string(DeviceManagementConfigurationSettingApplicabilityTechnologies_Mdm),
		string(DeviceManagementConfigurationSettingApplicabilityTechnologies_MicrosoftSense),
		string(DeviceManagementConfigurationSettingApplicabilityTechnologies_MobileApplicationManagement),
		string(DeviceManagementConfigurationSettingApplicabilityTechnologies_None),
		string(DeviceManagementConfigurationSettingApplicabilityTechnologies_Windows10XManagement),
	}
}

func (s *DeviceManagementConfigurationSettingApplicabilityTechnologies) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSettingApplicabilityTechnologies(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSettingApplicabilityTechnologies(input string) (*DeviceManagementConfigurationSettingApplicabilityTechnologies, error) {
	vals := map[string]DeviceManagementConfigurationSettingApplicabilityTechnologies{
		"appleremotemanagement":       DeviceManagementConfigurationSettingApplicabilityTechnologies_AppleRemoteManagement,
		"configmanager":               DeviceManagementConfigurationSettingApplicabilityTechnologies_ConfigManager,
		"endpointprivilegemanagement": DeviceManagementConfigurationSettingApplicabilityTechnologies_EndpointPrivilegeManagement,
		"enrollment":                  DeviceManagementConfigurationSettingApplicabilityTechnologies_Enrollment,
		"exchangeonline":              DeviceManagementConfigurationSettingApplicabilityTechnologies_ExchangeOnline,
		"linuxmdm":                    DeviceManagementConfigurationSettingApplicabilityTechnologies_LinuxMdm,
		"mdm":                         DeviceManagementConfigurationSettingApplicabilityTechnologies_Mdm,
		"microsoftsense":              DeviceManagementConfigurationSettingApplicabilityTechnologies_MicrosoftSense,
		"mobileapplicationmanagement": DeviceManagementConfigurationSettingApplicabilityTechnologies_MobileApplicationManagement,
		"none":                        DeviceManagementConfigurationSettingApplicabilityTechnologies_None,
		"windows10xmanagement":        DeviceManagementConfigurationSettingApplicabilityTechnologies_Windows10XManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingApplicabilityTechnologies(input)
	return &out, nil
}

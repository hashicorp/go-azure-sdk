package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies string

const (
	DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_AppleRemoteManagement       DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies = "appleRemoteManagement"
	DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_ConfigManager               DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies = "configManager"
	DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_EndpointPrivilegeManagement DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies = "endpointPrivilegeManagement"
	DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_Enrollment                  DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies = "enrollment"
	DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_ExchangeOnline              DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies = "exchangeOnline"
	DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_LinuxMdm                    DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies = "linuxMdm"
	DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_Mdm                         DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies = "mdm"
	DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_MicrosoftSense              DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies = "microsoftSense"
	DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_MobileApplicationManagement DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies = "mobileApplicationManagement"
	DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_None                        DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies = "none"
	DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_Windows10XManagement        DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies = "windows10XManagement"
)

func PossibleValuesForDeviceManagementConfigurationWindowsSettingApplicabilityTechnologies() []string {
	return []string{
		string(DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_AppleRemoteManagement),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_ConfigManager),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_EndpointPrivilegeManagement),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_Enrollment),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_ExchangeOnline),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_LinuxMdm),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_Mdm),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_MicrosoftSense),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_MobileApplicationManagement),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_None),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_Windows10XManagement),
	}
}

func (s *DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationWindowsSettingApplicabilityTechnologies(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationWindowsSettingApplicabilityTechnologies(input string) (*DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies, error) {
	vals := map[string]DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies{
		"appleremotemanagement":       DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_AppleRemoteManagement,
		"configmanager":               DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_ConfigManager,
		"endpointprivilegemanagement": DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_EndpointPrivilegeManagement,
		"enrollment":                  DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_Enrollment,
		"exchangeonline":              DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_ExchangeOnline,
		"linuxmdm":                    DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_LinuxMdm,
		"mdm":                         DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_Mdm,
		"microsoftsense":              DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_MicrosoftSense,
		"mobileapplicationmanagement": DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_MobileApplicationManagement,
		"none":                        DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_None,
		"windows10xmanagement":        DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies_Windows10XManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies(input)
	return &out, nil
}

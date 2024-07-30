package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies string

const (
	DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_AppleRemoteManagement       DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies = "appleRemoteManagement"
	DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_ConfigManager               DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies = "configManager"
	DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_EndpointPrivilegeManagement DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies = "endpointPrivilegeManagement"
	DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_Enrollment                  DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies = "enrollment"
	DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_ExchangeOnline              DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies = "exchangeOnline"
	DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_LinuxMdm                    DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies = "linuxMdm"
	DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_Mdm                         DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies = "mdm"
	DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_MicrosoftSense              DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies = "microsoftSense"
	DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_MobileApplicationManagement DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies = "mobileApplicationManagement"
	DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_None                        DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies = "none"
	DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_Windows10XManagement        DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies = "windows10XManagement"
)

func PossibleValuesForDeviceManagementConfigurationApplicationSettingApplicabilityTechnologies() []string {
	return []string{
		string(DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_AppleRemoteManagement),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_ConfigManager),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_EndpointPrivilegeManagement),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_Enrollment),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_ExchangeOnline),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_LinuxMdm),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_Mdm),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_MicrosoftSense),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_MobileApplicationManagement),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_None),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_Windows10XManagement),
	}
}

func (s *DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationApplicationSettingApplicabilityTechnologies(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationApplicationSettingApplicabilityTechnologies(input string) (*DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies, error) {
	vals := map[string]DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies{
		"appleremotemanagement":       DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_AppleRemoteManagement,
		"configmanager":               DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_ConfigManager,
		"endpointprivilegemanagement": DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_EndpointPrivilegeManagement,
		"enrollment":                  DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_Enrollment,
		"exchangeonline":              DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_ExchangeOnline,
		"linuxmdm":                    DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_LinuxMdm,
		"mdm":                         DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_Mdm,
		"microsoftsense":              DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_MicrosoftSense,
		"mobileapplicationmanagement": DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_MobileApplicationManagement,
		"none":                        DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_None,
		"windows10xmanagement":        DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies_Windows10XManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies(input)
	return &out, nil
}

package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies string

const (
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_AppleRemoteManagement       DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies = "appleRemoteManagement"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_ConfigManager               DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies = "configManager"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_EndpointPrivilegeManagement DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies = "endpointPrivilegeManagement"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_Enrollment                  DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies = "enrollment"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_ExchangeOnline              DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies = "exchangeOnline"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_LinuxMdm                    DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies = "linuxMdm"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_Mdm                         DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies = "mdm"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_MicrosoftSense              DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies = "microsoftSense"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_MobileApplicationManagement DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies = "mobileApplicationManagement"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_None                        DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies = "none"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_Windows10XManagement        DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies = "windows10XManagement"
)

func PossibleValuesForDeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies() []string {
	return []string{
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_AppleRemoteManagement),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_ConfigManager),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_EndpointPrivilegeManagement),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_Enrollment),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_ExchangeOnline),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_LinuxMdm),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_Mdm),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_MicrosoftSense),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_MobileApplicationManagement),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_None),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_Windows10XManagement),
	}
}

func (s *DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies(input string) (*DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies, error) {
	vals := map[string]DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies{
		"appleremotemanagement":       DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_AppleRemoteManagement,
		"configmanager":               DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_ConfigManager,
		"endpointprivilegemanagement": DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_EndpointPrivilegeManagement,
		"enrollment":                  DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_Enrollment,
		"exchangeonline":              DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_ExchangeOnline,
		"linuxmdm":                    DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_LinuxMdm,
		"mdm":                         DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_Mdm,
		"microsoftsense":              DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_MicrosoftSense,
		"mobileapplicationmanagement": DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_MobileApplicationManagement,
		"none":                        DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_None,
		"windows10xmanagement":        DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies_Windows10XManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies(input)
	return &out, nil
}

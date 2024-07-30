package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyTechnologies string

const (
	DeviceManagementConfigurationPolicyTechnologies_AppleRemoteManagement       DeviceManagementConfigurationPolicyTechnologies = "appleRemoteManagement"
	DeviceManagementConfigurationPolicyTechnologies_ConfigManager               DeviceManagementConfigurationPolicyTechnologies = "configManager"
	DeviceManagementConfigurationPolicyTechnologies_EndpointPrivilegeManagement DeviceManagementConfigurationPolicyTechnologies = "endpointPrivilegeManagement"
	DeviceManagementConfigurationPolicyTechnologies_Enrollment                  DeviceManagementConfigurationPolicyTechnologies = "enrollment"
	DeviceManagementConfigurationPolicyTechnologies_ExchangeOnline              DeviceManagementConfigurationPolicyTechnologies = "exchangeOnline"
	DeviceManagementConfigurationPolicyTechnologies_LinuxMdm                    DeviceManagementConfigurationPolicyTechnologies = "linuxMdm"
	DeviceManagementConfigurationPolicyTechnologies_Mdm                         DeviceManagementConfigurationPolicyTechnologies = "mdm"
	DeviceManagementConfigurationPolicyTechnologies_MicrosoftSense              DeviceManagementConfigurationPolicyTechnologies = "microsoftSense"
	DeviceManagementConfigurationPolicyTechnologies_MobileApplicationManagement DeviceManagementConfigurationPolicyTechnologies = "mobileApplicationManagement"
	DeviceManagementConfigurationPolicyTechnologies_None                        DeviceManagementConfigurationPolicyTechnologies = "none"
	DeviceManagementConfigurationPolicyTechnologies_Windows10XManagement        DeviceManagementConfigurationPolicyTechnologies = "windows10XManagement"
)

func PossibleValuesForDeviceManagementConfigurationPolicyTechnologies() []string {
	return []string{
		string(DeviceManagementConfigurationPolicyTechnologies_AppleRemoteManagement),
		string(DeviceManagementConfigurationPolicyTechnologies_ConfigManager),
		string(DeviceManagementConfigurationPolicyTechnologies_EndpointPrivilegeManagement),
		string(DeviceManagementConfigurationPolicyTechnologies_Enrollment),
		string(DeviceManagementConfigurationPolicyTechnologies_ExchangeOnline),
		string(DeviceManagementConfigurationPolicyTechnologies_LinuxMdm),
		string(DeviceManagementConfigurationPolicyTechnologies_Mdm),
		string(DeviceManagementConfigurationPolicyTechnologies_MicrosoftSense),
		string(DeviceManagementConfigurationPolicyTechnologies_MobileApplicationManagement),
		string(DeviceManagementConfigurationPolicyTechnologies_None),
		string(DeviceManagementConfigurationPolicyTechnologies_Windows10XManagement),
	}
}

func (s *DeviceManagementConfigurationPolicyTechnologies) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationPolicyTechnologies(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationPolicyTechnologies(input string) (*DeviceManagementConfigurationPolicyTechnologies, error) {
	vals := map[string]DeviceManagementConfigurationPolicyTechnologies{
		"appleremotemanagement":       DeviceManagementConfigurationPolicyTechnologies_AppleRemoteManagement,
		"configmanager":               DeviceManagementConfigurationPolicyTechnologies_ConfigManager,
		"endpointprivilegemanagement": DeviceManagementConfigurationPolicyTechnologies_EndpointPrivilegeManagement,
		"enrollment":                  DeviceManagementConfigurationPolicyTechnologies_Enrollment,
		"exchangeonline":              DeviceManagementConfigurationPolicyTechnologies_ExchangeOnline,
		"linuxmdm":                    DeviceManagementConfigurationPolicyTechnologies_LinuxMdm,
		"mdm":                         DeviceManagementConfigurationPolicyTechnologies_Mdm,
		"microsoftsense":              DeviceManagementConfigurationPolicyTechnologies_MicrosoftSense,
		"mobileapplicationmanagement": DeviceManagementConfigurationPolicyTechnologies_MobileApplicationManagement,
		"none":                        DeviceManagementConfigurationPolicyTechnologies_None,
		"windows10xmanagement":        DeviceManagementConfigurationPolicyTechnologies_Windows10XManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationPolicyTechnologies(input)
	return &out, nil
}

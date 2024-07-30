package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyTemplateTechnologies string

const (
	DeviceManagementConfigurationPolicyTemplateTechnologies_AppleRemoteManagement       DeviceManagementConfigurationPolicyTemplateTechnologies = "appleRemoteManagement"
	DeviceManagementConfigurationPolicyTemplateTechnologies_ConfigManager               DeviceManagementConfigurationPolicyTemplateTechnologies = "configManager"
	DeviceManagementConfigurationPolicyTemplateTechnologies_EndpointPrivilegeManagement DeviceManagementConfigurationPolicyTemplateTechnologies = "endpointPrivilegeManagement"
	DeviceManagementConfigurationPolicyTemplateTechnologies_Enrollment                  DeviceManagementConfigurationPolicyTemplateTechnologies = "enrollment"
	DeviceManagementConfigurationPolicyTemplateTechnologies_ExchangeOnline              DeviceManagementConfigurationPolicyTemplateTechnologies = "exchangeOnline"
	DeviceManagementConfigurationPolicyTemplateTechnologies_LinuxMdm                    DeviceManagementConfigurationPolicyTemplateTechnologies = "linuxMdm"
	DeviceManagementConfigurationPolicyTemplateTechnologies_Mdm                         DeviceManagementConfigurationPolicyTemplateTechnologies = "mdm"
	DeviceManagementConfigurationPolicyTemplateTechnologies_MicrosoftSense              DeviceManagementConfigurationPolicyTemplateTechnologies = "microsoftSense"
	DeviceManagementConfigurationPolicyTemplateTechnologies_MobileApplicationManagement DeviceManagementConfigurationPolicyTemplateTechnologies = "mobileApplicationManagement"
	DeviceManagementConfigurationPolicyTemplateTechnologies_None                        DeviceManagementConfigurationPolicyTemplateTechnologies = "none"
	DeviceManagementConfigurationPolicyTemplateTechnologies_Windows10XManagement        DeviceManagementConfigurationPolicyTemplateTechnologies = "windows10XManagement"
)

func PossibleValuesForDeviceManagementConfigurationPolicyTemplateTechnologies() []string {
	return []string{
		string(DeviceManagementConfigurationPolicyTemplateTechnologies_AppleRemoteManagement),
		string(DeviceManagementConfigurationPolicyTemplateTechnologies_ConfigManager),
		string(DeviceManagementConfigurationPolicyTemplateTechnologies_EndpointPrivilegeManagement),
		string(DeviceManagementConfigurationPolicyTemplateTechnologies_Enrollment),
		string(DeviceManagementConfigurationPolicyTemplateTechnologies_ExchangeOnline),
		string(DeviceManagementConfigurationPolicyTemplateTechnologies_LinuxMdm),
		string(DeviceManagementConfigurationPolicyTemplateTechnologies_Mdm),
		string(DeviceManagementConfigurationPolicyTemplateTechnologies_MicrosoftSense),
		string(DeviceManagementConfigurationPolicyTemplateTechnologies_MobileApplicationManagement),
		string(DeviceManagementConfigurationPolicyTemplateTechnologies_None),
		string(DeviceManagementConfigurationPolicyTemplateTechnologies_Windows10XManagement),
	}
}

func (s *DeviceManagementConfigurationPolicyTemplateTechnologies) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationPolicyTemplateTechnologies(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationPolicyTemplateTechnologies(input string) (*DeviceManagementConfigurationPolicyTemplateTechnologies, error) {
	vals := map[string]DeviceManagementConfigurationPolicyTemplateTechnologies{
		"appleremotemanagement":       DeviceManagementConfigurationPolicyTemplateTechnologies_AppleRemoteManagement,
		"configmanager":               DeviceManagementConfigurationPolicyTemplateTechnologies_ConfigManager,
		"endpointprivilegemanagement": DeviceManagementConfigurationPolicyTemplateTechnologies_EndpointPrivilegeManagement,
		"enrollment":                  DeviceManagementConfigurationPolicyTemplateTechnologies_Enrollment,
		"exchangeonline":              DeviceManagementConfigurationPolicyTemplateTechnologies_ExchangeOnline,
		"linuxmdm":                    DeviceManagementConfigurationPolicyTemplateTechnologies_LinuxMdm,
		"mdm":                         DeviceManagementConfigurationPolicyTemplateTechnologies_Mdm,
		"microsoftsense":              DeviceManagementConfigurationPolicyTemplateTechnologies_MicrosoftSense,
		"mobileapplicationmanagement": DeviceManagementConfigurationPolicyTemplateTechnologies_MobileApplicationManagement,
		"none":                        DeviceManagementConfigurationPolicyTemplateTechnologies_None,
		"windows10xmanagement":        DeviceManagementConfigurationPolicyTemplateTechnologies_Windows10XManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationPolicyTemplateTechnologies(input)
	return &out, nil
}

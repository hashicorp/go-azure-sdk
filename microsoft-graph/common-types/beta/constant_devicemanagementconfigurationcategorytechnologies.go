package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationCategoryTechnologies string

const (
	DeviceManagementConfigurationCategoryTechnologies_AppleRemoteManagement       DeviceManagementConfigurationCategoryTechnologies = "appleRemoteManagement"
	DeviceManagementConfigurationCategoryTechnologies_ConfigManager               DeviceManagementConfigurationCategoryTechnologies = "configManager"
	DeviceManagementConfigurationCategoryTechnologies_EndpointPrivilegeManagement DeviceManagementConfigurationCategoryTechnologies = "endpointPrivilegeManagement"
	DeviceManagementConfigurationCategoryTechnologies_Enrollment                  DeviceManagementConfigurationCategoryTechnologies = "enrollment"
	DeviceManagementConfigurationCategoryTechnologies_ExchangeOnline              DeviceManagementConfigurationCategoryTechnologies = "exchangeOnline"
	DeviceManagementConfigurationCategoryTechnologies_LinuxMdm                    DeviceManagementConfigurationCategoryTechnologies = "linuxMdm"
	DeviceManagementConfigurationCategoryTechnologies_Mdm                         DeviceManagementConfigurationCategoryTechnologies = "mdm"
	DeviceManagementConfigurationCategoryTechnologies_MicrosoftSense              DeviceManagementConfigurationCategoryTechnologies = "microsoftSense"
	DeviceManagementConfigurationCategoryTechnologies_MobileApplicationManagement DeviceManagementConfigurationCategoryTechnologies = "mobileApplicationManagement"
	DeviceManagementConfigurationCategoryTechnologies_None                        DeviceManagementConfigurationCategoryTechnologies = "none"
	DeviceManagementConfigurationCategoryTechnologies_Windows10XManagement        DeviceManagementConfigurationCategoryTechnologies = "windows10XManagement"
)

func PossibleValuesForDeviceManagementConfigurationCategoryTechnologies() []string {
	return []string{
		string(DeviceManagementConfigurationCategoryTechnologies_AppleRemoteManagement),
		string(DeviceManagementConfigurationCategoryTechnologies_ConfigManager),
		string(DeviceManagementConfigurationCategoryTechnologies_EndpointPrivilegeManagement),
		string(DeviceManagementConfigurationCategoryTechnologies_Enrollment),
		string(DeviceManagementConfigurationCategoryTechnologies_ExchangeOnline),
		string(DeviceManagementConfigurationCategoryTechnologies_LinuxMdm),
		string(DeviceManagementConfigurationCategoryTechnologies_Mdm),
		string(DeviceManagementConfigurationCategoryTechnologies_MicrosoftSense),
		string(DeviceManagementConfigurationCategoryTechnologies_MobileApplicationManagement),
		string(DeviceManagementConfigurationCategoryTechnologies_None),
		string(DeviceManagementConfigurationCategoryTechnologies_Windows10XManagement),
	}
}

func (s *DeviceManagementConfigurationCategoryTechnologies) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationCategoryTechnologies(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationCategoryTechnologies(input string) (*DeviceManagementConfigurationCategoryTechnologies, error) {
	vals := map[string]DeviceManagementConfigurationCategoryTechnologies{
		"appleremotemanagement":       DeviceManagementConfigurationCategoryTechnologies_AppleRemoteManagement,
		"configmanager":               DeviceManagementConfigurationCategoryTechnologies_ConfigManager,
		"endpointprivilegemanagement": DeviceManagementConfigurationCategoryTechnologies_EndpointPrivilegeManagement,
		"enrollment":                  DeviceManagementConfigurationCategoryTechnologies_Enrollment,
		"exchangeonline":              DeviceManagementConfigurationCategoryTechnologies_ExchangeOnline,
		"linuxmdm":                    DeviceManagementConfigurationCategoryTechnologies_LinuxMdm,
		"mdm":                         DeviceManagementConfigurationCategoryTechnologies_Mdm,
		"microsoftsense":              DeviceManagementConfigurationCategoryTechnologies_MicrosoftSense,
		"mobileapplicationmanagement": DeviceManagementConfigurationCategoryTechnologies_MobileApplicationManagement,
		"none":                        DeviceManagementConfigurationCategoryTechnologies_None,
		"windows10xmanagement":        DeviceManagementConfigurationCategoryTechnologies_Windows10XManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationCategoryTechnologies(input)
	return &out, nil
}

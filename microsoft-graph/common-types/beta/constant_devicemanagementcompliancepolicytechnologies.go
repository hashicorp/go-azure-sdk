package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementCompliancePolicyTechnologies string

const (
	DeviceManagementCompliancePolicyTechnologies_AppleRemoteManagement       DeviceManagementCompliancePolicyTechnologies = "appleRemoteManagement"
	DeviceManagementCompliancePolicyTechnologies_ConfigManager               DeviceManagementCompliancePolicyTechnologies = "configManager"
	DeviceManagementCompliancePolicyTechnologies_EndpointPrivilegeManagement DeviceManagementCompliancePolicyTechnologies = "endpointPrivilegeManagement"
	DeviceManagementCompliancePolicyTechnologies_Enrollment                  DeviceManagementCompliancePolicyTechnologies = "enrollment"
	DeviceManagementCompliancePolicyTechnologies_ExchangeOnline              DeviceManagementCompliancePolicyTechnologies = "exchangeOnline"
	DeviceManagementCompliancePolicyTechnologies_LinuxMdm                    DeviceManagementCompliancePolicyTechnologies = "linuxMdm"
	DeviceManagementCompliancePolicyTechnologies_Mdm                         DeviceManagementCompliancePolicyTechnologies = "mdm"
	DeviceManagementCompliancePolicyTechnologies_MicrosoftSense              DeviceManagementCompliancePolicyTechnologies = "microsoftSense"
	DeviceManagementCompliancePolicyTechnologies_MobileApplicationManagement DeviceManagementCompliancePolicyTechnologies = "mobileApplicationManagement"
	DeviceManagementCompliancePolicyTechnologies_None                        DeviceManagementCompliancePolicyTechnologies = "none"
	DeviceManagementCompliancePolicyTechnologies_Windows10XManagement        DeviceManagementCompliancePolicyTechnologies = "windows10XManagement"
)

func PossibleValuesForDeviceManagementCompliancePolicyTechnologies() []string {
	return []string{
		string(DeviceManagementCompliancePolicyTechnologies_AppleRemoteManagement),
		string(DeviceManagementCompliancePolicyTechnologies_ConfigManager),
		string(DeviceManagementCompliancePolicyTechnologies_EndpointPrivilegeManagement),
		string(DeviceManagementCompliancePolicyTechnologies_Enrollment),
		string(DeviceManagementCompliancePolicyTechnologies_ExchangeOnline),
		string(DeviceManagementCompliancePolicyTechnologies_LinuxMdm),
		string(DeviceManagementCompliancePolicyTechnologies_Mdm),
		string(DeviceManagementCompliancePolicyTechnologies_MicrosoftSense),
		string(DeviceManagementCompliancePolicyTechnologies_MobileApplicationManagement),
		string(DeviceManagementCompliancePolicyTechnologies_None),
		string(DeviceManagementCompliancePolicyTechnologies_Windows10XManagement),
	}
}

func (s *DeviceManagementCompliancePolicyTechnologies) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementCompliancePolicyTechnologies(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementCompliancePolicyTechnologies(input string) (*DeviceManagementCompliancePolicyTechnologies, error) {
	vals := map[string]DeviceManagementCompliancePolicyTechnologies{
		"appleremotemanagement":       DeviceManagementCompliancePolicyTechnologies_AppleRemoteManagement,
		"configmanager":               DeviceManagementCompliancePolicyTechnologies_ConfigManager,
		"endpointprivilegemanagement": DeviceManagementCompliancePolicyTechnologies_EndpointPrivilegeManagement,
		"enrollment":                  DeviceManagementCompliancePolicyTechnologies_Enrollment,
		"exchangeonline":              DeviceManagementCompliancePolicyTechnologies_ExchangeOnline,
		"linuxmdm":                    DeviceManagementCompliancePolicyTechnologies_LinuxMdm,
		"mdm":                         DeviceManagementCompliancePolicyTechnologies_Mdm,
		"microsoftsense":              DeviceManagementCompliancePolicyTechnologies_MicrosoftSense,
		"mobileapplicationmanagement": DeviceManagementCompliancePolicyTechnologies_MobileApplicationManagement,
		"none":                        DeviceManagementCompliancePolicyTechnologies_None,
		"windows10xmanagement":        DeviceManagementCompliancePolicyTechnologies_Windows10XManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementCompliancePolicyTechnologies(input)
	return &out, nil
}

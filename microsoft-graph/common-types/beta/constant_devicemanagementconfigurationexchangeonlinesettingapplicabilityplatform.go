package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform string

const (
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform_Android    DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform = "android"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform_IOS        DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform = "iOS"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform_Linux      DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform = "linux"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform_MacOS      DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform = "macOS"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform_None       DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform = "none"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform_Windows10  DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform = "windows10"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform_Windows10X DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform = "windows10X"
)

func PossibleValuesForDeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform() []string {
	return []string{
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform_Android),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform_IOS),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform_Linux),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform_MacOS),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform_None),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform_Windows10),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform_Windows10X),
	}
}

func (s *DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform(input string) (*DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform, error) {
	vals := map[string]DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform{
		"android":    DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform_Android,
		"ios":        DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform_IOS,
		"linux":      DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform_Linux,
		"macos":      DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform_MacOS,
		"none":       DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform_None,
		"windows10":  DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform_Windows10,
		"windows10x": DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform_Windows10X,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform(input)
	return &out, nil
}

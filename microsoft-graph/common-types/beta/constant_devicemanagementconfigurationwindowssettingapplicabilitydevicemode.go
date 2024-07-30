package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationWindowsSettingApplicabilityDeviceMode string

const (
	DeviceManagementConfigurationWindowsSettingApplicabilityDeviceMode_Kiosk DeviceManagementConfigurationWindowsSettingApplicabilityDeviceMode = "kiosk"
	DeviceManagementConfigurationWindowsSettingApplicabilityDeviceMode_None  DeviceManagementConfigurationWindowsSettingApplicabilityDeviceMode = "none"
)

func PossibleValuesForDeviceManagementConfigurationWindowsSettingApplicabilityDeviceMode() []string {
	return []string{
		string(DeviceManagementConfigurationWindowsSettingApplicabilityDeviceMode_Kiosk),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityDeviceMode_None),
	}
}

func (s *DeviceManagementConfigurationWindowsSettingApplicabilityDeviceMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationWindowsSettingApplicabilityDeviceMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationWindowsSettingApplicabilityDeviceMode(input string) (*DeviceManagementConfigurationWindowsSettingApplicabilityDeviceMode, error) {
	vals := map[string]DeviceManagementConfigurationWindowsSettingApplicabilityDeviceMode{
		"kiosk": DeviceManagementConfigurationWindowsSettingApplicabilityDeviceMode_Kiosk,
		"none":  DeviceManagementConfigurationWindowsSettingApplicabilityDeviceMode_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationWindowsSettingApplicabilityDeviceMode(input)
	return &out, nil
}

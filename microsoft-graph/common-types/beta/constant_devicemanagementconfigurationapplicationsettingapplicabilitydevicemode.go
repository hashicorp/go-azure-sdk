package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode string

const (
	DeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode_Kiosk DeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode = "kiosk"
	DeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode_None  DeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode = "none"
)

func PossibleValuesForDeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode() []string {
	return []string{
		string(DeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode_Kiosk),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode_None),
	}
}

func (s *DeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode(input string) (*DeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode, error) {
	vals := map[string]DeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode{
		"kiosk": DeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode_Kiosk,
		"none":  DeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode(input)
	return &out, nil
}

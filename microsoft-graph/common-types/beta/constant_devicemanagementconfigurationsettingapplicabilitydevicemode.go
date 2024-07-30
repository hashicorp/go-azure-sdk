package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingApplicabilityDeviceMode string

const (
	DeviceManagementConfigurationSettingApplicabilityDeviceMode_Kiosk DeviceManagementConfigurationSettingApplicabilityDeviceMode = "kiosk"
	DeviceManagementConfigurationSettingApplicabilityDeviceMode_None  DeviceManagementConfigurationSettingApplicabilityDeviceMode = "none"
)

func PossibleValuesForDeviceManagementConfigurationSettingApplicabilityDeviceMode() []string {
	return []string{
		string(DeviceManagementConfigurationSettingApplicabilityDeviceMode_Kiosk),
		string(DeviceManagementConfigurationSettingApplicabilityDeviceMode_None),
	}
}

func (s *DeviceManagementConfigurationSettingApplicabilityDeviceMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSettingApplicabilityDeviceMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSettingApplicabilityDeviceMode(input string) (*DeviceManagementConfigurationSettingApplicabilityDeviceMode, error) {
	vals := map[string]DeviceManagementConfigurationSettingApplicabilityDeviceMode{
		"kiosk": DeviceManagementConfigurationSettingApplicabilityDeviceMode_Kiosk,
		"none":  DeviceManagementConfigurationSettingApplicabilityDeviceMode_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingApplicabilityDeviceMode(input)
	return &out, nil
}

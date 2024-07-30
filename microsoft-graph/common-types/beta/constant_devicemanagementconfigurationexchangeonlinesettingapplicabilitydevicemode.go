package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode string

const (
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode_Kiosk DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode = "kiosk"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode_None  DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode = "none"
)

func PossibleValuesForDeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode() []string {
	return []string{
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode_Kiosk),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode_None),
	}
}

func (s *DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode(input string) (*DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode, error) {
	vals := map[string]DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode{
		"kiosk": DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode_Kiosk,
		"none":  DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode(input)
	return &out, nil
}

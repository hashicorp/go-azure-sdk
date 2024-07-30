package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentNotificationConfigurationPlatformType string

const (
	DeviceEnrollmentNotificationConfigurationPlatformType_AllPlatforms   DeviceEnrollmentNotificationConfigurationPlatformType = "allPlatforms"
	DeviceEnrollmentNotificationConfigurationPlatformType_Android        DeviceEnrollmentNotificationConfigurationPlatformType = "android"
	DeviceEnrollmentNotificationConfigurationPlatformType_AndroidForWork DeviceEnrollmentNotificationConfigurationPlatformType = "androidForWork"
	DeviceEnrollmentNotificationConfigurationPlatformType_Ios            DeviceEnrollmentNotificationConfigurationPlatformType = "ios"
	DeviceEnrollmentNotificationConfigurationPlatformType_Linux          DeviceEnrollmentNotificationConfigurationPlatformType = "linux"
	DeviceEnrollmentNotificationConfigurationPlatformType_Mac            DeviceEnrollmentNotificationConfigurationPlatformType = "mac"
	DeviceEnrollmentNotificationConfigurationPlatformType_Windows        DeviceEnrollmentNotificationConfigurationPlatformType = "windows"
	DeviceEnrollmentNotificationConfigurationPlatformType_WindowsPhone   DeviceEnrollmentNotificationConfigurationPlatformType = "windowsPhone"
)

func PossibleValuesForDeviceEnrollmentNotificationConfigurationPlatformType() []string {
	return []string{
		string(DeviceEnrollmentNotificationConfigurationPlatformType_AllPlatforms),
		string(DeviceEnrollmentNotificationConfigurationPlatformType_Android),
		string(DeviceEnrollmentNotificationConfigurationPlatformType_AndroidForWork),
		string(DeviceEnrollmentNotificationConfigurationPlatformType_Ios),
		string(DeviceEnrollmentNotificationConfigurationPlatformType_Linux),
		string(DeviceEnrollmentNotificationConfigurationPlatformType_Mac),
		string(DeviceEnrollmentNotificationConfigurationPlatformType_Windows),
		string(DeviceEnrollmentNotificationConfigurationPlatformType_WindowsPhone),
	}
}

func (s *DeviceEnrollmentNotificationConfigurationPlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceEnrollmentNotificationConfigurationPlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceEnrollmentNotificationConfigurationPlatformType(input string) (*DeviceEnrollmentNotificationConfigurationPlatformType, error) {
	vals := map[string]DeviceEnrollmentNotificationConfigurationPlatformType{
		"allplatforms":   DeviceEnrollmentNotificationConfigurationPlatformType_AllPlatforms,
		"android":        DeviceEnrollmentNotificationConfigurationPlatformType_Android,
		"androidforwork": DeviceEnrollmentNotificationConfigurationPlatformType_AndroidForWork,
		"ios":            DeviceEnrollmentNotificationConfigurationPlatformType_Ios,
		"linux":          DeviceEnrollmentNotificationConfigurationPlatformType_Linux,
		"mac":            DeviceEnrollmentNotificationConfigurationPlatformType_Mac,
		"windows":        DeviceEnrollmentNotificationConfigurationPlatformType_Windows,
		"windowsphone":   DeviceEnrollmentNotificationConfigurationPlatformType_WindowsPhone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentNotificationConfigurationPlatformType(input)
	return &out, nil
}

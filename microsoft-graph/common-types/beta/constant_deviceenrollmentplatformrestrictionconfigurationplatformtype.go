package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentPlatformRestrictionConfigurationPlatformType string

const (
	DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_AllPlatforms   DeviceEnrollmentPlatformRestrictionConfigurationPlatformType = "allPlatforms"
	DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_Android        DeviceEnrollmentPlatformRestrictionConfigurationPlatformType = "android"
	DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_AndroidForWork DeviceEnrollmentPlatformRestrictionConfigurationPlatformType = "androidForWork"
	DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_Ios            DeviceEnrollmentPlatformRestrictionConfigurationPlatformType = "ios"
	DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_Linux          DeviceEnrollmentPlatformRestrictionConfigurationPlatformType = "linux"
	DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_Mac            DeviceEnrollmentPlatformRestrictionConfigurationPlatformType = "mac"
	DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_Windows        DeviceEnrollmentPlatformRestrictionConfigurationPlatformType = "windows"
	DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_WindowsPhone   DeviceEnrollmentPlatformRestrictionConfigurationPlatformType = "windowsPhone"
)

func PossibleValuesForDeviceEnrollmentPlatformRestrictionConfigurationPlatformType() []string {
	return []string{
		string(DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_AllPlatforms),
		string(DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_Android),
		string(DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_AndroidForWork),
		string(DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_Ios),
		string(DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_Linux),
		string(DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_Mac),
		string(DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_Windows),
		string(DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_WindowsPhone),
	}
}

func (s *DeviceEnrollmentPlatformRestrictionConfigurationPlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceEnrollmentPlatformRestrictionConfigurationPlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceEnrollmentPlatformRestrictionConfigurationPlatformType(input string) (*DeviceEnrollmentPlatformRestrictionConfigurationPlatformType, error) {
	vals := map[string]DeviceEnrollmentPlatformRestrictionConfigurationPlatformType{
		"allplatforms":   DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_AllPlatforms,
		"android":        DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_Android,
		"androidforwork": DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_AndroidForWork,
		"ios":            DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_Ios,
		"linux":          DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_Linux,
		"mac":            DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_Mac,
		"windows":        DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_Windows,
		"windowsphone":   DeviceEnrollmentPlatformRestrictionConfigurationPlatformType_WindowsPhone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentPlatformRestrictionConfigurationPlatformType(input)
	return &out, nil
}

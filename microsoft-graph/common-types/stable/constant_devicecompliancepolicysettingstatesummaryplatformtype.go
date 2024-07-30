package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicySettingStateSummaryPlatformType string

const (
	DeviceCompliancePolicySettingStateSummaryPlatformType_All               DeviceCompliancePolicySettingStateSummaryPlatformType = "all"
	DeviceCompliancePolicySettingStateSummaryPlatformType_Android           DeviceCompliancePolicySettingStateSummaryPlatformType = "android"
	DeviceCompliancePolicySettingStateSummaryPlatformType_AndroidForWork    DeviceCompliancePolicySettingStateSummaryPlatformType = "androidForWork"
	DeviceCompliancePolicySettingStateSummaryPlatformType_IOS               DeviceCompliancePolicySettingStateSummaryPlatformType = "iOS"
	DeviceCompliancePolicySettingStateSummaryPlatformType_MacOS             DeviceCompliancePolicySettingStateSummaryPlatformType = "macOS"
	DeviceCompliancePolicySettingStateSummaryPlatformType_Windows10AndLater DeviceCompliancePolicySettingStateSummaryPlatformType = "windows10AndLater"
	DeviceCompliancePolicySettingStateSummaryPlatformType_Windows81AndLater DeviceCompliancePolicySettingStateSummaryPlatformType = "windows81AndLater"
	DeviceCompliancePolicySettingStateSummaryPlatformType_WindowsPhone81    DeviceCompliancePolicySettingStateSummaryPlatformType = "windowsPhone81"
)

func PossibleValuesForDeviceCompliancePolicySettingStateSummaryPlatformType() []string {
	return []string{
		string(DeviceCompliancePolicySettingStateSummaryPlatformType_All),
		string(DeviceCompliancePolicySettingStateSummaryPlatformType_Android),
		string(DeviceCompliancePolicySettingStateSummaryPlatformType_AndroidForWork),
		string(DeviceCompliancePolicySettingStateSummaryPlatformType_IOS),
		string(DeviceCompliancePolicySettingStateSummaryPlatformType_MacOS),
		string(DeviceCompliancePolicySettingStateSummaryPlatformType_Windows10AndLater),
		string(DeviceCompliancePolicySettingStateSummaryPlatformType_Windows81AndLater),
		string(DeviceCompliancePolicySettingStateSummaryPlatformType_WindowsPhone81),
	}
}

func (s *DeviceCompliancePolicySettingStateSummaryPlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceCompliancePolicySettingStateSummaryPlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceCompliancePolicySettingStateSummaryPlatformType(input string) (*DeviceCompliancePolicySettingStateSummaryPlatformType, error) {
	vals := map[string]DeviceCompliancePolicySettingStateSummaryPlatformType{
		"all":               DeviceCompliancePolicySettingStateSummaryPlatformType_All,
		"android":           DeviceCompliancePolicySettingStateSummaryPlatformType_Android,
		"androidforwork":    DeviceCompliancePolicySettingStateSummaryPlatformType_AndroidForWork,
		"ios":               DeviceCompliancePolicySettingStateSummaryPlatformType_IOS,
		"macos":             DeviceCompliancePolicySettingStateSummaryPlatformType_MacOS,
		"windows10andlater": DeviceCompliancePolicySettingStateSummaryPlatformType_Windows10AndLater,
		"windows81andlater": DeviceCompliancePolicySettingStateSummaryPlatformType_Windows81AndLater,
		"windowsphone81":    DeviceCompliancePolicySettingStateSummaryPlatformType_WindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceCompliancePolicySettingStateSummaryPlatformType(input)
	return &out, nil
}

package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingComparisonComparisonResult string

const (
	DeviceManagementSettingComparisonComparisonResult_Added    DeviceManagementSettingComparisonComparisonResult = "added"
	DeviceManagementSettingComparisonComparisonResult_Equal    DeviceManagementSettingComparisonComparisonResult = "equal"
	DeviceManagementSettingComparisonComparisonResult_NotEqual DeviceManagementSettingComparisonComparisonResult = "notEqual"
	DeviceManagementSettingComparisonComparisonResult_Removed  DeviceManagementSettingComparisonComparisonResult = "removed"
	DeviceManagementSettingComparisonComparisonResult_Unknown  DeviceManagementSettingComparisonComparisonResult = "unknown"
)

func PossibleValuesForDeviceManagementSettingComparisonComparisonResult() []string {
	return []string{
		string(DeviceManagementSettingComparisonComparisonResult_Added),
		string(DeviceManagementSettingComparisonComparisonResult_Equal),
		string(DeviceManagementSettingComparisonComparisonResult_NotEqual),
		string(DeviceManagementSettingComparisonComparisonResult_Removed),
		string(DeviceManagementSettingComparisonComparisonResult_Unknown),
	}
}

func (s *DeviceManagementSettingComparisonComparisonResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementSettingComparisonComparisonResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementSettingComparisonComparisonResult(input string) (*DeviceManagementSettingComparisonComparisonResult, error) {
	vals := map[string]DeviceManagementSettingComparisonComparisonResult{
		"added":    DeviceManagementSettingComparisonComparisonResult_Added,
		"equal":    DeviceManagementSettingComparisonComparisonResult_Equal,
		"notequal": DeviceManagementSettingComparisonComparisonResult_NotEqual,
		"removed":  DeviceManagementSettingComparisonComparisonResult_Removed,
		"unknown":  DeviceManagementSettingComparisonComparisonResult_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementSettingComparisonComparisonResult(input)
	return &out, nil
}

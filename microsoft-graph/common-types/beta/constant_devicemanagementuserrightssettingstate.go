package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementUserRightsSettingState string

const (
	DeviceManagementUserRightsSettingState_Allowed       DeviceManagementUserRightsSettingState = "allowed"
	DeviceManagementUserRightsSettingState_Blocked       DeviceManagementUserRightsSettingState = "blocked"
	DeviceManagementUserRightsSettingState_NotConfigured DeviceManagementUserRightsSettingState = "notConfigured"
)

func PossibleValuesForDeviceManagementUserRightsSettingState() []string {
	return []string{
		string(DeviceManagementUserRightsSettingState_Allowed),
		string(DeviceManagementUserRightsSettingState_Blocked),
		string(DeviceManagementUserRightsSettingState_NotConfigured),
	}
}

func (s *DeviceManagementUserRightsSettingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementUserRightsSettingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementUserRightsSettingState(input string) (*DeviceManagementUserRightsSettingState, error) {
	vals := map[string]DeviceManagementUserRightsSettingState{
		"allowed":       DeviceManagementUserRightsSettingState_Allowed,
		"blocked":       DeviceManagementUserRightsSettingState_Blocked,
		"notconfigured": DeviceManagementUserRightsSettingState_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementUserRightsSettingState(input)
	return &out, nil
}

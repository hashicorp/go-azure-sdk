package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceActionItemActionType string

const (
	DeviceComplianceActionItemActionType_Block                        DeviceComplianceActionItemActionType = "block"
	DeviceComplianceActionItemActionType_NoAction                     DeviceComplianceActionItemActionType = "noAction"
	DeviceComplianceActionItemActionType_Notification                 DeviceComplianceActionItemActionType = "notification"
	DeviceComplianceActionItemActionType_PushNotification             DeviceComplianceActionItemActionType = "pushNotification"
	DeviceComplianceActionItemActionType_RemoveResourceAccessProfiles DeviceComplianceActionItemActionType = "removeResourceAccessProfiles"
	DeviceComplianceActionItemActionType_Retire                       DeviceComplianceActionItemActionType = "retire"
	DeviceComplianceActionItemActionType_Wipe                         DeviceComplianceActionItemActionType = "wipe"
)

func PossibleValuesForDeviceComplianceActionItemActionType() []string {
	return []string{
		string(DeviceComplianceActionItemActionType_Block),
		string(DeviceComplianceActionItemActionType_NoAction),
		string(DeviceComplianceActionItemActionType_Notification),
		string(DeviceComplianceActionItemActionType_PushNotification),
		string(DeviceComplianceActionItemActionType_RemoveResourceAccessProfiles),
		string(DeviceComplianceActionItemActionType_Retire),
		string(DeviceComplianceActionItemActionType_Wipe),
	}
}

func (s *DeviceComplianceActionItemActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceComplianceActionItemActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceComplianceActionItemActionType(input string) (*DeviceComplianceActionItemActionType, error) {
	vals := map[string]DeviceComplianceActionItemActionType{
		"block":                        DeviceComplianceActionItemActionType_Block,
		"noaction":                     DeviceComplianceActionItemActionType_NoAction,
		"notification":                 DeviceComplianceActionItemActionType_Notification,
		"pushnotification":             DeviceComplianceActionItemActionType_PushNotification,
		"removeresourceaccessprofiles": DeviceComplianceActionItemActionType_RemoveResourceAccessProfiles,
		"retire":                       DeviceComplianceActionItemActionType_Retire,
		"wipe":                         DeviceComplianceActionItemActionType_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceActionItemActionType(input)
	return &out, nil
}

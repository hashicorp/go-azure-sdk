package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementComplianceActionItemActionType string

const (
	DeviceManagementComplianceActionItemActionType_Block                        DeviceManagementComplianceActionItemActionType = "block"
	DeviceManagementComplianceActionItemActionType_NoAction                     DeviceManagementComplianceActionItemActionType = "noAction"
	DeviceManagementComplianceActionItemActionType_Notification                 DeviceManagementComplianceActionItemActionType = "notification"
	DeviceManagementComplianceActionItemActionType_PushNotification             DeviceManagementComplianceActionItemActionType = "pushNotification"
	DeviceManagementComplianceActionItemActionType_RemoteLock                   DeviceManagementComplianceActionItemActionType = "remoteLock"
	DeviceManagementComplianceActionItemActionType_RemoveResourceAccessProfiles DeviceManagementComplianceActionItemActionType = "removeResourceAccessProfiles"
	DeviceManagementComplianceActionItemActionType_Retire                       DeviceManagementComplianceActionItemActionType = "retire"
	DeviceManagementComplianceActionItemActionType_Wipe                         DeviceManagementComplianceActionItemActionType = "wipe"
)

func PossibleValuesForDeviceManagementComplianceActionItemActionType() []string {
	return []string{
		string(DeviceManagementComplianceActionItemActionType_Block),
		string(DeviceManagementComplianceActionItemActionType_NoAction),
		string(DeviceManagementComplianceActionItemActionType_Notification),
		string(DeviceManagementComplianceActionItemActionType_PushNotification),
		string(DeviceManagementComplianceActionItemActionType_RemoteLock),
		string(DeviceManagementComplianceActionItemActionType_RemoveResourceAccessProfiles),
		string(DeviceManagementComplianceActionItemActionType_Retire),
		string(DeviceManagementComplianceActionItemActionType_Wipe),
	}
}

func (s *DeviceManagementComplianceActionItemActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementComplianceActionItemActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementComplianceActionItemActionType(input string) (*DeviceManagementComplianceActionItemActionType, error) {
	vals := map[string]DeviceManagementComplianceActionItemActionType{
		"block":                        DeviceManagementComplianceActionItemActionType_Block,
		"noaction":                     DeviceManagementComplianceActionItemActionType_NoAction,
		"notification":                 DeviceManagementComplianceActionItemActionType_Notification,
		"pushnotification":             DeviceManagementComplianceActionItemActionType_PushNotification,
		"remotelock":                   DeviceManagementComplianceActionItemActionType_RemoteLock,
		"removeresourceaccessprofiles": DeviceManagementComplianceActionItemActionType_RemoveResourceAccessProfiles,
		"retire":                       DeviceManagementComplianceActionItemActionType_Retire,
		"wipe":                         DeviceManagementComplianceActionItemActionType_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementComplianceActionItemActionType(input)
	return &out, nil
}

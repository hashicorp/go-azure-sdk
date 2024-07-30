package devicecompliancepolicy

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceActionItemActionType string

const (
	DeviceComplianceActionItemActionTypeBlock                        DeviceComplianceActionItemActionType = "block"
	DeviceComplianceActionItemActionTypeNoAction                     DeviceComplianceActionItemActionType = "noAction"
	DeviceComplianceActionItemActionTypeNotification                 DeviceComplianceActionItemActionType = "notification"
	DeviceComplianceActionItemActionTypePushNotification             DeviceComplianceActionItemActionType = "pushNotification"
	DeviceComplianceActionItemActionTypeRemoveResourceAccessProfiles DeviceComplianceActionItemActionType = "removeResourceAccessProfiles"
	DeviceComplianceActionItemActionTypeRetire                       DeviceComplianceActionItemActionType = "retire"
	DeviceComplianceActionItemActionTypeWipe                         DeviceComplianceActionItemActionType = "wipe"
)

func PossibleValuesForDeviceComplianceActionItemActionType() []string {
	return []string{
		string(DeviceComplianceActionItemActionTypeBlock),
		string(DeviceComplianceActionItemActionTypeNoAction),
		string(DeviceComplianceActionItemActionTypeNotification),
		string(DeviceComplianceActionItemActionTypePushNotification),
		string(DeviceComplianceActionItemActionTypeRemoveResourceAccessProfiles),
		string(DeviceComplianceActionItemActionTypeRetire),
		string(DeviceComplianceActionItemActionTypeWipe),
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
		"block":                        DeviceComplianceActionItemActionTypeBlock,
		"noaction":                     DeviceComplianceActionItemActionTypeNoAction,
		"notification":                 DeviceComplianceActionItemActionTypeNotification,
		"pushnotification":             DeviceComplianceActionItemActionTypePushNotification,
		"removeresourceaccessprofiles": DeviceComplianceActionItemActionTypeRemoveResourceAccessProfiles,
		"retire":                       DeviceComplianceActionItemActionTypeRetire,
		"wipe":                         DeviceComplianceActionItemActionTypeWipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceActionItemActionType(input)
	return &out, nil
}

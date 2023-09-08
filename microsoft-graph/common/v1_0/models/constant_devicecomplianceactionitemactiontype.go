package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceActionItemActionType string

const (
	DeviceComplianceActionItemActionTypeblock                        DeviceComplianceActionItemActionType = "Block"
	DeviceComplianceActionItemActionTypenoAction                     DeviceComplianceActionItemActionType = "NoAction"
	DeviceComplianceActionItemActionTypenotification                 DeviceComplianceActionItemActionType = "Notification"
	DeviceComplianceActionItemActionTypepushNotification             DeviceComplianceActionItemActionType = "PushNotification"
	DeviceComplianceActionItemActionTyperemoveResourceAccessProfiles DeviceComplianceActionItemActionType = "RemoveResourceAccessProfiles"
	DeviceComplianceActionItemActionTyperetire                       DeviceComplianceActionItemActionType = "Retire"
	DeviceComplianceActionItemActionTypewipe                         DeviceComplianceActionItemActionType = "Wipe"
)

func PossibleValuesForDeviceComplianceActionItemActionType() []string {
	return []string{
		string(DeviceComplianceActionItemActionTypeblock),
		string(DeviceComplianceActionItemActionTypenoAction),
		string(DeviceComplianceActionItemActionTypenotification),
		string(DeviceComplianceActionItemActionTypepushNotification),
		string(DeviceComplianceActionItemActionTyperemoveResourceAccessProfiles),
		string(DeviceComplianceActionItemActionTyperetire),
		string(DeviceComplianceActionItemActionTypewipe),
	}
}

func parseDeviceComplianceActionItemActionType(input string) (*DeviceComplianceActionItemActionType, error) {
	vals := map[string]DeviceComplianceActionItemActionType{
		"block":                        DeviceComplianceActionItemActionTypeblock,
		"noaction":                     DeviceComplianceActionItemActionTypenoAction,
		"notification":                 DeviceComplianceActionItemActionTypenotification,
		"pushnotification":             DeviceComplianceActionItemActionTypepushNotification,
		"removeresourceaccessprofiles": DeviceComplianceActionItemActionTyperemoveResourceAccessProfiles,
		"retire":                       DeviceComplianceActionItemActionTyperetire,
		"wipe":                         DeviceComplianceActionItemActionTypewipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceActionItemActionType(input)
	return &out, nil
}

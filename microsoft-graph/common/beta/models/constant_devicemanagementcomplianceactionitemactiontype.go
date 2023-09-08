package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementComplianceActionItemActionType string

const (
	DeviceManagementComplianceActionItemActionTypeblock                        DeviceManagementComplianceActionItemActionType = "Block"
	DeviceManagementComplianceActionItemActionTypenoAction                     DeviceManagementComplianceActionItemActionType = "NoAction"
	DeviceManagementComplianceActionItemActionTypenotification                 DeviceManagementComplianceActionItemActionType = "Notification"
	DeviceManagementComplianceActionItemActionTypepushNotification             DeviceManagementComplianceActionItemActionType = "PushNotification"
	DeviceManagementComplianceActionItemActionTyperemoteLock                   DeviceManagementComplianceActionItemActionType = "RemoteLock"
	DeviceManagementComplianceActionItemActionTyperemoveResourceAccessProfiles DeviceManagementComplianceActionItemActionType = "RemoveResourceAccessProfiles"
	DeviceManagementComplianceActionItemActionTyperetire                       DeviceManagementComplianceActionItemActionType = "Retire"
	DeviceManagementComplianceActionItemActionTypewipe                         DeviceManagementComplianceActionItemActionType = "Wipe"
)

func PossibleValuesForDeviceManagementComplianceActionItemActionType() []string {
	return []string{
		string(DeviceManagementComplianceActionItemActionTypeblock),
		string(DeviceManagementComplianceActionItemActionTypenoAction),
		string(DeviceManagementComplianceActionItemActionTypenotification),
		string(DeviceManagementComplianceActionItemActionTypepushNotification),
		string(DeviceManagementComplianceActionItemActionTyperemoteLock),
		string(DeviceManagementComplianceActionItemActionTyperemoveResourceAccessProfiles),
		string(DeviceManagementComplianceActionItemActionTyperetire),
		string(DeviceManagementComplianceActionItemActionTypewipe),
	}
}

func parseDeviceManagementComplianceActionItemActionType(input string) (*DeviceManagementComplianceActionItemActionType, error) {
	vals := map[string]DeviceManagementComplianceActionItemActionType{
		"block":                        DeviceManagementComplianceActionItemActionTypeblock,
		"noaction":                     DeviceManagementComplianceActionItemActionTypenoAction,
		"notification":                 DeviceManagementComplianceActionItemActionTypenotification,
		"pushnotification":             DeviceManagementComplianceActionItemActionTypepushNotification,
		"remotelock":                   DeviceManagementComplianceActionItemActionTyperemoteLock,
		"removeresourceaccessprofiles": DeviceManagementComplianceActionItemActionTyperemoveResourceAccessProfiles,
		"retire":                       DeviceManagementComplianceActionItemActionTyperetire,
		"wipe":                         DeviceManagementComplianceActionItemActionTypewipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementComplianceActionItemActionType(input)
	return &out, nil
}

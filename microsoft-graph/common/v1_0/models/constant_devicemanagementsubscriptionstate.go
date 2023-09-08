package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSubscriptionState string

const (
	DeviceManagementSubscriptionStateactive    DeviceManagementSubscriptionState = "Active"
	DeviceManagementSubscriptionStateblocked   DeviceManagementSubscriptionState = "Blocked"
	DeviceManagementSubscriptionStatedeleted   DeviceManagementSubscriptionState = "Deleted"
	DeviceManagementSubscriptionStatedisabled  DeviceManagementSubscriptionState = "Disabled"
	DeviceManagementSubscriptionStatelockedOut DeviceManagementSubscriptionState = "LockedOut"
	DeviceManagementSubscriptionStatepending   DeviceManagementSubscriptionState = "Pending"
	DeviceManagementSubscriptionStatewarning   DeviceManagementSubscriptionState = "Warning"
)

func PossibleValuesForDeviceManagementSubscriptionState() []string {
	return []string{
		string(DeviceManagementSubscriptionStateactive),
		string(DeviceManagementSubscriptionStateblocked),
		string(DeviceManagementSubscriptionStatedeleted),
		string(DeviceManagementSubscriptionStatedisabled),
		string(DeviceManagementSubscriptionStatelockedOut),
		string(DeviceManagementSubscriptionStatepending),
		string(DeviceManagementSubscriptionStatewarning),
	}
}

func parseDeviceManagementSubscriptionState(input string) (*DeviceManagementSubscriptionState, error) {
	vals := map[string]DeviceManagementSubscriptionState{
		"active":    DeviceManagementSubscriptionStateactive,
		"blocked":   DeviceManagementSubscriptionStateblocked,
		"deleted":   DeviceManagementSubscriptionStatedeleted,
		"disabled":  DeviceManagementSubscriptionStatedisabled,
		"lockedout": DeviceManagementSubscriptionStatelockedOut,
		"pending":   DeviceManagementSubscriptionStatepending,
		"warning":   DeviceManagementSubscriptionStatewarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementSubscriptionState(input)
	return &out, nil
}

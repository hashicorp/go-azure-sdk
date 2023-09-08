package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExchangeAccessRuleAccessLevel string

const (
	DeviceManagementExchangeAccessRuleAccessLevelallow      DeviceManagementExchangeAccessRuleAccessLevel = "Allow"
	DeviceManagementExchangeAccessRuleAccessLevelblock      DeviceManagementExchangeAccessRuleAccessLevel = "Block"
	DeviceManagementExchangeAccessRuleAccessLevelnone       DeviceManagementExchangeAccessRuleAccessLevel = "None"
	DeviceManagementExchangeAccessRuleAccessLevelquarantine DeviceManagementExchangeAccessRuleAccessLevel = "Quarantine"
)

func PossibleValuesForDeviceManagementExchangeAccessRuleAccessLevel() []string {
	return []string{
		string(DeviceManagementExchangeAccessRuleAccessLevelallow),
		string(DeviceManagementExchangeAccessRuleAccessLevelblock),
		string(DeviceManagementExchangeAccessRuleAccessLevelnone),
		string(DeviceManagementExchangeAccessRuleAccessLevelquarantine),
	}
}

func parseDeviceManagementExchangeAccessRuleAccessLevel(input string) (*DeviceManagementExchangeAccessRuleAccessLevel, error) {
	vals := map[string]DeviceManagementExchangeAccessRuleAccessLevel{
		"allow":      DeviceManagementExchangeAccessRuleAccessLevelallow,
		"block":      DeviceManagementExchangeAccessRuleAccessLevelblock,
		"none":       DeviceManagementExchangeAccessRuleAccessLevelnone,
		"quarantine": DeviceManagementExchangeAccessRuleAccessLevelquarantine,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementExchangeAccessRuleAccessLevel(input)
	return &out, nil
}

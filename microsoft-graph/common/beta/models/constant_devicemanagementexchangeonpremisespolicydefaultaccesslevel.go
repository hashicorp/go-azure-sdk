package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel string

const (
	DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevelallow      DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel = "Allow"
	DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevelblock      DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel = "Block"
	DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevelnone       DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel = "None"
	DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevelquarantine DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel = "Quarantine"
)

func PossibleValuesForDeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel() []string {
	return []string{
		string(DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevelallow),
		string(DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevelblock),
		string(DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevelnone),
		string(DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevelquarantine),
	}
}

func parseDeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel(input string) (*DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel, error) {
	vals := map[string]DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel{
		"allow":      DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevelallow,
		"block":      DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevelblock,
		"none":       DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevelnone,
		"quarantine": DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevelquarantine,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementExchangeOnPremisesPolicyDefaultAccessLevel(input)
	return &out, nil
}

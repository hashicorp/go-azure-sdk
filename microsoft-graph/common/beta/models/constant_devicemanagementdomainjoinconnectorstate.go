package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementDomainJoinConnectorState string

const (
	DeviceManagementDomainJoinConnectorStateactive   DeviceManagementDomainJoinConnectorState = "Active"
	DeviceManagementDomainJoinConnectorStateerror    DeviceManagementDomainJoinConnectorState = "Error"
	DeviceManagementDomainJoinConnectorStateinactive DeviceManagementDomainJoinConnectorState = "Inactive"
)

func PossibleValuesForDeviceManagementDomainJoinConnectorState() []string {
	return []string{
		string(DeviceManagementDomainJoinConnectorStateactive),
		string(DeviceManagementDomainJoinConnectorStateerror),
		string(DeviceManagementDomainJoinConnectorStateinactive),
	}
}

func parseDeviceManagementDomainJoinConnectorState(input string) (*DeviceManagementDomainJoinConnectorState, error) {
	vals := map[string]DeviceManagementDomainJoinConnectorState{
		"active":   DeviceManagementDomainJoinConnectorStateactive,
		"error":    DeviceManagementDomainJoinConnectorStateerror,
		"inactive": DeviceManagementDomainJoinConnectorStateinactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementDomainJoinConnectorState(input)
	return &out, nil
}
